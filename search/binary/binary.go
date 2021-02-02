// Package binary -
package binary

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
	mid := (len(input) - 1) / 2
	top := len(input) - 1
	bot := 0
	finished := false
	for finished == false {
		switch {
		case input[mid] == find:
			return mid, nil
		case input[mid] <= find:
			if top == mid {
				finished = true
			}
			mid = mid + ((top - mid) / 2) + 1
			bot = mid
		case input[mid] >= find:
			if bot == mid {
				finished = true
			}
			mid = mid - ((mid - bot) / 2) - 1
			top = mid
		default:
			finished = true
		}
	}

	return 0, fmt.Errorf("could not find %d", find)
}
