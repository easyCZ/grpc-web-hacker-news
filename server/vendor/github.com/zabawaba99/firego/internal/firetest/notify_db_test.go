package firetest

import (
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zabawaba99/firego/sync"
)

func TestNotifyDBAdd(t *testing.T) {
	for _, test := range []struct {
		path string
		node *sync.Node
	}{
		{
			path: "scalars/string",
			node: sync.NewNode("", "foo"),
		},
		{
			path: "s/c/a/l/a/r/s/s/t/r/i/n/g",
			node: sync.NewNode("", []interface{}{"foo", "bar"}),
		},
	} {
		db := newNotifyDB()

		// listen for notifications
		notifications := db.watch("")
		exited := make(chan struct{})
		go func() {
			n, ok := <-notifications
			assert.True(t, ok)
			assert.Equal(t, "put", n.Name)
			assert.Equal(t, test.path, n.Data.Path, "wat?")
			assert.Equal(t, test.node, n.Data.Data)
			close(exited)
		}()

		db.add(test.path, test.node)

		select {
		case <-exited:
		case <-time.After(250 * time.Millisecond):
		}
		db.stopWatching(test.path, notifications)
	}
}

func TestNotifyDBDel(t *testing.T) {
	existingNodes := []string{
		"root/only/two",
		"root/only/three",
		"root/only/one/child/here",
	}
	db := newNotifyDB()
	for _, p := range existingNodes {
		db.add(p, sync.NewNode("", 1))
	}

	// listen for notifications
	notifications := db.watch("")
	exited := make(chan struct{})
	go func() {
		regex := regexp.MustCompile("(root/only/one/child|root)")
		n, ok := <-notifications
		assert.True(t, ok)
		assert.Equal(t, "put", n.Name)
		assert.Regexp(t, regex, n.Data.Path)

		n, ok = <-notifications
		assert.True(t, ok)
		assert.Equal(t, "put", n.Name)
		assert.Regexp(t, regex, n.Data.Path)
		close(exited)
	}()

	db.del("root/only/one/child")
	db.del("root")

	select {
	case <-exited:
	case <-time.After(250 * time.Millisecond):
	}
	db.stopWatching("", notifications)
}
