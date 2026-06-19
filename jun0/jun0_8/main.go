package main

import (
	"fmt"
	"slices"
	"strings"
)

func ReverseWords(s string) string {
	temp := strings.Fields(s)
	slices.Reverse(temp)

	return strings.Join(temp, " ")
}

func main() {
	fmt.Println(ReverseWords("hello world"))             // "world hello"
	fmt.Println(ReverseWords("   go   is   awesome   ")) // "awesome is go"
	fmt.Println(ReverseWords("a"))                       // "a"
	fmt.Println(ReverseWords(""))                        // ""
}
