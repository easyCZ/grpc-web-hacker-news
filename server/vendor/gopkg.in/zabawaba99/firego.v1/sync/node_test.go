package sync

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestNodeWithKids(children map[string]*Node) *Node {
	n := &Node{}
	for _, child := range children {
		child.Parent = n
	}
	n.Children = children
	return n
}

func equalNodes(expected, actual *Node) error {
	if ec, ac := len(expected.Children), len(actual.Children); ec != ac {
		return fmt.Errorf("Children count is not the same\n\tExpected: %d\n\tActual: %d", ec, ac)
	}

	if len(expected.Children) == 0 {
		if !assert.ObjectsAreEqualValues(expected.Value, actual.Value) {
			return fmt.Errorf("Node values not equal\n\tExpected: %T %v\n\tActual: %T %v", expected.Value, expected.Value, actual.Value, actual.Value)
		}
		return nil
	}

	for child, n := range expected.Children {
		n2, ok := actual.Children[child]
		if !ok {
			return fmt.Errorf("Expected node to have child: %s", child)
		}

		err := equalNodes(n, n2)
		if err != nil {
			return err
		}
	}
	return nil
}

func TestNewNode(t *testing.T) {

	for _, test := range []struct {
		name string
		node *Node
	}{
		{
			name: "scalars/string",
			node: NewNode("", "foo"),
		},
		{
			name: "scalars/number",
			node: NewNode("", 2),
		},
		{
			name: "scalars/decimal",
			node: NewNode("", 2.2),
		},
		{
			name: "scalars/boolean",
			node: NewNode("", false),
		},
		{
			name: "arrays/strings",
			node: NewNode("", []interface{}{"foo", "bar"}),
		},
		{
			name: "arrays/booleans",
			node: NewNode("", []interface{}{true, false}),
		},
		{
			name: "arrays/numbers",
			node: NewNode("", []interface{}{1, 2, 3}),
		},
		{
			name: "arrays/decimals",
			node: NewNode("", []interface{}{1.1, 2.2, 3.3}),
		},
		{
			name: "objects/simple",
			node: newTestNodeWithKids(map[string]*Node{
				"foo": NewNode("", "bar"),
			}),
		},
		{
			name: "objects/complex",
			node: newTestNodeWithKids(map[string]*Node{
				"foo":  NewNode("", "bar"),
				"foo1": NewNode("", 2),
				"foo2": NewNode("", true),
				"foo3": NewNode("", 3.42),
			}),
		},
		{
			name: "objects/nested",
			node: newTestNodeWithKids(map[string]*Node{
				"dinosaurs": newTestNodeWithKids(map[string]*Node{
					"bruhathkayosaurus": newTestNodeWithKids(map[string]*Node{
						"appeared": NewNode("", -70000000),
						"height":   NewNode("", 25),
						"length":   NewNode("", 44),
						"order":    NewNode("", "saurischia"),
						"vanished": NewNode("", -70000000),
						"weight":   NewNode("", 135000),
					}),
					"lambeosaurus": newTestNodeWithKids(map[string]*Node{
						"appeared": NewNode("", -76000000),
						"height":   NewNode("", 2.1),
						"length":   NewNode("", 12.5),
						"order":    NewNode("", "ornithischia"),
						"vanished": NewNode("", -75000000),
						"weight":   NewNode("", 5000),
					}),
				}),
				"scores": newTestNodeWithKids(map[string]*Node{
					"bruhathkayosaurus": NewNode("", 55),
					"lambeosaurus":      NewNode("", 21),
				}),
			}),
		},
		{
			name: "objects/with_arrays",
			node: newTestNodeWithKids(map[string]*Node{
				"regular":  NewNode("", "item"),
				"booleans": NewNode("", []interface{}{false, true}),
				"numbers":  NewNode("", []interface{}{1, 2}),
				"decimals": NewNode("", []interface{}{1.1, 2.2}),
				"strings":  NewNode("", []interface{}{"foo", "bar"}),
			}),
		},
	} {
		data, err := ioutil.ReadFile("fixtures/" + test.name + ".json")
		require.NoError(t, err, test.name)

		var v interface{}
		require.NoError(t, json.Unmarshal(data, &v), test.name)

		n := NewNode("", v)
		assert.NoError(t, equalNodes(test.node, n), test.name)
	}
}

func TestObjectify(t *testing.T) {
	for _, test := range []struct {
		name   string
		object interface{}
	}{
		{
			name: "nil",
		},
		{
			name:   "string",
			object: "foo",
		},
		{
			name:   "number",
			object: 2,
		},
		{
			name:   "decimal",
			object: 2.2,
		},
		{
			name:   "boolean",
			object: false,
		},
		{
			name:   "arrays",
			object: []interface{}{"foo", 2, 2.2, false},
		},
		{
			name: "object",
			object: map[string]interface{}{
				"one_fish":     "two_fish",
				"red_fish":     2.2,
				"netflix_list": []interface{}{"Orange is the New Black", "House of Cards"},
				"shopping_list": map[string]interface{}{
					"publix":  "milk",
					"walmart": "reese's pieces",
				},
			},
		},
	} {
		node := NewNode("", test.object)
		assert.Equal(t, test.object, node.Objectify())
	}
}

func TestChild(t *testing.T) {
	node := NewNode("", map[string]interface{}{
		"one": map[string]interface{}{
			"two": map[string]interface{}{
				"three": true,
			},
		},
	})

	one, ok := node.Child("one")
	require.True(t, ok)
	require.NotNil(t, one)

	two, ok := one.Child("two")
	require.True(t, ok)
	require.NotNil(t, two)

	three, ok := two.Child("three")
	require.True(t, ok)
	require.NotNil(t, three)

	threeNested, ok := node.Child("one/two/three")
	require.True(t, ok)
	assert.Equal(t, three, threeNested)

	one, ok = node.Child("nope")
	assert.False(t, ok)
	assert.Nil(t, one)
}

func TestMerge(t *testing.T) {
	base := &Node{Children: map[string]*Node{
		"foo":        NewNode("", "bar"),
		"notreplace": NewNode("", "yes"),
	}}

	newNode := &Node{Children: map[string]*Node{
		"foo": NewNode("", "troll"),
		"new": NewNode("", "lala"),
	}}

	base.merge(newNode)

	expected := NewNode("", map[string]string{
		"foo":        "troll",
		"notreplace": "yes",
		"new":        "lala",
	})

	err := equalNodes(expected, base)
	assert.NoError(t, err)
}

func TestPrune(t *testing.T) {
	/*
		Children:	0
		Value:		Non nil
		Parent: 	nil
	*/
	n := NewNode("", "foo")
	assert.Nil(t, n.prune())

	/*
		Children:	0
		Value:		Non nil
		Parent: 	Non nil
	*/
	n = NewNode("", "foo")
	n.Parent = NewNode("", 1)
	assert.Nil(t, n.prune())

	/*
		Children:	0
		Value:		nil
		Parent: 	Non nil
	*/
	n = &Node{}
	parent := newTestNodeWithKids(map[string]*Node{"foo": n})
	parentFromPrune := n.prune()

	assert.NotNil(t, parentFromPrune)
	assert.Equal(t, parent, parentFromPrune)
	assert.Nil(t, n.Parent)
	assert.Nil(t, n.Children)

	/*
		Children:	1
		Value:		nil
		Parent: 	Non nil
	*/
	n = newTestNodeWithKids(map[string]*Node{"c1": n})
	parent = newTestNodeWithKids(map[string]*Node{"foo": n})
	assert.Nil(t, n.prune())

	/*
		Children:	1
		Value:		nil
		Parent: 	nil
	*/
	n = newTestNodeWithKids(map[string]*Node{"c1": n})
	assert.Nil(t, n.prune())

	/*
		Children:	2
		Value:		nil
		Parent: 	Non nil
	*/
	n = newTestNodeWithKids(map[string]*Node{
		"c1": NewNode("", 1),
		"c2": NewNode("", 2),
	})
	n.Parent = NewNode("", "hello!")
	assert.Nil(t, n.prune())
}
