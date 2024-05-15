package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(isSorted(arr[:], 0))
}

func isSorted(arr []int, i int) bool {
	if i == len(arr)-1 {
		return true
	}
	if arr[i] > arr[i+1] {
		return false
	}
	return isSorted(arr, i+1)
}
