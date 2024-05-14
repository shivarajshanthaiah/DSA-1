package main

import "fmt"

func main() {
	str := "Hello Go World hi!!"
	fmt.Println("Vowels Are :")
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == 'a' || str[i] == 'e' || str[i] == 'i' || str[i] == 'o' || str[i] == 'u' {
			count++
			fmt.Printf("%c\t", str[i])
		}
	}
	fmt.Println("count = ", count)
}