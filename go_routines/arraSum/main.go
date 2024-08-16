package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numChan := make(chan int)
	resChan := make(chan int)

	go func() {
		defer close(numChan)
		for i := 0; i < len(arr); i++ {
			numChan <- arr[i]
		}
	}()

	go func() {
		defer close(resChan)
		sum := 0
		for num := range numChan {
			sum += num
		}
		resChan <- sum
	}()
	res := <-resChan
	fmt.Println(res)
}
