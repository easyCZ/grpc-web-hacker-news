package firetest

import (
	"strings"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zabawaba99/firego/sync"
)

func TestRequireAuth(t *testing.T) {
	for _, require := range []bool{true, false} {
		ft := New()
		ft.RequireAuth(require)
		var expected int32
		if require {
			expected = 1
		}
		assert.Equal(t, expected, atomic.LoadInt32(ft.requireAuth))
	}
}

func TestCreate(t *testing.T) {
	var (
		ft = New()
		v  = true
	)

	for _, p := range []string{"path/hi", ""} {
		name := ft.Create(p, v)
		assert.True(t, strings.HasPrefix(name, "~"), "name is missing `~` prefix")

		n := ft.db.get(sanitizePath(p + "/" + name))
		assert.Equal(t, v, n.Value)
	}
}

func TestDelete(t *testing.T) {
	var (
		ft   = New()
		path = "foo/bar"
		v    = true
	)

	// delete path directly
	ft.db.add(path, sync.NewNode("", v))
	ft.Delete(path)
	assert.Nil(t, ft.db.get(path))

	// delete parent
	ft.db.add(path, sync.NewNode("", v))
	ft.Delete("foo")
	assert.Nil(t, ft.db.get(path))
}

func TestUpdate(t *testing.T) {
	var (
		ft   = New()
		path = "foo/bar"
		v    = map[string]interface{}{
			"1": "one",
			"2": "two",
			"3": "three",
		}
	)
	ft.db.add(path, sync.NewNode("", v))

	ft.Update(path, map[string]interface{}{
		"1": "three",
		"3": "one",
	})

	one := ft.db.get(path + "/1")
	three := ft.db.get(path + "/3")
	assert.Equal(t, "three", one.Value)
	assert.Equal(t, "one", three.Value)
}

func TestUpdateNil(t *testing.T) {
	var (
		ft   = New()
		path = "foo/bar"
		v    = map[string]string{
			"1": "one",
			"2": "two",
			"3": "three",
		}
	)
	ft.db.add(path, sync.NewNode("", v))

	ft.Update(path, nil)
	assert.Nil(t, ft.db.get(path))
	assert.Nil(t, ft.db.get(path+"/1"))
	assert.Nil(t, ft.db.get(path+"/2"))
	assert.Nil(t, ft.db.get(path+"/3"))
}

func TestSet(t *testing.T) {
	var (
		ft   = New()
		path = "foo/bar"
		v    = true
	)
	ft.Set(path, v)

	n := ft.db.get(path)
	assert.Equal(t, v, n.Value)
}

func TestGet(t *testing.T) {
	var (
		ft   = New()
		path = "foo/bar"
		v    = true
	)
	ft.db.add(path, sync.NewNode("", v))

	val := ft.Get(path)
	assert.Equal(t, v, val)
}
