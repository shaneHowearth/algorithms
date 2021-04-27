package redblack_test

import (
	"fmt"
	"testing"

	redblack "github.com/shanehowearth/algorithms/trees/search/red_black"

	"github.com/stretchr/testify/assert"
)

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
				assert.Equal(t, tc.output.Key, node.Key, "expected %d but got %d in search", tc.output.Key, node.Key)
			} else {
				assert.NotNil(t, err, "was expecting an error searching for %d", tc.key)
				assert.Nil(t, node, "was not expecting a node when searching for %d", tc.key)
			}
		})
	}
}

func TestGetMin(t *testing.T) {
	testcases := map[string]struct {
		input  []int
		output *redblack.Node
		err    error
	}{
		"negative min": {
			input:  []int{7, 18, 3, 10, -17, 22, 8, 11, 26},
			output: &redblack.Node{Key: -17},
		},
		"simple example": {
			input:  []int{7, 18, 3, 10, 22, 8, 11, 26},
			output: &redblack.Node{Key: 3},
		},
		"no root node": {
			err: fmt.Errorf("no nodes in the tree"),
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			rt := redblack.Tree{}
			for i := range tc.input {
				err := rt.Insert(tc.input[i], nil)
				assert.Nil(t, err, "inserting %d generated error %v", tc.input[i], err)
			}

			node, err := rt.GetMin()
			if tc.err == nil {
				assert.Nil(t, err, "was not expecting an error %v", err)
				assert.NotNil(t, node, "was expecting a node")
				assert.Equal(t, tc.output.Key, node.Key, "wanted %d, got %d", tc.output.Key, node.Key)
			} else {
				assert.NotNil(t, err, "was expecting an error")
				assert.EqualError(t, tc.err, err.Error(), "wanted %s, got %s", tc.err.Error(), err.Error())
			}
		})
	}
}

func TestGetMax(t *testing.T) {
	testcases := map[string]struct {
		input  []int
		output *redblack.Node
		err    error
	}{
		"simple example": {
			input:  []int{7, 18, 3, 10, 22, 8, 11, 26},
			output: &redblack.Node{Key: 26},
		},
		"no root node": {
			err: fmt.Errorf("no nodes in the tree"),
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			rt := redblack.Tree{}
			for i := range tc.input {
				err := rt.Insert(tc.input[i], nil)
				assert.Nil(t, err, "inserting %d generated error %v", tc.input[i], err)
			}

			node, err := rt.GetMax()
			if tc.err == nil {
				assert.Nil(t, err, "was not expecting an error %v", err)
				assert.NotNil(t, node, "was expecting a node")
				assert.Equal(t, tc.output.Key, node.Key, "wanted %d, got %d", tc.output.Key, node.Key)
			} else {
				assert.NotNil(t, err, "was expecting an error")
				assert.EqualError(t, tc.err, err.Error(), "wanted %s, got %s", tc.err.Error(), err.Error())
			}

		})
	}
}

func TestGetPredecessor(t *testing.T) {
	testcases := map[string]struct {
		input  []int
		output *redblack.Node
		base   int // key of the node that we want the predecessor of
		err    error
	}{
		"no node supplied": {
			err: fmt.Errorf("no node supplied"),
		},
		"simple example": {
			input:  []int{7, 18, 3, 10, 22, 8, 11, 26},
			output: &redblack.Node{Key: 26},
			base:   18,
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			rt := redblack.Tree{}
			for i := range tc.input {
				err := rt.Insert(tc.input[i], nil)
				assert.Nil(t, err, "inserting %d generated error %v", tc.input[i], err)
			}

			var node *redblack.Node
			var err error
			if tc.err == nil {
				node, err = rt.Search(tc.base)
				assert.Nil(t, err, "searching for %d generated error %v", tc.base, err)
			}

			var output *redblack.Node
			output, err = rt.GetPredecessor(node)

			if tc.err == nil {
				assert.Nil(t, err, "searching for predecessor of %d generated an unexpected error %v", tc.base, err)
				assert.NotNil(t, output, "predecessor of %d should not be nil", tc.base)
				assert.Equal(t, output.Key, tc.output.Key, "got %d when %d was expected", output.Key, tc.output.Key)
			} else {
				assert.NotNil(t, err, "expected an error when searching for predecessor of %d", tc.base)
				assert.EqualError(t, err, tc.err.Error(), "wanted %s, got %s", tc.err.Error(), err.Error())
			}
		})
	}
}

func TestGetSuccessor(t *testing.T) {
	testcases := map[string]struct {
		input  []int
		output *redblack.Node
		base   int // key of the node that we want the successor of
		err    error
	}{
		"no node supplied": {
			err: fmt.Errorf("no node supplied"),
		},
		"simple example": {
			input:  []int{7, 18, 3, 10, 22, 8, 11, 26, 15},
			output: &redblack.Node{Key: 11},
			base:   18,
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			rt := redblack.Tree{}
			for i := range tc.input {
				err := rt.Insert(tc.input[i], nil)
				assert.Nil(t, err, "inserting %d generated error %v", tc.input[i], err)
			}

			var node *redblack.Node
			var err error
			if tc.err == nil {
				node, err = rt.Search(tc.base)
				assert.Nil(t, err, "searching for %d generated error %v", tc.base, err)
			}

			var output *redblack.Node
			output, err = rt.GetSuccessor(node)

			if tc.err == nil {
				assert.Nil(t, err, "searching for successor of %d generated an unexpected error %v", tc.base, err)
				assert.NotNil(t, output, "successor of %d should not be nil", tc.base)
				assert.Equal(t, tc.output.Key, output.Key, "got %d when %d was expected", output.Key, tc.output.Key)
			} else {
				assert.NotNil(t, err, "expected an error when searching for successor of %d", tc.base)
				assert.EqualError(t, err, tc.err.Error(), "wanted %s, got %s", tc.err.Error(), err.Error())
			}
		})
	}
}
