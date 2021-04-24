// Package redblack -
/*
Wikipedia: Red–black trees offer worst-case guarantees for insertion time, deletion time, and search time. Not only does this make them valuable in time-sensitive applications such as real-time applications, but it makes them valuable building blocks in other data structures which provide worst-case guarantees; for example, many data structures used in computational geometry can be based on red–black trees, and the Completely Fair Scheduler used in current Linux kernels and epoll system call implementation[22] uses red–black trees.
*/
package redblack

import (
	"fmt"
	"sync"
)

// Tree - Red-Black Tree
type Tree struct {
	l    sync.RWMutex
	Root *Node // Nil if tree empty
}

// Colour -
type Colour int

// Colours -
const (
	Black Colour = iota
	Red
)

// Node - Red Black Node
type Node struct {
	Key    int
	Value  interface{}
	Parent *Node
	Child  [2]*Node
	Colour Colour
}

// Direction -
type Direction int

// Constants -
const (
	Left Direction = iota
	Right
)

// RotateDirRoot -
func (rt *Tree) RotateDirRoot(
	Parent *Node, // Parent of node that's out of place
	dir Direction,
) error {
	rt.l.Lock()
	defer rt.l.Unlock()

	if Parent.Parent != nil {
	}
	Grandparent := Parent.Parent
	Sibling := Parent.Child[1-dir]

	if Sibling != nil {
		C := Sibling.Child[dir]
		Parent.Child[1-dir] = C
		if C != nil {
			C.Parent = Parent
		}
		Sibling.Child[dir] = Parent
		Parent.Parent = Sibling
		Sibling.Parent = Grandparent
		if Grandparent != nil {
			d := Left
			if Parent == Grandparent.Child[Right] {
				d = Right
			}
			Grandparent.Child[d] = Sibling
		} else {
			rt.Root = Sibling
		}
		return nil
	}
	return fmt.Errorf("parent's other child is nil")
}

// Insert -
func (rt *Tree) Insert(
	key int,
	value interface{},
) error {
	// Ensure new node is coloured red
	node := &Node{Key: key, Colour: Red, Value: value}

	// Find correct point to place new node
	if rt.Root == nil {
		rt.Root = node
		node.Colour = Black
		return nil
	}

	// Walk down to node's new home
	cmp := rt.Root
	for {
		if node.Key < cmp.Key {
			if cmp.Child[Left] != nil {
				cmp = cmp.Child[Left]
				continue
			} else {
				cmp.Child[Left] = node
				node.Parent = cmp
				break
			}
		}
		if node.Key > cmp.Key {
			if cmp.Child[Right] != nil {
				cmp = cmp.Child[Right]
				continue
			} else {
				cmp.Child[Right] = node
				node.Parent = cmp
				break
			}
		}
		if node.Key == cmp.Key {
			return fmt.Errorf("key already exists")
		}
	}

	// Check that its insertion hasn't violated any of the RBTree conditions
	for node != rt.Root && node.Parent.Colour == Red {
		// Insertion scenarios
		// Recurse upwards until we find root, or no violations (no red parent with red
		// children)
		// rt.Root should be Black so there is no need to check
		// Case 1
		if node.Parent.Parent != nil && node.Parent.Parent.Child[Left] != nil && node.Parent.Parent.Child[Left].Colour == Red && node.Parent.Parent.Child[Right] != nil && node.Parent.Parent.Child[Right].Colour == Red {
			// Case 1: Recolour fix
			// If grandparent is black (it /should/ be) and parent's sibling is red
			// Recolor parent and sibling from red to black, and grandparent from black to
			// red.
			node.Parent.Parent.Colour = Red
			node.Parent.Parent.Child[Left].Colour = Black
			node.Parent.Parent.Child[Right].Colour = Black
			node = node.Parent.Parent
			continue
		}

		// Case 2
		changeDirection := Left
		zigzag := false
		if node == node.Parent.Child[Left] && node.Parent == node.Parent.Parent.Child[Right] {
			// Case 2: Zigzag fix
			// Node being inspected (X) is left child of its parent, and parent is right
			// child of its parent (X's grandparent) OR X is the right child of its parent,
			// and the parent is the left child of X's grandparent
			// Make X the parent of its parent, and child of its grandparent such that the
			// new child is the same side as X is on its (now) parent
			// ie, both are left children, or both are right children
			changeDirection = Right
			zigzag = true
		} else if node == node.Parent.Child[Right] && node.Parent == node.Parent.Parent.Child[Left] {
			changeDirection = Left
			zigzag = true
		}
		if zigzag {
			err := rt.RotateDirRoot(node.Parent, changeDirection)
			if err != nil {
				return fmt.Errorf("case 2: rotating %d, direction %d generated %w", node.Parent.Key, changeDirection, err)
			}

		}
		// Case 3
		// X's parent A, and X's grand parent C
		// move X to top of tree (making C and A it's direct children) and recolour
		// recolour
		node.Parent.Colour = 1 - node.Parent.Colour
		if node == node.Parent.Child[Left] {
			changeDirection = Right
		} else {
			changeDirection = Left
		}
		err := rt.RotateDirRoot(node.Parent, changeDirection)
		if err != nil {
			return fmt.Errorf("case 3: rotating %d, direction %d generated %w", node.Parent.Key, changeDirection, err)
		}
	}

	// Set the root to black
	rt.Root.Colour = Black
	return nil
}

// GetMax -
func (rt *Tree) GetMax() (*Node, error) {
	rt.l.RLock()
	defer rt.l.RUnlock()
	if rt.Root == nil {
		return nil, fmt.Errorf("no nodes in the tree")
	}
	cmp := rt.Root
	for {
		if cmp.Child[Right] == nil {
			return cmp, nil
		}

	}
}

// GetMin -
func (rt *Tree) GetMin() (*Node, error) {
	rt.l.RLock()
	defer rt.l.RUnlock()
	if rt.Root == nil {
		return nil, fmt.Errorf("no nodes in the tree")
	}
	cmp := rt.Root
	for {
		if cmp.Child[Left] == nil {
			return cmp, nil
		}
	}
}

// Search -
func (rt *Tree) Search(key int) (*Node, error) {
	rt.l.RLock()
	defer rt.l.RUnlock()
	if rt.Root == nil {
		return nil, fmt.Errorf("no nodes in the tree")
	}
	cmp := rt.Root
	for {
		if key < cmp.Key {
			if cmp.Child[Left] != nil {
				cmp = cmp.Child[Left]
				continue
			} else {
				return nil, fmt.Errorf("%d not found", key)
			}
		}
		if key > cmp.Key {
			if cmp.Child[Right] != nil {
				cmp = cmp.Child[Right]
				continue
			} else {
				return nil, fmt.Errorf("%d not found", key)
			}
		}
		if key == cmp.Key {
			return cmp, nil
		}
	}
}

// Successor -
func (rt *Tree) Successor(key int) (*Node, error) {
	if rt.Root == nil {
		return nil, fmt.Errorf("no nodes in the tree")
	}
	node, err := rt.Search(key)
	if err != nil {
		return nil, err
	}
	// WTH is the successor?
	// if node.
	return node, nil
}

// Predessor -
func (rt *Tree) Predessor(key int) (*Node, error) {
	if rt.Root == nil {
		return nil, fmt.Errorf("no nodes in the tree")
	}
	node, err := rt.Search(key)
	if err != nil {
		return nil, err
	}
	// WTH is the successor?
	// if node.
	return node, nil
}
