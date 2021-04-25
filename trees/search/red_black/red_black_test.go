package redblack_test

import (
	"fmt"
	"testing"

	redblack "github.com/shanehowearth/algorithms/trees/search/red_black"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	testcases := map[string]struct {
		input  []int
		output map[int]redblack.Node
		err    error
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
		"second difficult insertion": {
			input: []int{7, 18, 3, 10, 22, 8, 11, 26, 5, 6},
			output: map[int]redblack.Node{
				3:  redblack.Node{Key: 3, Colour: redblack.Black, Child: [2]*redblack.Node{nil, {Key: 6}}, Parent: &redblack.Node{Key: 7}},
				5:  redblack.Node{Key: 5, Colour: redblack.Black, Parent: &redblack.Node{Key: 6}},
				6:  redblack.Node{Key: 6, Colour: redblack.Red, Child: [2]*redblack.Node{{Key: 5}}, Parent: &redblack.Node{Key: 3}},
				7:  redblack.Node{Key: 7, Colour: redblack.Black, Child: [2]*redblack.Node{{Key: 3}, {Key: 18}}},
				8:  redblack.Node{Key: 8, Colour: redblack.Red, Parent: &redblack.Node{Key: 10}},
				10: redblack.Node{Key: 10, Colour: redblack.Black, Child: [2]*redblack.Node{{Key: 8}, {Key: 11}}, Parent: &redblack.Node{Key: 18}},
				11: redblack.Node{Key: 11, Colour: redblack.Red, Parent: &redblack.Node{Key: 10}},
				18: redblack.Node{Key: 18, Colour: redblack.Red, Child: [2]*redblack.Node{{Key: 10}, {Key: 22}}, Parent: &redblack.Node{Key: 7}},
				22: redblack.Node{Key: 22, Colour: redblack.Black, Child: [2]*redblack.Node{nil, {Key: 26}}, Parent: &redblack.Node{Key: 18}},
				26: redblack.Node{Key: 26, Colour: redblack.Red, Parent: &redblack.Node{Key: 22}},
			},
		},
		"duplicate input": {
			input: []int{7, 18, 3, 10, 22, 8, 11, 26},
			err:   fmt.Errorf("key already exists"),
		},
	}
	for name, tc := range testcases {

		t.Run(name, func(t *testing.T) {
			rt := redblack.Tree{}
			for i := range tc.input {
				err := rt.Insert(tc.input[i], nil)
				assert.Nil(t, err, "inserting %d generated error %v", tc.input[i], err)
			}
			if tc.err != nil {
				err := rt.Insert(tc.input[0], nil)
				assert.NotNil(t, err, "expected an error")
				assert.EqualError(t, tc.err, err.Error(), "")
				return
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
						assert.Equal(t, tc.output[node.Key].Child[d].Key, node.Child[d].Key, "node %d child keys were different for direction %d", node.Key, d)
					} else {
						assert.Nil(t, node.Child[d], "Child %d of %d should be nil", d, node.Key)
					}
				}
				assert.Equal(t, node.Colour, tc.output[node.Key].Colour, "%d colour was different", node.Key)
				if tc.output[node.Key].Parent != nil {
					assert.NotNil(t, node.Parent, "Parent expected for node %d, but missing", node.Key)
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

func TestSearch(t *testing.T) {
	testcases := map[string]struct {
		input  []int
		key    int
		output *redblack.Node
		err    error
	}{
		"simple example exercises cmp.Child[Left]": {
			input:  []int{7, 18, 3, 10, 22, 8, 11, 26},
			key:    3,
			output: &redblack.Node{Key: 3},
		},
		"simple example exercises cmp.Child[Right]": {
			input:  []int{7, 18, 3, 10, 22, 8, 11, 26},
			key:    22,
			output: &redblack.Node{Key: 22},
		},
		"simple example exercises both possible paths": {
			input:  []int{7, 18, 3, 10, 22, 8, 11, 26},
			key:    11,
			output: &redblack.Node{Key: 11},
		},
		"simple example exercises gets root": {
			input:  []int{7, 18, 3, 10, 22, 8, 11, 26},
			key:    7,
			output: &redblack.Node{Key: 7},
		},
		"no root": {
			err: fmt.Errorf("no nodes in the tree"),
		},
		"simple example with error key < cmp.Key": {
			input: []int{7, 18, 3, 10, 22, 8, 11, 26},
			key:   23,
			err:   fmt.Errorf("%d not found", 23),
		},
		"simple example with error key > cmp.Key": {
			input: []int{7, 18, 3, 10, 22, 8, 11, 26},
			key:   27,
			err:   fmt.Errorf("%d not found", 27),
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			rt := redblack.Tree{}
			for i := range tc.input {
				err := rt.Insert(tc.input[i], nil)
				assert.Nil(t, err, "inserting %d generated error %v", tc.input[i], err)
			}
			node, err := rt.Search(tc.key)
			if tc.err == nil {
				assert.Nil(t, err, "got unexpected error searching for %d", tc.key)
				assert.NotNil(t, node, "was expecting a node when searching for %d", tc.key)
				assert.Equal(t, node.Key, tc.output.Key)
			} else {
				assert.NotNil(t, err, "was expecting an error searching for %d", tc.key)
				assert.Nil(t, node, "was not expecting a node when searching for %d", tc.key)
			}
		})
	}
}
