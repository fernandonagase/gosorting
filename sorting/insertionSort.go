package sorting

// InsertionSort sorts an array of integers using the insertion sort algorithm
func InsertionSort(array []int) []int {
	sortedArray := make([]int, len(array))
	copy(sortedArray, array)

	for i := 1; i < len(sortedArray); i++ {
		current := sortedArray[i]
		j := i - 1
		for j >= 0 && current < sortedArray[j] {
			sortedArray[j + 1] = sortedArray[j]
			j--
		}
		sortedArray[j + 1] = current
	}

	return sortedArray
}