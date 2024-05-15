package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 1122322114

	s := strconv.Itoa(num)

	fmt.Println(isPalNum(s, 0, len(s)-1))
}

func isPalNum(n string, left, right int) bool {
	if left > right {
		return true
	}

	if n[left] != n[right] {
		return false
	}
	return isPalNum(n, left+1, right-1)
}
