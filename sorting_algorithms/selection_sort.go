package sortingalgorithms

func SelectionSort(elements []int) []int {
	var n = len(elements)

	for i := 0; i < n; i++ {
		min_idx := i
		for j := i; j < n; j++ {
			if elements[j] < elements[min_idx] {
				min_idx = j
			}
		}
		elements[i], elements[min_idx] = elements[min_idx], elements[i]
	}
	return elements
}
