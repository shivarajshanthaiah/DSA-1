package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	traverse(arr, 0)
}

func traverse(nums []int, i int) {
	if i == len(nums) {
		return
	}

	fmt.Println(nums[i])
	traverse(nums, i+1)
}
