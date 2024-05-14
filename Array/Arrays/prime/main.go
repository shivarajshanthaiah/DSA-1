package main

import "fmt"

func main() {

	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, n := range num {
		if prime(n) {
			fmt.Print(" ", n)
		}
	}
	fmt.Println()

}

func prime(nums int) bool {
	if nums < 2 {
		return true
	}

	for i := 2; i < nums/2; i++ {
		if nums%i == 0 {
			return false
		}
	}
	return true
}
