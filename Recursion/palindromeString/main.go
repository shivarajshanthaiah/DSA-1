package main

import "fmt"

func main() {
	str := "radar"
	fmt.Println(isPalindrome(str, 0, len(str)-1))

}

func isPalindrome(s string, left int, right int) bool {
	if left > right {
		return true
	}
	if s[left] != s[right] {
		return false
	}
	return isPalindrome(s, left+1, right-1)
}
