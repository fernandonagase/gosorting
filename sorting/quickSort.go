package sorting

func QuickSort(array []int) []int {
	sortedArray := make([]int, len(array))
	copy(sortedArray, array)

	quickSortRecursive(sortedArray)

	return sortedArray
}

func quickSortRecursive(array []int) {
	p, r := 0, len(array) - 1

	if len(array[p:r + 1]) <= 1 {
		return
	}
	pivot := r

	i := p - 1
	for j := p; j < r; j++ {
		if array[j] < array[pivot] {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}
	array[i + 1], array[r] = array[r], array[i + 1]

	quickSortRecursive(array[:i + 1])
	quickSortRecursive(array[i + 1:])
}