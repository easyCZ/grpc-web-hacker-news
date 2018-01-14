package firego

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zabawaba99/firego/sync"
)

func TestDataSnapshotKey(t *testing.T) {
	n := sync.Node{Key: "foo"}
	d := newSnapshot(&n)

	assert.Equal(t, n.Key, d.Key)
}

func TestDataSnapshotValue(t *testing.T) {
	n := sync.Node{Value: "foo"}
	d := newSnapshot(&n)

	assert.Equal(t, n.Value, d.Value)
}

func TestDataSnapshotChild(t *testing.T) {
	n := sync.NewNode("", map[string]interface{}{
		"one": map[string]interface{}{
			"two": map[string]interface{}{
				"three": true,
			},
		},
	})
	d := newSnapshot(n)

	one, ok := d.Child("one")
	require.True(t, ok)
	assert.Equal(t, one.Key, "one")
	assert.Equal(t, one.Value, map[string]interface{}{
		"two": map[string]interface{}{
			"three": true,
		},
	})

	two, ok := one.Child("two")
	require.True(t, ok)
	assert.Equal(t, two.Key, "two")
	assert.Equal(t, two.Value, map[string]interface{}{
		"three": true,
	})

	three, ok := two.Child("three")
	require.True(t, ok)
	assert.Equal(t, three.Key, "three")
	assert.Equal(t, three.Value, true)

	three2, ok := d.Child("one/two/three")
	require.True(t, ok)
	assert.Equal(t, three, three2)
}
