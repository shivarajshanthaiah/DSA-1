package main

import "fmt"

func main() {
	num := 10

	res := factorial(num)
	fmt.Println("Answer is :", res)
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}
