// Package chunk -
package chunk

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
	// chunk the input
	chunkSize := 100
	length := len(input) - 1
	i := length
	once := false
	for i >= 0 {
		if input[i] <= find {
			// search this chunk
			for j := i; j <= i+chunkSize; j++ {
				if j > length {
					break
				}
				if input[j] == find {
					return j, nil
				}
			}
		}
		i -= chunkSize
		if i < 0 && !once {
			once = true
			i = 0
		}
	}

	return 0, fmt.Errorf("could not find %d", find)
}
