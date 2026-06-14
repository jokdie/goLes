package main

import "fmt"

func RemoveAt(nums []int, index int) []int {
	count := len(nums)

	if count == 0 {
		return nums
	}
	if index < 0 || index >= count {
		return nums
	}

	result := make([]int, 0, count)

	for k, v := range nums {
		if k == index {
			continue
		}

		result = append(result, v)
	}

	return result
}

func main() {
	fmt.Println(RemoveAt([]int{1, 2, 3, 4}, 2)) // [1,2,4]
	fmt.Println(RemoveAt([]int{10, 20, 30}, 0)) // [20,30]
	fmt.Println(RemoveAt([]int{5}, 0))          // []
}
