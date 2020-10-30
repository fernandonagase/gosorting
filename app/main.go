package main

import "fmt"

import "github.com/nagasefernando/gosorting/sorting"

func main() {
	slice := []int{3, 6, 4, 1, 9, 7, 0}

	fmt.Println("Bubble Sort: ", sorting.BubbleSort(slice))
	fmt.Println("Insertion Sort: ", sorting.InsertionSort(slice))
	fmt.Println("Selection Sort: ", sorting.SelectionSort(slice))
	fmt.Println("Quick Sort: ", sorting.QuickSort(slice))
}
