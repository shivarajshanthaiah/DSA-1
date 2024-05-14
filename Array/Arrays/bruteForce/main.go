package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, -2, 6, -1, 3}
	maxSubarrySum(arr[:])

}

func maxSubarrySum(nums []int) {
	currSum := 0
	maxSum := nums[0]

	for i := 0; i < len(nums); i++ {
		currSum = 0
		for j := i; j < len(nums); j++ {
			currSum += nums[j]
			if currSum > maxSum {
				maxSum = currSum
			}
		}
	}
	fmt.Println("Max sum is :", maxSum)

}
