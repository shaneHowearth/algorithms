// Package bubble -
package bubble

// Sort - Sorts array
// Note: This sorts the array in place, it relies on the way that Go passes
// slices around (that is, a slice is a struct with len, cap, and a pointer to
// the beginning of the underlying array, when Go passes a slice around it
// passes a copy of that struct, meaning that the copy you have has a pointer to
// the original array)
func Sort(array []int) {
	length := len(array)
	for i := range array {
		for j := 0; j < length-i-1; j++ {
			if array[j] > array[j+1] {
				//swap
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}
