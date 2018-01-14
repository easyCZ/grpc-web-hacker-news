package firego

import (
	"strings"

	"github.com/zabawaba99/firego/sync"
)

// DataSnapshot instances contains data from a Firebase reference.
type DataSnapshot struct {
	// Key retrieves the key for the source location of this snapshot
	Key string

	// Value retrieves the data contained in this snapshot.
	Value interface{}
}

func newSnapshot(node *sync.Node) DataSnapshot {
	return DataSnapshot{
		Key:   node.Key,
		Value: node.Objectify(),
	}
}

// Child gets a DataSnapshot for the location at the specified relative path.
// The relative path can either be a simple child key (e.g. 'fred') or a deeper
// slash-separated path (e.g. 'fred/name/first').
func (d *DataSnapshot) Child(name string) (DataSnapshot, bool) {
	name = strings.Trim(name, "/")
	rabbitHole := strings.Split(name, "/")

	current := *d
	for _, tkn := range rabbitHole {
		children, ok := current.Value.(map[string]interface{})
		if !ok {
			return current, false
		}

		v, ok := children[tkn]
		if !ok {
			return current, ok
		}

		current = DataSnapshot{
			Key:   tkn,
			Value: v,
		}
	}

	return current, true
}
