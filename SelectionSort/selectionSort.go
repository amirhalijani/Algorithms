package main

import "fmt"

// SelectionSort sorts an array of integers in ascending order using the Selection Sort algorithm.
func SelectionSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		// Find the minimum element in the unsorted portion of the array
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// Swap the minimum element with the first element in the unsorted portion
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {
	arr := []int{11, 7, 43, 2, 34, 78, 56}
	fmt.Println("Unsorted Array:", arr)

	SelectionSort(arr)

	fmt.Println("Sorted Array:", arr)
}
