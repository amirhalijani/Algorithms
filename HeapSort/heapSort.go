package main

import "fmt"

// Heapify function converts an array into a max heap.
func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	// Find the largest element among the root, left child, and right child.
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// If the largest element is not the root, swap it with the root.
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		// Recursively heapify the affected sub-tree.
		heapify(arr, n, largest)
	}
}

// HeapSort function performs the heap sort algorithm on an array.
func heapSort(arr []int) {
	n := len(arr)

	// Build a max heap.
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Extract elements from the heap one by one.
	for i := n - 1; i > 0; i-- {
		// Move the current root to the end of the array.
		arr[0], arr[i] = arr[i], arr[0]
		// Call heapify on the reduced heap.
		heapify(arr, i, 0)
	}
}

func main() {
	arr := []int{25, 43, 12, 7, 2, 34, 67, 78}
	fmt.Println("Unsorted Array:", arr)

	heapSort(arr)
	fmt.Println("Sorted Array:", arr)
}
