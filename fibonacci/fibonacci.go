package main

// Deret Fibonacci adalah urutan angka di mana setiap angka berikutnya adalah jumlah dari dua angka sebelumnya. Dimulai dari 0 dan 1, deret ini berjalan sebagai berikut:
// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...
// Dalam notasi matematis, deret Fibonacci dapat ditulis sebagai:
// F(0) = 0
// F(1) = 1
// F(n) = F(n-1) + F(n-2) untuk n ≥ 2

// Pada sesi ini saya akan membuat program untuk menampilkan deret fibonacci dengan 3 cara, yaitu
// 1. Perulangan
// 2. Rekursif
// 3. Mencari nilai ke-n dari deret fibonacci menggunakan rekursif

import (
	"fmt"
)

func main() {
	fmt.Println("loop:")
	loopFibo(5)
	fmt.Println("recursive")
	fmt.Println(recursFibo(5))
}

// fungsi menampilkan deret dengan perulangan
// 2 variabel 0 dan 1, dan variabel temp untuk menampung hasil
// looping sebanyak n
func loopFibo(n uint) {
	a, b := 0, 1
	var fibo int

	for i := uint(0); i < n; i++ {
		if i >= 2 {
			fibo = a + b
			a = b
			b = fibo
			fmt.Printf("%d ", fibo)
			continue
		}
		fibo = a
		a = b
		// b = fibo
		fmt.Printf("%d ", fibo)
	}
}

// recursif fibonaci
// untuk n >= 2, f(n) = f(n-1)+f(n-2)
func recursFibo(n int) int {
	if n <= 1 {
		return n
	}
	return recursFibo(n-1) + recursFibo(n-2)
}
