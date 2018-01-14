package sync

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

// Node represents an object linked in Database. This object
// should not be created by hand, use NewNode when creating
// a new instance of Node.
type Node struct {
	mtx      sync.RWMutex
	Key      string
	Value    interface{}
	Children map[string]*Node

	Parent    *Node
	sliceKids bool
}

// NewNode converts the given data into a node.
func NewNode(key string, data interface{}) *Node {
	n := &Node{
		Key: key,
	}

	n.mtx.Lock()
	n.Children = map[string]*Node{}
	n.mtx.Unlock()

	if data == nil {
		return n
	}

	switch val := reflect.ValueOf(data); val.Kind() {
	case reflect.Map:
		for _, k := range val.MapKeys() {
			v := val.MapIndex(k)
			key := fmt.Sprintf("%s", k.Interface())

			child := NewNode(key, v.Interface())
			child.Parent = n
			n.Children[key] = child
		}

	case reflect.Array, reflect.Slice:
		n.sliceKids = true

		for i := 0; i < val.Len(); i++ {
			v := val.Index(i)
			key := strconv.FormatInt(int64(i), 10)

			child := NewNode(key, v.Interface())
			child.Parent = n
			n.Children[key] = child
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fallthrough
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fallthrough
	case reflect.Float32, reflect.Float64:
		fallthrough
	case reflect.String, reflect.Bool, reflect.Interface:
		n.Value = val.Interface()
	default:
		fmt.Printf("Unsupported type %s(%#v)If you see this log please report an issue on https://github.com/zabawaba99/firego", data, data)
	}

	return n
}

// MarshalJSON turns the node object into JSON bytes.
func (n *Node) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Objectify())
}

// Objectify turns the node and all its children into a go type.
// If a node was created from a slice initially, a slice will be return.
// If a node has child nodes, a map will be returned.
// Otherwise, a primitive type will be returned.
func (n *Node) Objectify() interface{} {
	n.mtx.RLock()
	defer n.mtx.RUnlock()

	if n.isNil() {
		return nil
	}

	if n.Value != nil {
		return n.Value
	}

	if n.sliceKids {
		obj := make([]interface{}, len(n.Children))
		for k, v := range n.Children {
			index, err := strconv.Atoi(k)
			if err != nil {
				continue
			}
			obj[index] = v.Objectify()
		}
		return obj
	}

	obj := map[string]interface{}{}
	for k, v := range n.Children {
		obj[k] = v.Objectify()
	}

	return obj
}

// Child gets a DataSnapshot for the location at the specified relative path.
// The relative path can either be a simple child key (e.g. 'fred') or a deeper
// slash-separated path (e.g. 'fred/name/first').
func (n *Node) Child(name string) (*Node, bool) {
	n.mtx.RLock()
	defer n.mtx.RUnlock()

	name = strings.Trim(name, "/")
	rabbitHole := strings.Split(name, "/")

	current := n
	for i := 0; i < len(rabbitHole); i++ {
		next, ok := current.Children[rabbitHole[i]]
		if !ok {
			// item does not exist, no need to do anything
			return nil, false
		}

		current = next
	}
	return current, true
}

func (n *Node) isNil() bool {
	n.mtx.RLock()
	defer n.mtx.RUnlock()

	return n.Value == nil && len(n.Children) == 0
}

func (n *Node) merge(newNode *Node) {
	n.mtx.Lock()
	defer n.mtx.Unlock()

	for k, v := range newNode.Children {
		n.Children[k] = v
	}
	n.Value = newNode.Value
}

func (n *Node) prune() *Node {
	n.mtx.Lock()
	defer n.mtx.Unlock()

	if len(n.Children) > 0 || n.Value != nil {
		return nil
	}

	parent := n.Parent
	n.Parent = nil
	return parent
}
