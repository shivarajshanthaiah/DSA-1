package main

import "fmt"

func main() {
	arr := []int{12, 15, 17, 19, 21, 33, 44, 55, 66, 77, 88}

	tar := 12
	ans := binarySearch(arr, 0, len(arr)-1, tar)
	if ans == -1 {
		fmt.Println("Not found")
	} else {
		fmt.Printf("Forund at index: %d", ans)
	}
	fmt.Println("")
}

func binarySearch(arr []int, left, right, target int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2

	if arr[mid] == target {
		return mid
	} else if arr[mid] > target {
		return binarySearch(arr, left, mid-1, target)
	} else {
		return binarySearch(arr, mid+1, right, target)
	}
}
