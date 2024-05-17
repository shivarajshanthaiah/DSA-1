package main

import "fmt"

func main() {
	num := 10
	printFib(num)

	// fmt.Println("Fibunacci is :", fib(num))
}

func fib(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return n
	}
	fib1 := fib(n - 1)
	fib2 := fib(n - 2)

	fibunacci := fib1 + fib2

	return fibunacci
}

func printFib(n int) {
	fmt.Println("Series ")
	for i := 0; i <= n; i++ {
		fmt.Printf("%d, ", fib(i))
	}
}
