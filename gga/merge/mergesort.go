// Package merge -
package merge

// Sort -
func Sort(array []int) {
	mergeSort(array, 0, len(array)-1)
}

func mergeSort(array []int, l, r int) {
	if l < r {
		mid := (l + r) / 2
		mergeSort(array, l, mid)   // LHS
		mergeSort(array, mid+1, r) // RHS
		merge(array, l, mid, r)
	}
}

func merge(array []int, l, m, r int) {
	// Find sizes of two subarrays to be merged
	n1 := m - l + 1
	n2 := r - m

	/* Create temp arrays */
	L := make([]int, n1)
	R := make([]int, n2)

	/*Copy data to temp arrays*/
	for i := 0; i < n1; i++ {

		L[i] = array[l+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = array[m+1+j]

	}

	/* Merge the temp arrays */

	// Initial indexes of first and second subarrays
	i := 0
	j := 0

	// Initial index of merged subarry array
	k := l
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			array[k] = L[i]
			i++
		} else {
			array[k] = R[j]
			j++
		}
		k++
	}

	/* Copy remaining elements of L[] if any */
	for i < n1 {
		array[k] = L[i]
		i++
		k++
	}

	/* Copy remaining elements of R[] if any */
	for j < n2 {
		array[k] = R[j]
		j++
		k++
	}
}
