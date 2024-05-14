package main

import "fmt"

func main() {

	str := "Hello go world hi!!!"

	for i := len(str) - 1; i >= 0; i-- {
		fmt.Printf("%c", str[i])
	}
}
