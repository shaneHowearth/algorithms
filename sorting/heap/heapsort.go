// Package heap -
package heap

// Sort -
func Sort(array []int) {
	heapSort(array, len(array))
}

// To heapify a subtree rooted with node i which is
// an index in arr[]. n is size of heap
func heapify(arr []int, n int, i int) {
	largest := i // Initialize largest as root
	l := 2*i + 1 // left = 2*i + 1
	r := 2*i + 2 // right = 2*i + 2

	// If left child is larger than root
	if l < n && arr[l] > arr[largest] {
		largest = l
	}

	// If right child is larger than largest so far
	if r < n && arr[r] > arr[largest] {
		largest = r
	}

	// If largest is not root
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		// Recursively heapify the affected sub-tree
		heapify(arr, n, largest)
	}
}

// main function to do heap sort
func heapSort(arr []int, n int) {
	// Build heap (rearrange array)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// One by one extract an element from heap
	for i := n - 1; i > 0; i-- {
		// Move current root to end
		arr[0], arr[i] = arr[i], arr[0]

		// call max heapify on the reduced heap
		heapify(arr, i, 0)
	}
}
