package main

import (
	"fmt"
	"strings"
)

func CountVowels(s string) int {
	count := 0

	for _, char := range strings.ToLower(s) {
		switch char {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(CountVowels("hello"))
	fmt.Println(CountVowels("HELLO"))
	fmt.Println(CountVowels("Go Lang"))
	fmt.Println(CountVowels(""))
	fmt.Println(CountVowels("bcdfg"))
}
