package swap

func Swap(a *int, b *int) {
	// fmt.Println("a dan b ", *a, " ", *b)
	temp := *a
	a = b
	b = &temp
	// fmt.Println("a dan b ", *a, " ", *b)
}
