package main

import "fmt"

func countingSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	// Find the maximum and minimum values in the input array
	max := arr[0]
	min := arr[0]

	for _, num := range arr {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	// Create the counting array and initialize it with zeros
	countingArray := make([]int, max-min+1)
	for _, num := range arr {
		countingArray[num-min]++
	}

	// Create the output array
	outputArray := make([]int, len(arr))
	index := 0

	// Fill the output array by reading counts from the counting array
	for i, count := range countingArray {
		for j := 0; j < count; j++ {
			outputArray[index] = i + min
			index++
		}
	}

	return outputArray
}

func main() {
	arr := []int{5, 2, 2, 3, 9, 34, 1, 1}
	fmt.Println("Unsorted Array:", arr)

	sortedArray := countingSort(arr)
	fmt.Println("Sorted Array:", sortedArray)
}
