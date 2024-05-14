package main

import "fmt"

func main() {
	arr := [5]int{55, 33, 78, 92, 30}

	fmt.Println("Largest element is:", largest(arr[:]))
}

func largest(n []int) int {
	temp := n[0]
	for i := 1; i < len(n); i++ {
		if n[i] > temp {
			temp = n[i]
		}
	}
	return temp
}
