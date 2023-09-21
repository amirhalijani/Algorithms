package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)

	// Repeat the process n times, where n is the number of elements in the array.
	for i := 0; i < n; i++ {
		// Each pass through the array compares adjacent pairs of elements and swaps them if they are in the wrong order.
		// After each pass, the largest unsorted element "bubbles up" to its correct position at the end of the array.
		// So, we can reduce the number of comparisons in each pass by ignoring the already sorted portion at the end of the array.
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{21, 4, 16, 90, 35, 78, 43}

	fmt.Println("Unsorted Array:", arr)

	bubbleSort(arr)

	fmt.Println("Sorted Array:", arr)
}
