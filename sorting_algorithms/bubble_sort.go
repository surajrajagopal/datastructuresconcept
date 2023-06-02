package sortingalgorithms

// []int{3, 5, 0, 8, 1}
// []int{3, 0, 5, 8, 1}
// []int{3, 0, 5, 1, 8}
// []int{0, 3, 5, 1, 8}
// []int{0, 3, 1, 5, 8}
// []int{0, 1, 3, 5, 8}

func Bubblesort(items []int) []int {
	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
	return items
}

func BubbleSort1(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
