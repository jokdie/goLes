package main

import "fmt"

func FirstUniqueChar(s string) rune {
	if len(s) == 0 {
		return 0
	}
	counts := make(map[rune]int)

	for _, c := range s {
		counts[c] += 1
	}

	for _, c := range s {
		if counts[c] == 1 {
			return c
		}
	}

	return 0
}

func main() {
	fmt.Printf("aabbcddee: %c\n", FirstUniqueChar("aabbcddee")) // 'c'
	fmt.Printf("leetcode: %c\n", FirstUniqueChar("leetcode"))   // 'l'
	fmt.Printf("aabb: %c\n", FirstUniqueChar("aabb"))           // 0
	fmt.Printf(": %c\n", FirstUniqueChar(""))                   // 0
	fmt.Printf("приветп: %c\n", FirstUniqueChar("приветп"))     // 'р'
}
