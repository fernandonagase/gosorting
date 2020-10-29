package sorting

// SelectionSort sorts an array of integers using the selection sort algorithm
func SelectionSort(array []int) []int {
	sortedArray := make([]int, len(array))
	copy(sortedArray, array)

	for i := 0; i < len(sortedArray) - 1; i++ {
		currentMin := sortedArray[i]
		minPosition := i
		for j := i + 1; j < len(sortedArray); j++ {
			if sortedArray[j] < currentMin {
				currentMin = sortedArray[j]
				minPosition = j
			}
		}
		if currentMin < sortedArray[i] {
			aux := sortedArray[i]
			sortedArray[i] = currentMin
			sortedArray[minPosition] = aux
		}
	}

	return sortedArray
}