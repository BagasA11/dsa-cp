package bublesort

import (
	_ "array/swap"
	"fmt"
	"time"
)

func BubleSort(arr []int) {
	fmt.Println("before: ", arr)
	start := time.Now()
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
			// time.Sleep(time.Second)
		}
	}
	fmt.Println("after: ", arr)
	fmt.Println("time: ", time.Until(start).Abs())
}
