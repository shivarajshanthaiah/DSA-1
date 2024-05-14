package main

import "fmt"

func main() {
	arr := [9]int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	key := 8

	fmt.Println("Found at Index :", binarySearch(arr[:], key))

}

func binarySearch(num []int, key int) int {
	start := 0
	end := len(num) - 1
	for start <= end {
		mid := (start + end) / 2

		if num[mid] == key {
			return mid
		}

		if num[mid] < key {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
