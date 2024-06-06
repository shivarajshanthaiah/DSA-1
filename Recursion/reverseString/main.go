package main

import "fmt"

func main() {
	str := "hello world"
	ans := reversed(str)
	fmt.Println(ans)
}

func reversed(s string) string {
	if len(s) == 0 {
		return s
	}
	return reversed(s[1:]) + string(s[0])
}
