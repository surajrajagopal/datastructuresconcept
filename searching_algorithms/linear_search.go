package searchingalgorithms

func LinearSearch1(arr []int, x int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == x {
			return true
		}
	}
	return false
}

func LinearSearch2(arr []int, x int) int {
	for i, v := range arr {
		if v == x {
			return i
		}
	}
	return -1
}
