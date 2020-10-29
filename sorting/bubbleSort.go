package sorting

// BubbleSort sorts an array of integers using the bubble sort algorithm
func BubbleSort(array []int) []int {
	hasSwapped := true
	iteration := 1
	sortedArray := array

	for iteration < len(sortedArray) && hasSwapped {
		hasSwapped = false
		for i := 0; i < len(sortedArray)-1; i++ {
			if sortedArray[i] > sortedArray[i+1] {
				hasSwapped = true
				aux := sortedArray[i]
				sortedArray[i] = sortedArray[i+1]
				sortedArray[i+1] = aux
			}
		}
	}

	return sortedArray
}
