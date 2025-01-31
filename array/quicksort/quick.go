package quicksort

import (
	"fmt"
)

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func partition(arr []int, low, high int) int {
	fmt.Println(arr)
	pivot := arr[high]
	index := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			index++
			swap(&arr[index], &arr[j])
		}
		// time.Sleep(time.Second)
	}
	swap(&arr[index+1], &arr[high])
	return index + 1
}

func Quicksort(arr []int, low, high int) {
	// start := time.Now()
	// fmt.Println("quick sort:")
	// fmt.Println("before: ", arr)

	if low < high {
		pivotIndex := partition(arr, low, high)
		Quicksort(arr, low, high-1)
		Quicksort(arr, pivotIndex+1, high)
	}
	fmt.Println(arr)
	// fmt.Println("after: ", arr)
	// fmt.Println("time: ", time.Until(start).Abs())
}
