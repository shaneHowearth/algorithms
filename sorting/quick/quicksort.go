// Package quick -
package quick

// Sort -
func Sort(array []int) {
	quickSort(array, 0, len(array)-1)

}

/* The main function that implements QuickSort
arr[] --> Array to be sorted,
low --> Starting index,
high --> Ending index */
func quickSort(arr []int, low int, high int) {
	if low < high {
		/* pi is partitioning index, arr[p] is now
		   at right place */
		pi := partition(arr, low, high)

		// Separately sort elements before
		// partition and after
		// partition
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}

}

/* This function takes last element as pivot, places
the pivot element at its correct position in sorted
array, and places all smaller (smaller than pivot)
to left of pivot and all greater elements to right
of pivot */
func partition(arr []int, low int, high int) int {
	pivot := arr[high] // pivot
	i := low - 1       // Index of smaller element and indicates the right position of pivot found so far

	for j := low; j <= high-1; j++ {
		// If current element is smaller than the pivot
		if arr[j] < pivot {
			i++ // increment index of smaller element
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return (i + 1)
}
