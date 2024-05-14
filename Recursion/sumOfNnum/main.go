package main

import (
	"fmt"
)

func main() {
	num := 10
	fmt.Print(sum(num))
	fmt.Println()
}

func sum(n int) int {
	if n == 1 {
		return 1
	}
	return n + sum(n-1)
}
