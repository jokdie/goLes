package main

import "fmt"

func RemoveDuplicatesMy(nums []int) []int {
	result := []int{}
	if len(nums) == 0 {
		return result
	}

	unique := make(map[int]int)

	for _, v := range nums {
		unique[v]++
	}

	for _, v := range nums {
		if _, ok := unique[v]; ok {
			result = append(result, v)
			delete(unique, v)
		}
	}

	return result
}

func RemoveDuplicates(nums []int) []int {
	result := []int{}
	seen := map[int]struct{}{}

	for _, v := range nums {
		if _, ok := seen[v]; !ok {
			result = append(result, v)
			seen[v] = struct{}{}
		}
	}

	return result
}

func main() {
	fmt.Println("[1, 2, 2, 3, 1, 4]:", RemoveDuplicates([]int{1, 2, 2, 3, 1, 4}))                                           // []int{1, 2, 3, 4}
	fmt.Println("[1, 2, 2, 3, 1, 4, 0]:", RemoveDuplicates([]int{1, 2, 2, 3, 1, 4, 0}))                                     // []int{1, 2, 3, 4, 0}
	fmt.Println("[5, 5, 5]:", RemoveDuplicates([]int{5, 5, 5}))                                                             // []int{5}
	fmt.Println("[]:", RemoveDuplicates([]int{}))                                                                           // []int{}
	fmt.Println("[1, 4, 2, 2, 9, 3, 1, 4, 2, 7, 7, 9, 8]:", RemoveDuplicates([]int{1, 4, 2, 2, 9, 3, 1, 4, 2, 7, 7, 9, 8})) // []int{1, 4, 2, 9, 3, 7, 8}
}
