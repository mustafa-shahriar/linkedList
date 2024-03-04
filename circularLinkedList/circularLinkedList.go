package circularlinkedlist

import (
	"errors"
	"fmt"
)

type CircularLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

type Node struct {
	Value int
	Next  *Node
}

func (list *CircularLinkedList) GetAtIndex(index int) (error, *Node) {
	if index < 0 || index >= list.Size {
		return errors.New("Invalid index"), nil
	}

	node := list.Head

	for i := 0; i < index; i++ {
		node = node.Next
	}

	return nil, node

}

func (list *CircularLinkedList) PrintAllNode() {
	if list.Size == 0 {
		fmt.Println("List is empty")
		return
	}

	node := list.Head

	for {
		fmt.Println(node)

		if node.Next == list.Head {
			break
		}

		node = node.Next
	}
}

func (list *CircularLinkedList) InsertFirst(value int) *Node {

	var node Node = Node{
		Value: value,
	}

	if list.Size == 0 {
		list.Head = &node
		list.Tail = &node
	} else {
		node.Next = list.Head
		list.Head = &node
		list.Tail.Next = list.Head
	}

	list.Size++

	return &node
}

func (list *CircularLinkedList) InsertLast(value int) *Node {
	node := Node{
		Value: value,
		Next:  list.Head,
	}

	list.Tail.Next = &node
	list.Tail = &node
	list.Size++

	return &node
}

func (list *CircularLinkedList) InsertAtIndex(value int, index int) (error, *Node) {
	if index == 0 {
		return nil, list.InsertFirst(value)
	}
	if index == list.Size {
		return nil, list.InsertLast(value)
	}

	err, node := list.GetAtIndex(index - 1)

	if err != nil {
		return err, node
	}
	newNode := Node{
		Value: value,
		Next:  node.Next,
	}
	node.Next = &newNode
	list.Size++

	return nil, &newNode
}

func (list *CircularLinkedList) DeleteFirst() *Node {
	node := list.Head
	list.Head = list.Head.Next
	list.Tail.Next = list.Head
	list.Size--

	return node
}

func (list *CircularLinkedList) DeleteLast() *Node {
	tail := list.Tail

	_, node := list.GetAtIndex(list.Size - 2)

	node.Next = list.Head
	list.Tail = node
	list.Size--

	return tail
}

func (list *CircularLinkedList) DeleteAtIndex(index int) *Node {
	if index == 0 {
		return list.DeleteFirst()
	}
	if index == list.Size-1 {
		return list.DeleteLast()
	}

	err, node := list.GetAtIndex(index - 1)

	if err != nil {
		return nil
	}

	node.Next = node.Next.Next
	list.Size--

	return node
}
