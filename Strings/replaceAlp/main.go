package main

import (
	"fmt"
)

func main() {

	str := "Hello World, I am learning golang"
	bytes := []byte(str)

	count := 0
	for i := 0; i < len(bytes); i++ {
		if bytes[i] != ' ' {
			count++
		}
		if count%3 == 0 {
			bytes[i] = 'H'
		}
	}

	res := string(bytes)
	fmt.Println(res)

}
