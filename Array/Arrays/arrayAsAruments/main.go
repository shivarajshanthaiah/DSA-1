package main

import "fmt"

func main() {
	arr := [3]int{1, 5, 9}
	fmt.Println(arr)
	update(&arr)
	fmt.Println(arr)
}

func update(up *[3]int) {
	for i := 0; i < len(*up); i++ {
		up[i] += 1
	}
}
