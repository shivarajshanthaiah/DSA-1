package main

import "fmt"

func main() {
	str := "Hello World I am learning Go"
	bytes := []byte(str)

	for i := 0; i < len(bytes); i++ {
		if bytes[i] >= 'a' && bytes[i] <= 'z' {
			bytes[i] = bytes[i] - 32
		}
	}
	res := string(bytes)
	fmt.Println(res)
}
