package main

import "fmt"

func main() {

	nums := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	key := 9
	index := search(nums[:], key)

	if index == -1 {
		fmt.Println("key not found")
	} else {
		fmt.Println("key found at index := ", index)
	}

}

func search(n []int, key int) int {

	for i := 0; i < len(n); i++ {
		if n[i] == key {
			return i
		}
	}
	return -1
}
