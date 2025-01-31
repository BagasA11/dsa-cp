package main

import (
	"array/bublesort"
	"array/insertion"
	"array/quicksort"
	sel "array/selection_sort"
	"fmt"
	"time"
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
	fmt.Println("quick sort")
	s := time.Now()
	var newArr = []int{-2, 3, 0, 4, 1}
	fmt.Println("before ", newArr)
	quicksort.Quicksort(newArr, 0, len(newArr)-1)
	fmt.Println("after ", newArr)
	fmt.Println("time: ", time.Until(s).Abs())
}
