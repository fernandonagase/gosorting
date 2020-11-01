package main

import "fmt"

import "github.com/nagasefernando/gosorting/sorting"

func main() {
	slice := []int{3, 6, 4, 1, 9, 7, 0}

	fmt.Println("Original array: ", slice)

	fmt.Println("Bubble Sort: ", sorting.BubbleSort(slice))
	fmt.Println("Insertion Sort: ", sorting.InsertionSort(slice))
	fmt.Println("Selection Sort: ", sorting.SelectionSort(slice))
	fmt.Println("Quick Sort: ", sorting.QuickSort(slice))
	fmt.Println("Merge Sort: ", sorting.MergeSort(slice))
	fmt.Println("Heap Sort: ", sorting.HeapSort(slice))
}
