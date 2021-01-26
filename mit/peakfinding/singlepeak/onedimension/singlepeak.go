// Package singlepeak -
package singlepeak

import "fmt"

// Peak -
type Peak struct {
	input []int
}

// NewPeak - create a new peak instance with data
func NewPeak(input []int) (*Peak, error) {
	if len(input) == 0 {
		return nil, fmt.Errorf("no input to check")
	}
	return &Peak{input: input}, nil
}

// Linear - O(n)
func (p *Peak) Linear() int {
	for i := range p.input {
		if i == len(p.input)-1 {
			return i
		}
		if p.input[i+1] < p.input[i] {
			return i
		}
	}
	// Will never reach here, should be deadcode
	return len(p.input) - 1
}

// DandC - Divide and Conquer O(logn)
func (p *Peak) DandC() int {
	// start in the middle
	i := len(p.input) / 2
	return p.recursiveFind(i)

}

func (p *Peak) recursiveFind(idx int) int {
	// if left bigger, move left (note: if bother are bigger, only the LHS will
	// be evaluated)
	// if right bigger, move right
	// if neither return (we have 'a' peak)
	if idx == 0 || idx == len(p.input)-1 {
		return idx
	}
	// lhs
	if p.input[idx-1] >= p.input[idx] {
		return p.recursiveFind(idx / 2)
	}
	// rhs
	// move to the mid point between the current index and the end of the input
	if p.input[idx+1] >= p.input[idx] {
		return p.recursiveFind(idx + ((len(p.input) - idx) / 2))
	}
	// peak
	return idx
}
