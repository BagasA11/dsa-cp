package main

import (
	"array/bublesort"
	"array/insertion"
	sel "array/selection_sort"
	"fmt"
)

func main() {
	// var array = []int{2, 1, 4, 3, 5}
	// duplicate := array // reset array after perform sort
	fmt.Println("bublesort:")
	bublesort.BubleSort([]int{2, 1, 4, 3, 5})
	// array = duplicate
	fmt.Println("Selection sort:")
	sel.SelectSort([]int{2, 1, 4, 3, 5})
	// array = duplicate
	fmt.Println("insertion sort")
	insertion.InsertionSort([]int{2, 1, 4, 3, 5})
}
