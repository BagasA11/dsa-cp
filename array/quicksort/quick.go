package quicksort

import (
	"fmt"
	"time"
)

func Quicksort(arr []int) {
	start := time.Now()
	fmt.Println("quick sort:")
	fmt.Println("before: ", arr)

	fmt.Println("after: ", arr)
	fmt.Println("time: ", time.Until(start).Abs())

}
