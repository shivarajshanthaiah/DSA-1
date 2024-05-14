package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Reversed", reverse(arr[:]))
}

func reverse(num []int) []int {
	first := 0
	last := len(num) - 1
	for first < last {
		num[first], num[last] = num[last], num[first]
		first++
		last--
	}
	return num
}
