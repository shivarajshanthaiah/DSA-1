package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) InsertAtEnd(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (list *LinkedList) PrintList() {
	current := list.head
	for current != nil {
		fmt.Printf("%d, ->", current.data)
		current= current.next
	}
	fmt.Println("nil")
}

func main(){
	list:=LinkedList{}

	list.InsertAtEnd(20)
	list.InsertAtEnd(30)
	list.InsertAtEnd(10)
	list.InsertAtEnd(60)
	list.PrintList()
}
