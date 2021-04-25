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
