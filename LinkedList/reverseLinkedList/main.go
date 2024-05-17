package main

import "fmt"

type Node struct {
	data byte
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) insertAtEnd(data byte) {
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

func (list *LinkedList) display() {
	temp := list.head
	for temp != nil {
		fmt.Printf("%c->,", temp.data)
		temp = temp.next
	}
	fmt.Println()
}

func (list *LinkedList) reverse() {
	var prev *Node
	current := list.head
	var next *Node
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	list.head = prev
}

func main() {

	list := LinkedList{}

	list.insertAtEnd('2')
	list.insertAtEnd('d')
	list.insertAtEnd('f')
	list.insertAtEnd('f')
	list.insertAtEnd('b')
	list.display()

	list.reverse()
	list.display()

}
