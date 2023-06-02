package sortingalgorithms

func InsertionSort(elements []int) []int {
	var n = len(elements)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if elements[j-1] > elements[j] {
				elements[j-1], elements[j] = elements[j], elements[j-1]
			}
			j = j - 1
		}
	}
	return elements
}
