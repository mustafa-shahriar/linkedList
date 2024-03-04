package singleLinkedList

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

type Node struct {
	Value int
	Next  *Node
}

func (linkedList *LinkedList) GetAtIndex(index int) *Node {
	node := linkedList.Head

	for i := 0; i < index; i++ {
		node = node.Next
	}

	return node
}

func (linkedList *LinkedList) InsertFirst(value int) *Node {

	var node Node = Node{
		Value: value,
	}

	if linkedList.Size == 0 {
		linkedList.Head = &node
		linkedList.Tail = &node
	} else {
		node.Next = linkedList.Head
		linkedList.Head = &node
	}

	linkedList.Size++

	return &node
}

func (linkedList *LinkedList) InsertLast(value int) *Node {

	var node Node = Node{
		Value: value,
	}

	if linkedList.Size == 0 {
		linkedList.Head = &node
		linkedList.Tail = &node
	} else {
		linkedList.Tail.Next = &node
		linkedList.Tail = &node
	}

	linkedList.Size++

	return &node
}

func (linkedList *LinkedList) InsertAtIndex(value int, index int) (error, *Node) {
	if index < 0 || index > linkedList.Size {
		return errors.New("Invalid index"), nil
	}
	if index == 0 {
		return nil, linkedList.InsertFirst(value)
	}
	if index == linkedList.Size {
		return nil, linkedList.InsertLast(value)
	}

	node := linkedList.GetAtIndex(index - 1)

	for i := 1; i < index; i++ {
		node = node.Next
	}

	newNode := Node{
		Value: value,
		Next:  node.Next,
	}
	node.Next = &newNode
	linkedList.Size++

	return nil, &newNode
}

func (linkedList *LinkedList) DeleteFirst() *Node {
	if linkedList.Size == 0 {
		return nil
	}
	if linkedList.Size == 1 {
		temp := linkedList.Head
		linkedList.Head = nil
		linkedList.Tail = nil
		return temp
	}
	temp := linkedList.Head
	linkedList.Head = linkedList.Head.Next
	linkedList.Size--
	return temp
}

func (linkedList *LinkedList) DeleteLast() *Node {
	if linkedList.Size == 0 {
		return nil
	}
	if linkedList.Size == 1 {
		temp := linkedList.Head
		linkedList.Head = nil
		linkedList.Tail = nil
		return temp
	}

	node := linkedList.GetAtIndex(linkedList.Size - 2)
	temp := node.Next
	node.Next = nil
	linkedList.Tail = node
	linkedList.Size--
	return temp
}

func (linkedList *LinkedList) DeleteAtIndex(index int) *Node {
	if index == 0 {
		return linkedList.DeleteFirst()
	}
	if index == linkedList.Size-1 {
		return linkedList.DeleteLast()
	}
	if index < 0 || index >= linkedList.Size {
		return nil
	}

	node := linkedList.GetAtIndex(index - 1)
	temp := node.Next
	node.Next = node.Next.Next
	linkedList.Size--
	return temp
}

func (linkedList *LinkedList) PrintAllNode() {
	if linkedList.Size == 0 {
		fmt.Println("List is empty")
	}

	node := linkedList.Head
	fmt.Println("---------")

	for node != nil {
		fmt.Println(node.Value)
		fmt.Println("---------")
		node = node.Next
	}
}
