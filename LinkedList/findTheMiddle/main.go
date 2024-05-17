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

func (list *LinkedList) findMiddle() {
	if list.head == nil {
		return
	}

	slow := list.head
	fast := list.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	fmt.Println("",slow.data)
}

func (list *LinkedList) display() {
	temp := list.head

	for temp != nil {
		fmt.Print("->", temp.data)
		temp = temp.next
	}
	fmt.Println()
}

func main() {

	list := LinkedList{}
	list.insertAtEnd(20)
	list.insertAtEnd(33)
	list.insertAtEnd(45)
	list.insertAtEnd(20)
	list.insertAtEnd(3)
	list.insertAtEnd(45)
	list.insertAtEnd(20)
	list.insertAtEnd(33)
	list.insertAtEnd(45)

	list.display()

	list.findMiddle()

}
