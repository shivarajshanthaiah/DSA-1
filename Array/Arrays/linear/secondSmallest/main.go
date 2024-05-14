package main

import (
	"fmt"
)

func main() {
	arr := [5]int{35, 38, 83, 49, 35}

	fmt.Println("Second smallest element is :", secSmall(arr[:]))
}

func secSmall(nums []int) int {
	// sort.Ints(nums)
	// return nums[1]

	smallest := nums[0]
	secSmallest := nums[1]
	if smallest > secSmallest {
		secSmallest, smallest = smallest, secSmallest
	}

	for i := 2; i < len(nums); i++ {
		if nums[i] < smallest {
			secSmallest, smallest = smallest, nums[i]
		} else if nums[i] < secSmallest && nums[i] != smallest {
			secSmallest = nums[i]
		}
	}
	return secSmallest
}
