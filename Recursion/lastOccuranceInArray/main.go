package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 6, 9, 34, 67, 3, 31}
	key := 31
	fmt.Println(lastOcc(arr, key, 0))
}

func lastOcc(nums []int, key int, i int) int {
	if i == len(nums) {
		return -1
	}
	found := (lastOcc(nums, key, i+1))
	if found == -1 && nums[i] == key {
		return i
	}
	return found
}