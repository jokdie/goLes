package main

import "fmt"

// Функция должна циклически сдвинуть массив вправо на k позиций.
func RotateRight(nums []int, k int) []int {
	lNums := len(nums)
	if lNums == 0 {
		return nums
	}

	k %= lNums
	if k == 0 {
		return nums
	}

	tempArr := make([]int, lNums)

	for i, v := range nums {
		tempArr[(i+k)%lNums] = v
	}

	return tempArr
}

func main() {
	fmt.Println("[], k =", 0, "result:", RotateRight([]int{}, 0), "OK: []")                                                            // []int{}
	fmt.Println("[1 2 3], k =", 0, "result:", RotateRight([]int{1, 2, 3}, 0), "OK: [1 2 3]")                                           // []int{1,2,3}
	fmt.Println("[2 3 4 5 1], k =", 1, "result:", RotateRight([]int{2, 3, 4, 5, 1}, 1), "OK: [1 2 3 4 5]")                             // []int{1,2,3,4,5}
	fmt.Println("[1 2 3 4 5], k =", 1, "result:", RotateRight([]int{1, 2, 3, 4, 5}, 1), "OK: [5 1 2 3 4]")                             // []int{5,1,2,3,4}
	fmt.Println("[1 2 3 4 5], k =", 2, "result:", RotateRight([]int{1, 2, 3, 4, 5}, 2), "OK: [4 5 1 2 3]")                             // []int{4,5,1,2,3}
	fmt.Println("[1 2 3], k =", 3, "result:", RotateRight([]int{1, 2, 3}, 3), "OK: [1 2 3]")                                           // []int{1,2,3}
	fmt.Println("[1 2 3], k =", 4, "result:", RotateRight([]int{1, 2, 3}, 4), "OK: [3 1 2]")                                           // []int{3,1,2}
	fmt.Println("[6 7 7 8 1 2 3 4 5], k =", 2, "result:", RotateRight([]int{6, 7, 7, 8, 1, 2, 3, 4, 5}, 2), "OK: [4 5 6 7 7 8 1 2 3]") // []int{4,5,6,7,7,8,1,2,3}
}
