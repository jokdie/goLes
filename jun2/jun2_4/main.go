package main

import (
	"fmt"
	"jun2_4/mathx"
	"jun2_4/stringx"
)

func main() {
	fmt.Println(stringx.Reverse("abc"))
	fmt.Println(stringx.IsEmpty("abc"))
	fmt.Printf("Max: %d\n", mathx.Max(1, 10))
	fmt.Printf("Min: %d\n", mathx.Min(1, 10))
}
