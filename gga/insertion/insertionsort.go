// Package insertion -
package insertion

// Sort -
func Sort(array []int) {
	for idx := range array {
		if idx == 0 {
			continue
		}
		key := array[idx]
		j := idx - 1
		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = key
	}
}
