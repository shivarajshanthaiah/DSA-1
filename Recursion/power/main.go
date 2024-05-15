package main

import "fmt"

func main() {
	num := 3
	x := 5
	fmt.Println(power(x, num))

}

func power(x int, n int) int {
	if n == 0 {
		return 1
	}
	return x * (power(x, n-1))
}
