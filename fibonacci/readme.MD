# README: Bilangan Fibonacci

## Pengantar
Bilangan Fibonacci adalah urutan angka yang didefinisikan sebagai setiap angka merupakan jumlah dari dua angka sebelumnya. Deret ini dinamakan setelah matematikawan Italia, Leonardo Fibonacci.

## Teori Fibonacci
Deret Fibonacci dimulai dengan dua angka pertama yaitu 0 dan 1, kemudian angka berikutnya adalah hasil penjumlahan dari dua angka sebelumnya. Urutannya adalah sebagai berikut:

0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...

Dalam notasi matematis, deret Fibonacci dapat dituliskan sebagai:
- F(0) = 0
- F(1) = 1
- F(n) = F(n-1) + F(n-2) untuk n ≥ 2

Deret ini memiliki banyak aplikasi dalam berbagai bidang seperti matematika, ilmu komputer, seni, dan biologi.

## Implementasi Algoritma Fibonacci

### Algoritma dengan Looping
Berikut adalah implementasi deret Fibonacci menggunakan looping dalam bahasa Python:

```python
# loop
def fibonacci_loop(n):
    a, b = 0, 1
    for _ in range(n):
        print(a, end=" ")
        a, b = b, a + b

# recursive
def fibonacci_recursive(n):
    if n <= 0:
        return 0
    elif n == 1:
        return 1
    else:
        return fibonacci_recursive(n - 1) + fibonacci_recursive(n - 2)

```

