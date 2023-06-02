package searchingalgorithms

import (
	"fmt"
	"sort"
)

func BinarySearch(arr []int, x int) int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println("sorted", arr)
	low := 0
	high := len(arr) - 1

	if low > high {
		return -1
	}

	for low <= high {
		mid_index := (low + high) / 2
		mid_value := arr[mid_index]
		if mid_value == x {
			return mid_index
		} else if mid_value < x {
			low = mid_index + 1
		} else {
			high = mid_index - 1
		}
	}
	return -1
}
