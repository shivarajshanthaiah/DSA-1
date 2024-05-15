package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 6, 9, 34, 67, 83, 31}
	key := 81
	fmt.Println(firstOcc(arr[:], key, 0))

}

func firstOcc(nums []int, key int, i int) (int int) {
	if i == len(nums) {
		return -1
	}
	if nums[i] == key {
		return i
	}
	return firstOcc(nums, key, i+1)
}
