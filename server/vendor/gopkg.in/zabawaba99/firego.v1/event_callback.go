package firego

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/zabawaba99/firego/sync"
)

// ChildEventFunc is the type of function that is called for every
// new child added under a firebase reference. The snapshot argument
// contains the data that was added. The previousChildKey argument
// contains the key of the previous child that this function was called for.
type ChildEventFunc func(snapshot DataSnapshot, previousChildKey string)

// ChildAdded listens on the firebase instance and executes the callback
// for every child that is added.
//
// You cannot set the same function twice on a Firebase reference, if you do
// the first function will be overridden and you will not be able to close the
// connection.
func (fb *Firebase) ChildAdded(fn ChildEventFunc) error {
	return fb.addEventFunc(fn, fn.childAdded)
}

func (fn ChildEventFunc) childAdded(db *sync.Database, prevKey *string, notifications chan Event) error {
	for event := range notifications {
		if event.Type == EventTypeError {
			err, ok := event.Data.(error)
			if !ok {
				err = fmt.Errorf("Got error from event %#v", event)
			}
			return err
		}

		if event.Type != EventTypePut {
			continue
		}

		child := strings.Split(event.Path[1:], "/")[0]
		if event.Data == nil {
			db.Del(child)
			continue
		}

		if _, ok := db.Get("").Child(child); ok {
			// if the child isn't being added, forget it
			continue
		}

		m, ok := event.Data.(map[string]interface{})
		if child == "" && ok {
			// if we were given a map at the root then we have
			// to send an event per child
			for _, k := range sortedKeys(m) {
				v := m[k]
				node := sync.NewNode(k, v)
				db.Add(k, node)
				fn(newSnapshot(node), *prevKey)
				*prevKey = k
			}
			continue
		}

		// we have a single event to process
		node := sync.NewNode(child, event.Data)
		db.Add(strings.Trim(child, "/"), node)

		fn(newSnapshot(node), *prevKey)
		*prevKey = child
	}
	return nil
}

// ChildChanged listens on the firebase instance and executes the callback
// for every child that is changed.
//
// You cannot set the same function twice on a Firebase reference, if you do
// the first function will be overridden and you will not be able to close the
// connection.
func (fb *Firebase) ChildChanged(fn ChildEventFunc) error {
	return fb.addEventFunc(fn, fn.childChanged)
}

func (fn ChildEventFunc) childChanged(db *sync.Database, prevKey *string, notifications chan Event) error {
	first, ok := <-notifications
	if !ok {
		return errors.New("channel closed")
	}

	db.Add("", sync.NewNode("", first.Data))
	for event := range notifications {
		if event.Type == EventTypeError {
			err, ok := event.Data.(error)
			if !ok {
				err = fmt.Errorf("Got error from event %#v", event)
			}
			return err
		}

		path := strings.Trim(event.Path, "/")
		if event.Data == nil {
			db.Del(path)
			continue
		}

		child := strings.Split(path, "/")[0]
		node := sync.NewNode(child, event.Data)

		dbNode := db.Get("")
		if _, ok := dbNode.Child(child); child != "" && !ok {
			// if the child is new, ignore it.
			db.Add(path, node)
			continue
		}

		if m, ok := event.Data.(map[string]interface{}); child == "" && ok {
			// we've got children so send an event per child
			for _, k := range sortedKeys(m) {
				v := m[k]
				node := sync.NewNode(k, v)
				newPath := strings.TrimPrefix(child+"/"+k, "/")
				if _, ok := dbNode.Child(k); !ok {
					db.Add(newPath, node)
					continue
				}

				db.Update(newPath, node)
				fn(newSnapshot(node), *prevKey)
				*prevKey = k
			}
			continue
		}

		db.Update(path, node)
		fn(newSnapshot(db.Get(child)), *prevKey)
		*prevKey = child
	}
	return nil
}

// ChildRemoved listens on the firebase instance and executes the callback
// for every child that is deleted.
//
// You cannot set the same function twice on a Firebase reference, if you do
// the first function will be overridden and you will not be able to close the
// connection.
func (fb *Firebase) ChildRemoved(fn ChildEventFunc) error {
	return fb.addEventFunc(fn, fn.childRemoved)
}

func (fn ChildEventFunc) childRemoved(db *sync.Database, prevKey *string, notifications chan Event) error {
	first, ok := <-notifications
	if !ok {
		return errors.New("channel closed")
	}

	node := sync.NewNode("", first.Data)
	db.Add("", node)

	for event := range notifications {
		if event.Type == EventTypeError {
			err, ok := event.Data.(error)
			if !ok {
				err = fmt.Errorf("Got error from event %#v", event)
			}
			return err
		}

		path := strings.Trim(event.Path, "/")
		node := sync.NewNode(path, event.Data)

		if event.Type == EventTypePatch {
			db.Update(path, node)
			continue
		}

		if event.Data != nil {
			db.Add(path, node)
			continue
		}

		if path == "" {
			// if node that is being listened to is deleted,
			// an event should be triggered for every child
			children := db.Get("").Children
			orderedChildren := make([]string, len(children))
			var i int
			for k := range children {
				orderedChildren[i] = k
				i++
			}

			sort.Strings(orderedChildren)

			for _, k := range orderedChildren {
				node := db.Get(k)
				fn(newSnapshot(node), "")
				db.Del(k)
			}

			db.Del(path)
			continue
		}

		node = db.Get(path)
		fn(newSnapshot(node), "")
		db.Del(path)
	}
	return nil
}

type handleSSEFunc func(*sync.Database, *string, chan Event) error

func (fb *Firebase) addEventFunc(fn ChildEventFunc, handleSSE handleSSEFunc) error {
	fb.eventMtx.Lock()
	defer fb.eventMtx.Unlock()

	stop := make(chan struct{})
	key := fmt.Sprintf("%v", fn)
	if _, ok := fb.eventFuncs[key]; ok {
		return nil
	}

	fb.eventFuncs[key] = stop
	notifications, err := fb.watch(stop)
	if err != nil {
		return err
	}

	db := sync.NewDB()
	prevKey := new(string)
	var run func(notifications chan Event, backoff time.Duration)
	run = func(notifications chan Event, backoff time.Duration) {
		fb.eventMtx.Lock()
		if _, ok := fb.eventFuncs[key]; !ok {
			fb.eventMtx.Unlock()
			// the func has been removed
			return
		}
		fb.eventMtx.Unlock()

		if err := handleSSE(db, prevKey, notifications); err == nil {
			// we returned gracefully
			return
		}

		// give firebase some time
		backoff *= 2
		time.Sleep(backoff)

		// try and reconnect
		for notifications, err = fb.watch(stop); err != nil; time.Sleep(backoff) {
			fb.eventMtx.Lock()
			if _, ok := fb.eventFuncs[key]; !ok {
				fb.eventMtx.Unlock()
				// func has been removed
				return
			}
			fb.eventMtx.Unlock()
		}

		// give this another shot
		run(notifications, backoff)
	}

	go run(notifications, fb.watchHeartbeat)
	return nil
}

// RemoveEventFunc removes the given function from the firebase
// reference.
func (fb *Firebase) RemoveEventFunc(fn ChildEventFunc) {
	fb.eventMtx.Lock()
	defer fb.eventMtx.Unlock()

	key := fmt.Sprintf("%v", fn)
	stop, ok := fb.eventFuncs[key]
	if !ok {
		return
	}

	delete(fb.eventFuncs, key)
	close(stop)
}

func sortedKeys(m map[string]interface{}) []string {
	orderedKeys := make([]string, len(m))
	var i int
	for k := range m {
		orderedKeys[i] = k
		i++
	}

	sort.Strings(orderedKeys)
	return orderedKeys
}
