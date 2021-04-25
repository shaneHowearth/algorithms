package redblack

import "fmt"

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
