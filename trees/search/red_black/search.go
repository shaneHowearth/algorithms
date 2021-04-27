package redblack

import "fmt"

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
		cmp = cmp.Child[Right]
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
		cmp = cmp.Child[Left]
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

// GetSuccessor -
// In order successor is
// the minimum element in its right subtree (which is the in-order successor)
func (rt *Tree) GetSuccessor(node *Node) (*Node, error) {
	rt.l.RLock()
	defer rt.l.RUnlock()
	if node == nil {
		return nil, fmt.Errorf("no node supplied")
	}

	for {
		if node.Child[Left] == nil {
			return node, nil
		}
		node = node.Child[Left]
	}
}

// GetPredecessor -
// the maximum element in its left subtree (which is the in-order predecessor)
func (rt *Tree) GetPredecessor(node *Node) (*Node, error) {
	rt.l.RLock()
	defer rt.l.RUnlock()
	if node == nil {
		return nil, fmt.Errorf("no node supplied")
	}

	for {
		if node.Child[Right] == nil {
			return node, nil
		}
		node = node.Child[Right]
	}
}
