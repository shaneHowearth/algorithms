// Package orderedlinear -
package orderedlinear

import (
	"fmt"
	"sort"
)

// Search -
func Search(find int, input []int) (int, error) {
	if len(input) == 0 {
		return 0, fmt.Errorf("nothing to search")
	}
	sort.Ints(input)
	for i := range input {
		if input[i] == find {
			return i, nil
		}
	}
	return 0, fmt.Errorf("could not find %d", find)
}
