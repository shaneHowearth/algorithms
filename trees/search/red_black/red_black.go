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
