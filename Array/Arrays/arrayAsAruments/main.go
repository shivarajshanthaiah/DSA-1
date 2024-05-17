package main

import "fmt"

func main() {
	arr := [3]int{1, 5, 9}
	fmt.Println(arr)
	update(&arr)
	fmt.Println(arr)
}

func update(update *[3]int) {
	for i := 0; i < len(*update); i++ {
		update[i] += 1
	}
}
