package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Insert(data int) {
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

func (list *LinkedList) deleteFirstNode() {
	if list.head == nil {
		fmt.Println("List is empty")
	}
	list.head = list.head.next
}

func (list *LinkedList) display() {
	temp := list.head

	for temp != nil {
		fmt.Print("-> ", temp.data)
		temp = temp.next
	}
	fmt.Println()
	
}

func main() {
	list := LinkedList{}

	list.Insert(30)
	list.Insert(20)
	list.Insert(50)

	list.display()

	list.deleteFirstNode()
	list.display()
}
