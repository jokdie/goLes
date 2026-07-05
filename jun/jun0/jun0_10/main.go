package main

import "fmt"

func Swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	x := 10
	y := 20

	Swap(&x, &y)

	fmt.Println(x)
	fmt.Println(y)
}
