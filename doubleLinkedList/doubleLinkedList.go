package doublelinkedlist

import (
	"errors"
	"fmt"
)

type DoubleLinkedList struct {
	Size int
	Head *Node
	Tail *Node
}

type Node struct {
	Value    int
	Previous *Node
	Next     *Node
}

func (list *DoubleLinkedList) GetAtIndex(index int) (error, *Node) {
	if list.Size == 0 {
		return nil, nil
	}

	if index < 0 || index >= list.Size {
		return errors.New("Invalid index"), nil
	}

	node := list.Head

	for i := 0; i < index; i++ {
		node = node.Next
	}

	return nil, node

}

func (list *DoubleLinkedList) InsertFirst(value int) *Node {
	node := Node{
		Value: value,
	}

	if list.Size == 0 {
		list.Head = &node
		list.Tail = &node
		list.Size++
		return &node
	}

	list.Head.Previous = &node
	node.Next = list.Head
	list.Head = &node
	list.Size++
	return &node
}

func (list *DoubleLinkedList) InsertLast(value int) *Node {
	if list.Size == 0 {
		return list.InsertFirst(value)
	}

	node := Node{
		Value:    value,
		Previous: list.Tail,
	}
	list.Tail.Next = &node
	list.Tail = &node
	list.Size++

	return &node
}

func (list *DoubleLinkedList) InsertAtIndex(value int, index int) (error, *Node) {
	if index == 0 {
		return nil, list.InsertFirst(value)
	}
	if index == list.Size {
		return nil, list.InsertLast(value)
	}

	err, node := list.GetAtIndex(index - 1)

	if err != nil {
		return err, nil
	}

	newNode := Node{
		Value:    value,
		Next:     node.Next,
		Previous: node,
	}

	node.Next = &newNode
	newNode.Next.Previous = &newNode
	list.Size++

	return nil, &newNode
}

func (list *DoubleLinkedList) DeleteFirst() *Node {
	if list.Size == 0 {
		return nil
	}

	if list.Size == 1 {
		node := list.Head
		list.Head = nil
		list.Tail = nil
		list.Size--
		return node
	}

	node := list.Head
	list.Head = list.Head.Next
	list.Head.Previous = nil
	list.Size--

	return node
}

func (list *DoubleLinkedList) DeleteLast() *Node {
	if list.Size <= 1 {
		return list.DeleteFirst()
	}

	node := list.Tail
	list.Tail = list.Tail.Previous
	list.Tail.Next = nil
	list.Size--

	return node
}

func (list *DoubleLinkedList) DeleteAtIndex(index int) (error, *Node) {
	if index == 0 {
		return nil, list.DeleteFirst()
	}
	if index == list.Size-1 {
		return nil, list.DeleteLast()
	}

	err, node := list.GetAtIndex(index)

	if err != nil {
		return err, node
	}

	node.Previous.Next = node.Next
	node.Next.Previous = node.Previous
	list.Size--

	return nil, node
}

func (list *DoubleLinkedList) PrintAllNode() {
	if list.Size == 0 {
		fmt.Println("List is Empty")
	}
	node := list.Head
	fmt.Println("------------------- size:", list.Size)
	fmt.Println("value | previous | next")
	for node != nil {
		fmt.Println(node)
		node = node.Next
	}
	fmt.Println("-------------------")
}
