package sorting

// HeapSort sorts an array of integers using the heap sort algorithm
func HeapSort(array []int) []int {
	sortedArray := make([]int, len(array))
	copy(sortedArray, array)

	buildHeap(sortedArray)
	for m := heapToIndex(len(sortedArray)); m >= heapToIndex(2); m-- {
		swap(sortedArray, heapToIndex(1), m)
		reorganize(sortedArray[:m])
	}

	return sortedArray
}

func buildHeap(array []int) {
	for i := 1; i < heapToIndex(len(array)); i++ {
		n := i + 1
		for n > 1 && array[heapToIndex(n)] >= array[heapToIndex(n / 2)] {
			swap(array, heapToIndex(n), heapToIndex(n / 2))
			n /= 2
		}
	}
}

func reorganize(array []int) {
	n := 2
	for n <= len(array) {
		if n < len(array) && array[heapToIndex(n + 1)] > array[heapToIndex(n)] {
			n++
		}
		if array[heapToIndex(n / 2)] >= array[heapToIndex(n)] {
			break
		}
		swap(array, heapToIndex(n), heapToIndex(n / 2))
		n *= 2
	}
}

func heapToIndex(heapIndex int) (index int) {
	return heapIndex - 1
}

func swap(array []int, x, y int) {
	array[x], array[y] = array[y], array[x]
}