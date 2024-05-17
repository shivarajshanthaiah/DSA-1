package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	reverse(arr, len(arr)-1)
}

func reverse(num []int, i int) {
	if i < 0 {
		return
	}
	fmt.Print(" ", num[i])
	reverse(num, i-1)
}
