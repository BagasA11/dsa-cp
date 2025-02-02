package main

// Operasi dasar pada stack adalah:
// - Push: Menambahkan data baru deri paling belakang.
// - Pop: Menghapus data baru dari urutan paling belakang.
// - Peek: Melihat data pada urutan terbelakang
// - IsEmpety: Mengecek apakah tumpukan kosong.
// - Size: Mencari jumlah data pada stack

import (
	"fmt"
	a "stack/stack_array"
	s "stack/stack_linklist"
)

// func newArr(arr [5]int, index uint) [5]int {
// 	var newArr []int
// 	if index == 0 {
// 		newArr = arr[1:5]
// 		return [5]int(newArr)
// 	}

// }

func main() {
	stackArr := [5]int{}
	a.Push(&stackArr, 4)
	fmt.Println(stackArr)
	fmt.Println("kosong? ", a.IsEmpety(stackArr))
	a.Push(&stackArr, 2)
	fmt.Println(stackArr)
	pop := a.Pop(&stackArr)
	fmt.Println("pop: ", pop)
	fmt.Println("peek: ", a.Peek(stackArr))
	fmt.Println(&stackArr)

	// ==========
	fmt.Println()
	stack := s.NewStack(1)
	fmt.Println("stack last value ", stack.GetLayer())
	fmt.Println("stack size ", stack.GetSize())
	stack.Push(3)
	stack.Push(-4)
	fmt.Println("stack last value ", stack.GetLayer())
	fmt.Println("stack size ", stack.GetSize())
	fmt.Println("deleting last layer: ", stack.Pop())
	fmt.Println("stack last value ", stack.GetLayer())
	fmt.Println("stack size ", stack.GetSize())
}
