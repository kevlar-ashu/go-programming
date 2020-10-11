package main

import (
	"github.com/kevlar-ashu/go-programming.git/searching"
	"github.com/kevlar-ashu/go-programming.git/sorting"

	"fmt"
)

func main() {
	fmt.Println("Linear Search......")

	myList := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	fmt.Println(searching.LinearSearch(myList, 58))

	fmt.Println("Binary Search....")
	arr := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(searching.BinarySearch(63, arr))

	fmt.Println("Bubble sort.....")

	slice := sorting.GenerateSlice(20)

	fmt.Println("\n--- Unsorted --- \n\n", slice)

	sorting.Bubblesort(slice)

	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")

	fmt.Println("QuickSort tuttorial.....going on...")
	data := []int{8, 7, 6, 1, 0, 9, 2}
	fmt.Println("Before sorting.....")
	fmt.Println(data)
	length := len(data) - 1
	sorting.QuickSort(data, 0, length)
	fmt.Println("After sorting....")
	fmt.Println(data)
}
