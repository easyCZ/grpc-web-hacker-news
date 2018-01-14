package sync

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	for _, test := range []struct {
		path string
		node *Node
	}{
		{
			path: "scalars/string",
			node: NewNode("", "foo"),
		},
		{
			path: "s/c/a/l/a/r/s/s/t/r/i/n/g",
			node: NewNode("", []interface{}{"foo", "bar"}),
		},
	} {
		db := NewDB()

		db.Add(test.path, test.node)

		rabbitHole := strings.Split(test.path, "/")
		previous := db.root
		for i := 0; i < len(rabbitHole); i++ {
			var ok bool
			previous, ok = previous.Children[rabbitHole[i]]
			assert.True(t, ok, test.path)
		}

		assert.NoError(t, equalNodes(test.node, previous), test.path)
	}
}

func TestUpdate(t *testing.T) {
	db := NewDB()

	db.Add("", NewNode("", map[string]interface{}{
		"hello": map[string]interface{}{
			"world":  2,
			"world2": 1,
		},
	}))

	db.Update("hello/world", NewNode("", "hi"))
	db.Update("hello/world3", NewNode("", "lol"))

	assert.Equal(t, "hi", db.Get("hello/world").Value)
	assert.Equal(t, 1, db.Get("hello/world2").Value)
	assert.Equal(t, "lol", db.Get("hello/world3").Value)
}

func TestGet(t *testing.T) {
	for _, test := range []struct {
		path string
		node *Node
	}{
		{
			path: "scalars/string",
			node: NewNode("", "foo"),
		},
		{
			path: "s/c/a/l/a/r/s/s/t/r/i/n/g",
			node: NewNode("", []interface{}{"foo", "bar"}),
		},
	} {
		db := NewDB()
		db.Add(test.path, test.node)

		assert.NoError(t, equalNodes(test.node, db.Get(test.path)), test.path)
	}
}

func TestDel(t *testing.T) {
	existingNodes := []string{
		"root/only/two",
		"root/only/three",
		"root/only/one/child/here",
	}
	db := NewDB()
	for _, p := range existingNodes {
		db.Add(p, NewNode("", 1))
	}

	db.Del("root/only/one/child")
	assert.Nil(t, db.Get("root/only/one/child/here"))
	assert.Nil(t, db.Get("root/only/one/child"))
	assert.Nil(t, db.Get("root/only/one"))

	n := db.Get("root/only")
	require.NotNil(t, n)

	assert.Len(t, n.Children, 2)
	_, ok := n.Children["one"]
	assert.False(t, ok)

	db.Del("root")
	n = db.Get("")
	require.NotNil(t, n)
	assert.Len(t, n.Children, 0)
}
