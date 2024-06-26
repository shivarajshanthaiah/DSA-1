package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) insert(data int) {
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
		fmt.Print("->", temp.data)
		temp = temp.next
	}
	fmt.Println()

}

func (list *LinkedList) deleteLastNode() {
	temp := list.head
	var prev *Node

	for temp.next != nil {
		prev = temp
		temp = temp.next
	}

	prev.next = nil
}

func main() {
	list := LinkedList{}
	list.insert(20)
	list.insert(30)
	list.insert(40)

	list.display()

	list.deleteLastNode()
	list.display()
}
