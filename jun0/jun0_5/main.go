package main

import (
	"fmt"
)

func MostFrequentChar(s string) rune {
	if len(s) == 0 {
		return 0
	}

	counts := map[rune]int{}

	for _, v := range s {
		counts[v]++
	}

	var bestChar rune
	bestCount := 0

	for _, v := range s {
		if counts[v] > bestCount {
			bestChar = v
			bestCount = counts[v]
		}
	}

	return bestChar
}

func main() {
	fmt.Printf("hello: %c\n", MostFrequentChar("hello"))           // 'l'
	fmt.Printf("aaabbcc: %c\n", MostFrequentChar("aaabbcc"))       // 'a'
	fmt.Printf("aaabbccccc: %c\n", MostFrequentChar("aaabbccccc")) // 'a'
	fmt.Printf("abc: %c\n", MostFrequentChar("abc"))               // 'a'
	fmt.Printf(": %c\n", MostFrequentChar(""))                     // 0
	fmt.Printf("приветтт: %c\n", MostFrequentChar("приветтт"))     // т
}
