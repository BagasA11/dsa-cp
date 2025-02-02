package stack_array

import "fmt"

func size(arr [5]int) uint {
	size := 0
	for _, v := range arr {
		if v != 0 {
			size++
		}
	}
	return uint(size)
}

func Push(arr *[5]int, data int) {
	// count size of arr -> if there not enough space then do nothing
	// x := *arr
	size := size(*arr)
	if data == 0 {
		data++ //if data 0, data++
	}
	if size == 5 {
		fmt.Println("stack sudah penuh")
		return
	}
	// array start from 0 to n-1
	if size == 0 {
		fmt.Println("menambahkan elemen pertama")
		arr[0] = data
		return
	}
	fmt.Println("menambahkan elemen ke-", size)
	arr[size] = data
}

func Pop(arr *[5]int) int {
	size := size(*arr)
	if size == 0 {
		fmt.Println("stack has empty")
		return 0
	}
	val := arr[size-1]
	arr[size-1] = 0
	return val
}

func Peek(arr [5]int) int {
	size := size(arr)
	return arr[size-1]
}

func IsEmpety(arr [5]int) bool {
	return size(arr) == 0
}
