package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("constant:")
	constant()
	fmt.Println("linier:")
	linier(4)
	fmt.Println("O(n2):")
	On2(4)
}

func constant() {
	start := time.Now()
	time.Sleep(time.Second)
	fmt.Println(time.Until(start).Abs())
}

func linier(len int) {
	start := time.Now()
	for i := 0; i < len; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%d. %v\n", i+1, time.Until(start).Abs())
	}
}

func On2(len int) {
	start := time.Now()
	for i := 0; i < len; i++ {
		// time.Sleep(time.Second)
		for j := 0; j < len; j++ {
			time.Sleep(time.Second)
			fmt.Printf("[%d][%d]...%v\n", i+1, j+1, time.Until(start).Abs())
		}
	}
}
