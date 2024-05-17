package main

import "fmt"

func main() {
	value := 10
	decreasing(value)
}

func decreasing(nums int) {
	if nums == 1 {
		fmt.Println(1)
		return
	}
	fmt.Print(nums," ")
	decreasing(nums-1)

}
