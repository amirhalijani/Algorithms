package main

import "fmt"

// MergeSort function divides and sorts the input slice using merge sort.
func MergeSort(arr []int) []int {
	// Base case: If the slice has zero or one element, it's already sorted.
	if len(arr) <= 1 {
		return arr
	}

	// Split the slice into two halves.
	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	// Recursively sort both halves.
	left = MergeSort(left)
	right = MergeSort(right)

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	l, r := 0, 0

	// Compare elements from the left and right slices and add the smaller
	// element to the result slice.
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	// Add any remaining elements from the left and right slices.
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}

func main() {
	arr := []int{11, 32, 56, 7, 21, 3, 72, 5}
	fmt.Println("Unsorted Array:", arr)

	sortedArr := MergeSort(arr)
	fmt.Println("Sorted Array:", sortedArr)
}
