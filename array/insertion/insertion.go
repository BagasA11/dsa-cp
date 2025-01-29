package insertion

import (
	"fmt"
	"time"
)

func InsertionSort(arr []int) {
	fmt.Println("before: ", arr)
	start := time.Now()
	n := len(arr)
	for i := 1; i < n; i++ {
		// outer loop to iterate over unsorted array
		curr := arr[i] //current array
		j := i - 1     // look over previous array
		for j >= 0 && arr[j] > curr {
			arr[j+1] = arr[j] //switch array with lower value
			time.Sleep(time.Second)
			j--
		}
		arr[j+1] = curr
	}

	fmt.Println("after: ", arr)
	fmt.Println("time: ", time.Until(start).Abs())
}
