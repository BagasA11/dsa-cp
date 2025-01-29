package selectionsort

import (
	"fmt"
	"time"
)

func SelectSort(arr []int) {
	start := time.Now()
	n := len(arr)
	fmt.Println("before: ", arr)
	for i := 0; i < n; i++ {
		low := arr[i]
		post := i
		for j := i; j < n; j++ {
			if arr[j] < low {
				low = arr[j]
				post = j
			}

			time.Sleep(time.Second)
		}
		temp := arr[i]
		arr[i] = low
		arr[post] = temp
	}
	fmt.Println("after: ", arr)
	fmt.Println("time: ", time.Until(start).Abs())
}
