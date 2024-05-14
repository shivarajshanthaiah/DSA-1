package main

import "fmt"

func main() {
	str := "hello"

	result := pal(str)

	if result {
		fmt.Println("It is a palindrome")
	} else {
		fmt.Println("Not a palindrome")
	}

}

func pal(name string) bool {
	for i := 0; i < len(name)/2; i++ {
		n := len(name)
		if name[i] != name[n-i-1] {
			return false
		}
	}
	return true
}
