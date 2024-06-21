package singleLinkedList

import (
	"errors"
)

type LinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	Size int
}

type Node[T any] struct {
	Value *T
	Next  *Node[T]
}

func (linkedList *LinkedList[T]) GetAtIndex(index int) *Node[T] {
	if index < 0 || index >= linkedList.Size {
		return nil
	}

	node := linkedList.Head

	for i := 0; i < index; i++ {
		node = node.Next
	}

	return node
}

func (linkedList *LinkedList[T]) InsertFirst(value *T) *Node[T] {

	var node Node[T] = Node[T]{
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

func (linkedList *LinkedList[T]) InsertLast(value *T) *Node[T] {

	var node Node[T] = Node[T]{
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

func (linkedList *LinkedList[T]) InsertAtIndex(value *T, index int) (error, *Node[T]) {
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

	newNode := Node[T]{
		Value: value,
		Next:  node.Next,
	}
	node.Next = &newNode
	linkedList.Size++

	return nil, &newNode
}

func (linkedList *LinkedList[T]) DeleteFirst() *Node[T] {
	if linkedList.Size == 0 {
		return nil
	}

	if linkedList.Size == 1 {
		temp := linkedList.Head
		linkedList.Head = nil
		linkedList.Tail = nil
		linkedList.Size--
		return temp
	}

	temp := linkedList.Head
	linkedList.Head = linkedList.Head.Next
	linkedList.Size--

	return temp
}

func (linkedList *LinkedList[T]) DeleteLast() *Node[T] {
	if linkedList.Size <= 1 {
		return linkedList.DeleteFirst()
	}

	node := linkedList.GetAtIndex(linkedList.Size - 2)
	temp := node.Next
	node.Next = nil
	linkedList.Tail = node
	linkedList.Size--

	return temp
}

func (linkedList *LinkedList[T]) DeleteAtIndex(index int) *Node[T] {
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

// func (linkedList *LinkedList[T]) PrintAllNode() {
// 	if linkedList.Size == 0 {
// 		fmt.Println("List is empty")
// 	}
//
// 	node := linkedList.Head
// 	fmt.Println("---------")
//
// 	for node != nil {
// 		fmt.Println(node.Value)
// 		fmt.Println("---------")
// 		node = node.Next
// 	}
// }
