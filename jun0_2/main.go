package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(s string) bool {
	switch len(s) {
	case 0, 1:
		return true
	}

	sSlice := []rune(strings.ToLower(s))
	lenSlice := len(sSlice)

	for i := 0; i < lenSlice/2; i++ {
		if sSlice[i] != sSlice[lenSlice-1-i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Printf("abba: %t\n", IsPalindrome("abba"))     // true
	fmt.Printf("level: %t\n", IsPalindrome("level"))   // true
	fmt.Printf("lev2el: %t\n", IsPalindrome("lev2el")) // false
	fmt.Printf("hello: %t\n", IsPalindrome("hello"))   // false
	fmt.Printf("a: %t\n", IsPalindrome("a"))           // true
	fmt.Printf(" : %t\n", IsPalindrome(""))            // true
}
