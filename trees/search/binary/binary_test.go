package binary_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/shanehowearth/algorithms/trees/search/binary"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	bt := binary.Tree{}
	for _, v := range []int{88, 65, 92, 54, 80, 76, 82, 81} {
		err := bt.Insert(v, fmt.Sprintf("Test data %d", v))
		assert.Nil(t, err, "error inserting %d into %#v", v, bt)
	}
	testcases := map[string]struct {
		input int
		key   int
		value string
		err   error
	}{
		"Root": {
			input: 88,
			key:   88,
			value: "Test data 88",
		},
		"Non-existant": {
			input: 12,
			err:   fmt.Errorf("nothing found"),
		},
		"two levels in": {
			input: 76,
			key:   76,
			value: "Test data 76",
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			output, err := bt.Search(tc.input)
			if tc.err == nil {
				assert.Nil(t, err, "Got an unexpected error %v", err)
				assert.Equal(t, tc.value, output, "Values did not match, expecting %d, got %d", tc.value, output)
			} else {
				assert.NotNil(t, err, "Expected an error but got nil, expected: %v", tc.err)
				assert.EqualError(t, tc.err, err.Error(), "Error messages did not match, expected %q, got %q", tc.err.Error(), err.Error())
			}
		})
	}
}

func TestInsert(t *testing.T) {
	testcases := map[string]struct {
		keys []int
		vals []interface{}
		size int
		err  error
	}{
		"Single node": {
			keys: []int{21},
			vals: []interface{}{"Test"},
			size: 1,
		},
		"Zero node": {},
		"Eight node": {
			keys: []int{88, 65, 92, 54, 80, 76, 82, 81},
			vals: []interface{}{"Test 88", "Test 65", "Test 92", "Test 54", "Test 80", "Test 76", "Test 82", "Test 81"},
			size: 8,
		},
		"seven node, because of double up": {
			keys: []int{88, 65, 92, 54, 80, 92, 82, 81},
			vals: []interface{}{"Test 88", "Test 65", "Test 92", "Test 54", "Test 80", "Test 92", "Test 82", "Test 81"},
			size: 7,
		},
		"worst case, all nodes to one path": {
			keys: []int{54, 65, 76, 80, 81, 82, 88, 92},
			vals: []interface{}{"Test 54", "Test 65", "Test 76", "Test 80", "Test 81", "Test 82", "Test 88", "Test 92"},
			size: 8,
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			bt := binary.Tree{}
			// fill the tree with nodes
			for idx := range tc.keys {
				err := bt.Insert(tc.keys[idx], tc.vals[idx])
				assert.Nil(t, err, "Did not expect an error inserting K: %d V: %v", tc.keys[idx], tc.vals[idx])
			}
			toVisit := []*binary.Node{bt.Root}
			size := 0
			keys := map[int]bool{}
			vals := map[interface{}]bool{}

			// Walk the tree to gather metrics on it
			for len(toVisit) != 0 {
				n := toVisit[0]
				// If the node is nil we break
				if n == nil {
					break
				}
				keys[n.Key] = true
				vals[n.Value] = true
				// track size in case the 'set' nature of the tree has been
				// violated
				size++

				toVisit = toVisit[1:]

				if n.Left != nil {
					toVisit = append(toVisit, n.Left)
				}
				if n.Right != nil {
					toVisit = append(toVisit, n.Right)
				}

			}

			// Check the number of nodes, keys, and vals is what is expected
			assert.Equalf(t, tc.size, size, "got wrong number of nodes for insertion, expected %d, got %d ", tc.size, size)
			assert.Equalf(t, tc.size, len(keys), "got wrong number of keys for insertion, expected %d, got %d ", tc.size, len(keys))
			assert.Equalf(t, tc.size, len(vals), "got wrong number of keys for insertion, expected %d, got %d ", tc.size, len(vals))

			// check that keys contains an instance of each possible key
			for idx := range tc.keys {
				assert.True(t, keys[tc.keys[idx]], "expected %d to be in keys", tc.keys[idx])
			}

			sort.Ints(tc.keys) // Make the search a little easier, really not needed
			// check that each key should have been created
			for k := range keys {
				assert.Containsf(t, tc.keys, k, "key %d not found in expected keys", k)
			}

			// check that vals contains an instance of each possible val
			for idx := range tc.vals {
				assert.True(t, vals[tc.vals[idx]], "expected %v to be in vals", tc.vals[idx])
			}

			// check that each val should have been created
			for k := range vals {
				assert.Containsf(t, tc.vals, k, "val %v not found in expected vals", k)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	testcases := map[string]struct {
		beforeKeys []int
		afterKeys  []int
		vals       []interface{}
		size       int
		toDelete   int
		err        error
	}{
		"Single node, no deletion": {
			beforeKeys: []int{21},
			afterKeys:  []int{21},
			vals:       []interface{}{"Test"},
			size:       1,
			toDelete:   55,
		},
		"Delete Single node": {
			beforeKeys: []int{21},
			afterKeys:  []int{},
			vals:       []interface{}{"Test"},
			size:       0,
			toDelete:   21,
		},
		"Zero node": {
			toDelete: 55,
			err:      fmt.Errorf("no root"),
		},
		"Eight node, no delete": {
			beforeKeys: []int{88, 65, 92, 54, 80, 76, 82, 81},
			afterKeys:  []int{88, 65, 92, 54, 80, 76, 82, 81},
			vals:       []interface{}{"Test 88", "Test 65", "Test 92", "Test 54", "Test 80", "Test 76", "Test 82", "Test 81"},
			size:       8,
			toDelete:   55,
		},
		"Delete Root": {
			beforeKeys: []int{88, 65, 92, 54, 80, 76, 82, 81},
			afterKeys:  []int{65, 92, 54, 80, 76, 82, 81},
			vals:       []interface{}{"Test 88", "Test 65", "Test 92", "Test 54", "Test 80", "Test 76", "Test 82", "Test 81"},
			size:       7,
			toDelete:   88,
		},
		"Delete Eight node": {
			beforeKeys: []int{88, 65, 92, 54, 80, 76, 82, 81},
			afterKeys:  []int{88, 65, 92, 80, 76, 82, 81},
			vals:       []interface{}{"Test 88", "Test 65", "Test 92", "Test 54", "Test 80", "Test 76", "Test 82", "Test 81"},
			size:       7,
			toDelete:   54,
		},
		"Delete Eight node replacement is a direct child": {
			beforeKeys: []int{88, 65, 92, 54, 80, 76, 82, 81},
			afterKeys:  []int{88, 65, 92, 54, 76, 82, 81},
			vals:       []interface{}{"Test 88", "Test 65", "Test 92", "Test 54", "Test 80", "Test 76", "Test 82", "Test 81"},
			size:       7,
			toDelete:   80,
		},
		"Delete Eight node replacement is not a direct child": {
			beforeKeys: []int{88, 65, 92, 54, 80, 76, 82, 81},
			afterKeys:  []int{88, 92, 54, 80, 76, 82, 81},
			vals:       []interface{}{"Test 88", "Test 65", "Test 92", "Test 54", "Test 80", "Test 76", "Test 82", "Test 81"},
			size:       7,
			toDelete:   65,
		},
		"Delete Eight tree is a single branch": {
			beforeKeys: []int{54, 65, 76, 80, 81, 82, 88, 92},
			afterKeys:  []int{88, 92, 54, 80, 65, 82, 81},
			vals:       []interface{}{"Test 88", "Test 65", "Test 92", "Test 54", "Test 80", "Test 76", "Test 82", "Test 81"},
			size:       7,
			toDelete:   76,
		},
		"Delete Nine right tree is a multi branch": {
			beforeKeys: []int{65, 76, 75, 80, 54, 81, 82, 58, 92, 77, 53},
			afterKeys:  []int{92, 54, 75, 58, 80, 65, 82, 81, 77, 53},
			vals:       []interface{}{"Test 88", "Test 88", "Test 88", "Test 65", "Test 92", "Test 54", "Test 80", "Test 76", "Test 82", "Test 81", "Test 77"},
			size:       10,
			toDelete:   76,
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			// fmt.Println("#################################################")
			bt := binary.Tree{}
			// fmt.Println("BEGIN: ", bt)
			// Build the tree
			for idx := range tc.beforeKeys {
				err := bt.Insert(tc.beforeKeys[idx], tc.vals[idx])
				assert.Nil(t, err, "Did not expect an error inserting K: %d V: %v", tc.beforeKeys[idx], tc.vals[idx])
			}

			// fmt.Println("POST: ", bt)
			// Call Delete
			err := bt.Delete(tc.toDelete)
			// fmt.Println("AFTER: ", bt)
			if tc.err != nil {
				assert.NotNil(t, err, "was expecting an error")

			} else {
				assert.Nil(t, err, "was not expecting an error")

			}

			toVisit := []*binary.Node{bt.Root}
			// fmt.Printf("toVisit: %#v\n", toVisit)
			size := 0
			keys := map[int]bool{}
			vals := map[interface{}]bool{}
			visited := map[*binary.Node]bool{}
			lap := false

			// Walk the tree to gather metrics on it
			for len(toVisit) != 0 {
				fmt.Println(toVisit)
				n := toVisit[0]
				// fmt.Printf("N: %#v\n", n)
				// If the node is nil we break
				if n == nil {
					break
				}
				if visited[n] {
					if lap {

						break
					} else {
						lap = true
					}
				}
				visited[n] = true
				keys[n.Key] = true
				vals[n.Value] = true
				// track size in case the 'set' nature of the tree has been
				// violated
				size++
				if size > 50 {
					assert.FailNow(t, "size greater than 50`")
				}

				toVisit = toVisit[1:]

				if n.Left != nil {
					toVisit = append(toVisit, n.Left)
				}
				if n.Right != nil {
					toVisit = append(toVisit, n.Right)
				}

			}
			// fmt.Println(size)

			// Check the number of nodes, keys, and vals is what is expected
			assert.Equalf(t, tc.size, size, "got wrong number of nodes after deletion, expected %d, got %d ", tc.size, size)
			assert.Equalf(t, tc.size, len(keys), "got wrong number of keys after deletion, expected %d, got %d ", tc.size, len(keys))
			// assert.Equalf(t, tc.size, len(vals), "got wrong number of vals after deletion, expected %d, got %d ", tc.size, len(vals))

			// check that keys contains an instance of each possible key
			for _, val := range tc.afterKeys {
				assert.True(t, keys[val], "expected %d to be in keys", val)
			}

			sort.Ints(tc.afterKeys) // Make the search a little easier, really not needed
			// check that each key should have been created
			for k := range keys {
				assert.Containsf(t, tc.afterKeys, k, "key %d not found in expected keys after deletion", k)
			}
			fmt.Println(keys)

			// check that vals contains an instance of each possible val
			// for idx := range tc.vals {
			// assert.True(t, vals[tc.vals[idx]], "expected %v to be in vals after deletion", tc.vals[idx])
			// }

			// check that each val should have been created
			for k := range vals {
				assert.Containsf(t, tc.vals, k, "val %v not found in expected vals after deletion", k)
			}
		})
	}
}
