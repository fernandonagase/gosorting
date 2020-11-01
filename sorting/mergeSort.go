package sorting

// MergeSort sorts an array of integers using the merge sort algorithm
func MergeSort(array []int) []int {
	sortedArray := make([]int, len(array))
	copy(sortedArray, array)

	mergeSortRecursive(sortedArray)

	return sortedArray
}

func mergeSortRecursive(array []int) {
	size := len(array)

	// se tamanho é 1: retorna
	if size == 1 {
		return
	}

	leftSlice := array[:size / 2]
	rightSlice := array[size / 2:]
	// ordenaEsquerda
	mergeSortRecursive(leftSlice)
	// ordenaDireita
	mergeSortRecursive(rightSlice)

	// combina esquerda e direita
	copy(array, merge(leftSlice, rightSlice))
}

func merge(left, right []int) (merged []int) {
	i, j, k := 0, 0, 0
	array := make([]int, len(left) + len(right))

	for j < len(left) && k < len(right) {
		var min int
		if left[j] < right[k] {
			min = left[j]
			j++
		} else {
			min = right[k]
			k++
		}
		array[i] = min
		i++
	}

	if j < len(left) {
		copy(array[i:], left[j:])
	}

	if k < len(right) {
		copy(array[i:], array[k:])
	}

	return array
}