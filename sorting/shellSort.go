package sorting

// ShellSort sorts an array of integers using the shell sort algorithm
func ShellSort(array []int) []int {
	sortedArray := make([]int, len(array))
	copy(sortedArray, array)

	for gap := gap(len(sortedArray)); gap > 0; gap /= 2 {
		for i := gap; i < len(sortedArray); i += gap {
			current := sortedArray[i]
			j := i - gap
			for j >= 0 && current < sortedArray[j] {
				sortedArray[j + gap] = sortedArray[j]
				j -= gap
			}
			sortedArray[j + gap] = current
		}
	}

	return sortedArray
}

func gap(size int) int {
	gap := 1
	for 3 * gap + 1 <= size {
		gap = 3 * gap + 1
	}
	return gap
}