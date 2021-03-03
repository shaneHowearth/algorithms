// Package binary -
package binary

import "fmt"

// Node -
type Node struct {
	Key   int
	Value interface{}
	Left  *Node
	Right *Node
}

// Tree -
type Tree struct {
	Root *Node
}

// Search - search tree for an element with the specified key
func (t *Tree) Search(key int) (interface{}, error) {
	n := t.Root
	for {
		if n == nil {
			return nil, fmt.Errorf("nothing found")
		}
		if key == n.Key {
			return n.Value, nil
		}
		if key < n.Key {
			n = n.Left
		}
		if n != nil && key > n.Key {
			n = n.Right
		}
	}
}

// Insert -
func (t *Tree) Insert(key int, val interface{}) error {
	if t.Root == nil {
		t.Root = &Node{Key: key, Value: val}
		return nil
	}
	// Find appropriate point to insert the new node (if required)
	n := t.Root
	for {
		if key == n.Key {
			return nil // Nothing to insert
		}
		if key < n.Key {
			if n.Left == nil {
				n.Left = &Node{Key: key, Value: val}
				return nil
			}
			n = n.Left
		}
		if key > n.Key {
			if n.Right == nil {
				n.Right = &Node{Key: key, Value: val}
				return nil
			}
			n = n.Right
		}
	}
}

// Delete -
func (t *Tree) Delete(key int) error {
	fmt.Println("#### Delete called with key", key)
	fmt.Println("###", t)
	fmt.Println("###", t.Root)
	// Nothing to delete
	if t.Root == nil {
		return fmt.Errorf("no root")
	}

	if key == t.Root.Key {
		// No children
		if t.Root.Left == nil && t.Root.Right == nil {
			t.Root = nil
			return nil
		}
		// One child
		if t.Root.Left != nil && t.Root.Right == nil {
			t.Root = t.Root.Left
			return nil
		}
		if t.Root.Right != nil && t.Root.Left == nil {
			t.Root = t.Root.Right
			return nil
		}
		fmt.Println("### Two children")
		// Two children :\
		// Find the max in the left hand branch
		parent, replacement := findMax(t.Root.Left)
		if parent == replacement {
			fmt.Println("Parent == Replacement", parent, replacement, t.Root)
			// t.Root.Left has no .Right
			replacement.Right = t.Root.Right
			t.Root = replacement
			return nil
		}
		// Delete and update pointers
		// replacement left becomes tmpParent's right
		parent.Right = replacement.Left
		// deleted nodes right becomes replacement's right
		replacement.Right = t.Root.Right
		// deleted nodes left becomes replacement's left
		replacement.Left = t.Root.Left
		// Move replacement to root
		t.Root = replacement
		return nil
	}

	fmt.Println("#### Setting node ans parent to root")
	node := t.Root
	parent := t.Root
	for {
		fmt.Println("#### infinite loop", node, parent)
		if key == node.Key {
			fmt.Println("Found the key")
			if node.Left == nil && node.Right == nil {
				fmt.Println("No children")
				// no children, update the correct parent pointer
				if parent.Left.Key == key {
					parent.Left = nil
				} else {
					parent.Right = nil
				}
				return nil
			}
			if node.Left == nil && node.Right != nil {
				parent.Right = node.Right
				return nil
			}
			if node.Left != nil && node.Right != nil {
				fmt.Println("Node Left")
				// Find the min child
				tmpParent, replacement := findMin(node.Right)
				fmt.Println("Parent, replacement", tmpParent.Key, replacement.Key)
				// Delete and update pointers
				// replacement.Left = node.Left
				// deleted nodes right becomes replacement's right
				// if node.Right != replacement {
				// replacement.Right = node.Right
				// }
				// deleted nodes left becomes replacement's left
				replacement.Left = node.Left
				// Move replacement to where the deleted node was
				if parent.Right == node {
					parent.Right = tmpParent
				} else {
					parent.Left = replacement
				}
				return nil
			}

			fmt.Println("#### Node right")
			tmpParent, replacement := findMin(node.Left)
			fmt.Println(tmpParent, replacement)
			// Delete and update pointers
			// replacement left becomes tmpParent's right
			tmpParent.Left = replacement.Right
			// deleted nodes right becomes replacement's right
			if node.Left != replacement {
				replacement.Left = node.Left
			}
			// deleted nodes left becomes replacement's left
			replacement.Right = node.Right
			// Move replacement to root
			if parent.Right == node {
				parent.Right = replacement
			} else {
				parent.Left = replacement
			}
			return nil
		}

		// Move down the tree
		if key < node.Key {
			// Nothing to delete
			if node.Left == nil {
				return nil
			}
			parent = node
			node = node.Left
		}
		if key > node.Key {
			// Nothing to delete
			if node.Right == nil {
				return nil
			}
			parent = node
			node = node.Right
		}
	}
}

func findMax(tmp *Node) (*Node, *Node) {
	tmpParent := tmp
	max := tmp
	for tmp != nil {
		tmpParent = max
		max = tmp
		tmp = tmp.Right
	}
	return tmpParent, max
}

func findMin(tmp *Node) (*Node, *Node) {
	tmpParent := tmp
	min := tmp
	for tmp != nil {
		tmpParent = min
		min = tmp
		tmp = tmp.Left
	}
	return tmpParent, min
}
