package redblack_test

import (
	"testing"

	redblack "github.com/shanehowearth/algorithms/trees/search/red_black"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	testcases := map[string]struct {
		input  []int
		output map[int]redblack.Node
	}{
		"class example": {
			input: []int{7, 18, 3, 10, 22, 8, 11, 26},
			output: map[int]redblack.Node{
				7:  redblack.Node{Key: 7, Colour: redblack.Black, Child: [2]*redblack.Node{{Key: 3}, {Key: 18}}},
				3:  redblack.Node{Key: 3, Colour: redblack.Black, Parent: &redblack.Node{Key: 7}},
				18: redblack.Node{Key: 18, Colour: redblack.Red, Child: [2]*redblack.Node{{Key: 10}, {Key: 22}}, Parent: &redblack.Node{Key: 7}},
				10: redblack.Node{Key: 10, Colour: redblack.Black, Child: [2]*redblack.Node{{Key: 8}, {Key: 11}}, Parent: &redblack.Node{Key: 18}},
				22: redblack.Node{Key: 22, Colour: redblack.Black, Child: [2]*redblack.Node{nil, {Key: 26}}, Parent: &redblack.Node{Key: 18}},
				8:  redblack.Node{Key: 8, Colour: redblack.Red, Parent: &redblack.Node{Key: 10}},
				11: redblack.Node{Key: 11, Colour: redblack.Red, Parent: &redblack.Node{Key: 10}},
				26: redblack.Node{Key: 26, Colour: redblack.Red, Parent: &redblack.Node{Key: 22}},
			},
		},
		"difficult insertion": {
			input: []int{7, 18, 3, 10, 22, 8, 11, 26, 15},
			output: map[int]redblack.Node{
				7:  redblack.Node{Key: 7, Colour: redblack.Red, Child: [2]*redblack.Node{{Key: 3}, {Key: 8}}, Parent: &redblack.Node{Key: 10}},
				3:  redblack.Node{Key: 3, Colour: redblack.Black, Parent: &redblack.Node{Key: 7}},
				18: redblack.Node{Key: 18, Colour: redblack.Red, Child: [2]*redblack.Node{{Key: 11}, {Key: 22}}, Parent: &redblack.Node{Key: 10}},
				10: redblack.Node{Key: 10, Colour: redblack.Black, Child: [2]*redblack.Node{{Key: 7}, {Key: 18}}},
				22: redblack.Node{Key: 22, Colour: redblack.Black, Child: [2]*redblack.Node{nil, {Key: 26}}, Parent: &redblack.Node{Key: 18}},
				8:  redblack.Node{Key: 8, Colour: redblack.Black, Parent: &redblack.Node{Key: 7}},
				11: redblack.Node{Key: 11, Colour: redblack.Black, Child: [2]*redblack.Node{nil, {Key: 15}}, Parent: &redblack.Node{Key: 18}},
				15: redblack.Node{Key: 15, Colour: redblack.Red, Parent: &redblack.Node{Key: 11}},
				26: redblack.Node{Key: 26, Colour: redblack.Red, Parent: &redblack.Node{Key: 22}},
			},
		},
	}
	for name, tc := range testcases {

		t.Run(name, func(t *testing.T) {
			rt := redblack.Tree{}
			for i := range tc.input {
				rt.Insert(tc.input[i], nil)
			}

			nodes := []*redblack.Node{rt.Root}
			count := 0
			for len(nodes) != 0 {
				count++
				node := nodes[0]
				nodes = nodes[1:]
				assert.Equal(t, tc.output[node.Key].Key, node.Key, "keys were different")
				for d := range []redblack.Direction{redblack.Left, redblack.Right} {
					if tc.output[node.Key].Child[d] != nil {
						assert.NotNil(t, node.Child[d], "expected the node %d to have a child[%d]", node.Key, d)
						assert.NotNil(t, tc.output[node.Key].Child[d].Key)
						assert.NotNil(t, node.Child[d].Key)
						assert.Equal(t, tc.output[node.Key].Child[d].Key, node.Child[d].Key, "keys were different for %d", d)
					} else {
						assert.Nil(t, node.Child[d], "Child %d of %d should be nil", d, node.Key)
					}
				}
				assert.Equal(t, node.Colour, tc.output[node.Key].Colour, "%d colour was different", node.Key)
				if tc.output[node.Key].Parent != nil {
					assert.NotNil(t, node.Parent, "Parent expected, but missing")
					assert.Equal(t, tc.output[node.Key].Parent.Key, node.Parent.Key, "parent keys for %d were different", node.Key)
				} else {
					assert.Nil(t, node.Parent, "Parent of %d should be nil", node.Key)
				}
				if node.Child[redblack.Left] != nil {
					nodes = append(nodes, node.Child[redblack.Left])
				}
				if node.Child[redblack.Right] != nil {
					nodes = append(nodes, node.Child[redblack.Right])
				}
			}
			assert.Equal(t, len(tc.output), count, "Counts do not match")
		})
	}
}
