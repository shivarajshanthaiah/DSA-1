package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) insertAtEnd(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
		return
	}

	temp := list.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode

}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func (list *LinkedList) printPrimes() {
	temp := list.head
	found := false
	for temp != nil {
		if isPrime(temp.data) {
			fmt.Print("", temp.data)
			found = true
		}
		temp = temp.next
	}
	if !found {
		fmt.Println("Not found")
	}
	fmt.Println()
}

func (list *LinkedList) display() {
	temp := list.head
	for temp != nil {
		fmt.Print("-> ", temp.data)
		temp = temp.next
	}
	fmt.Println("nil")

}

func main() {
	list := LinkedList{}
	list.insertAtEnd(20)
	list.insertAtEnd(40)
	list.insertAtEnd(31)

	list.display()
	list.printPrimes()
}
