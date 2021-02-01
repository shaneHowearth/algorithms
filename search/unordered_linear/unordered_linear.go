// Package unorderedlinear -
package unorderedlinear

import "fmt"

// Search -
func Search(find int, input []int) (int, error) {
	if len(input) == 0 {
		return 0, fmt.Errorf("nothing to search")
	}
	for i := range input {
		if input[i] == find {
			return i, nil
		}
	}
	return 0, fmt.Errorf("could not find %d", find)
}
