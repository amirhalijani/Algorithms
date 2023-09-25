package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 1 {
		return arr
	}

	pivot := arr[0] // Choose the first element as the pivot

	var less, greater []int

	for _, item := range arr[1:] {
		if item <= pivot {
			less = append(less, item)
		} else {
			greater = append(greater, item)
		}
	}

	// Recursively sort the sub-arrays and combine them
	less = quickSort(less)
	greater = quickSort(greater)

	// Concatenate less + pivot + greater to get the sorted array
	return append(append(less, pivot), greater...)
}

func main() {
	arr := []int{72, 1, 156, 62, 13, 27, 8, 36}
	fmt.Println("Unsorted Array:", arr)

	arr = quickSort(arr)
	fmt.Println("Sorted Array:", arr)
}
