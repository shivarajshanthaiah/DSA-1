package main

import "fmt"

func main() {
	str := "Hello World"
	reverse(str, len(str)-1)
}

func reverse(s string, i int) {
	if i < 0 {
		return
	}
	fmt.Printf("%c",s[i])
	reverse(s, i-1)

}
