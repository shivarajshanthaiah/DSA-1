package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := sumArray(arr, len(arr))
	fmt.Println(sum)
}

func sumArray(arr []int, n int) int {
	if n == 0 {
		return 0
	}

	return arr[n-1] + sumArray(arr, n-1)
}
