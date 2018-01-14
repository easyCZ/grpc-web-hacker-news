package firetest

import (
	"encoding/json"
	"strings"
	_sync "sync"
	"time"

	"github.com/zabawaba99/firego/sync"
)

type event struct {
	Name string
	Data eventData
}

type eventData struct {
	Path string     `json:"path"`
	Data *sync.Node `json:"data"`
}

func (ed eventData) MarshalJSON() ([]byte, error) {
	type eventData2 eventData
	ed2 := eventData2(ed)
	ed2.Path = "/" + ed2.Path
	return json.Marshal(ed2)
}

func newEvent(name, path string, n *sync.Node) event {
	return event{
		Name: "put",
		Data: eventData{
			Path: path,
			Data: n,
		},
	}
}

type notifyDB struct {
	intDB *sync.Database

	watchersMtx _sync.RWMutex
	watchers    map[string][]chan event
}

func newNotifyDB() *notifyDB {
	return &notifyDB{
		intDB:    sync.NewDB(),
		watchers: map[string][]chan event{},
	}
}

func (db *notifyDB) add(path string, n *sync.Node) {
	db.intDB.Add(path, n)
	go db.notify(newEvent("put", path, n))
}

func (db *notifyDB) update(path string, n *sync.Node) {
	db.intDB.Update(path, n)
	go db.notify(newEvent("patch", path, n))
}

func (db *notifyDB) del(path string) {
	db.intDB.Del(path)
	go db.notify(newEvent("put", path, nil))
}

func (db *notifyDB) get(path string) *sync.Node {
	return db.intDB.Get(path)
}

func (db *notifyDB) notify(e event) {
	db.watchersMtx.RLock()
	for path, listeners := range db.watchers {
		if !strings.HasPrefix(e.Data.Path, path) {
			continue
		}

		// Make sure to not return full path when notifying
		// only return the path relative to the watcher
		e.Data.Path = strings.TrimPrefix(e.Data.Path, path)
		e.Data.Path = sanitizePath(e.Data.Path)

		for _, c := range listeners {
			select {
			case c <- e:
			case <-time.After(250 * time.Millisecond):
				continue
			}
		}
	}
	db.watchersMtx.RUnlock()
}

func (db *notifyDB) stopWatching(path string, c chan event) {
	db.watchersMtx.Lock()
	index := -1
	for i, ch := range db.watchers[path] {
		if ch == c {
			index = i
			break
		}
	}

	if index > -1 {
		a := db.watchers[path]
		db.watchers[path] = append(a[:index], a[index+1:]...)
		close(c)
	}
	db.watchersMtx.Unlock()
}

func (db *notifyDB) watch(path string) chan event {
	c := make(chan event)

	db.watchersMtx.Lock()
	db.watchers[path] = append(db.watchers[path], c)
	db.watchersMtx.Unlock()

	return c
}
