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

func (list *LinkedList) insertInMiddle(position int, data int){
	newNode := &Node{data:data}
	if position < 0{
		newNode.next=list.head
		list.head = newNode
		return
	}

	temp := list.head
	for i:=0;i<position-1;i++{
		temp = temp.next
	}

	newNode.next = temp.next
	temp.next=newNode
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
	list.insertAtEnd(30)
	list.insertAtEnd(40)
	list.insertAtEnd(10)


	list.display()


	list.insertInMiddle(1,5)
	list.display()
}
