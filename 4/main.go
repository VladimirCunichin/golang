package main

import (
	"fmt"
)

//Item that is stored inside Linked List
type Item struct {
	Value interface{}
	Next  *Item
	Prev  *Item
}

//LinkedList that can be traversed in two directions
type LinkedList struct {
	Length int
	Head   *Item
	Tail   *Item
}

//NewItem constructor
func NewItem(v interface{}) *Item {
	return &Item{v, nil, nil}
}

//NewLinkedList constructor
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

//First element from Linked List
func (l LinkedList) First() Item {
	return *l.Head
}

func (l LinkedList) Len() int {
	return l.Length
}

//Last element from Linked List
func (l LinkedList) Last() Item {
	return *l.Tail
}

//PushBack Item into Linked List
func (l *LinkedList) PushBack(v interface{}) {
	newItem := NewItem(v)
	if l.Head == nil {
		l.Head = newItem
		l.Tail = newItem
	} else {
		newItem.Prev = l.Tail
		l.Tail.Next = newItem
		l.Tail = newItem
	}
	l.Length++
}

//PushFront Item into Linked List
func (l *LinkedList) PushFront(v interface{}) {
	newItem := NewItem(v)
	if l.Head == nil {
		l.Head = newItem
		l.Tail = newItem
	} else {
		newItem.Next = l.Head
		l.Head.Prev = newItem
		l.Head = newItem
	}
	l.Length++
}

func (l *LinkedList) Remove(i Item) {
	if l.Length == 0 {
		return
	}
	switch {
	case i.Prev == nil:
		l.Head = i.Next
		i.Next.Prev = nil
	case i.Next == nil:
		l.Tail = i.Prev
		i.Prev.Next = nil
	case i.Next != nil && i.Prev != nil:
		i.Prev.Next, i.Next.Prev = i.Next, i.Prev
	}
	l.Length--
}

func (l *LinkedList) Print() {
	if l.Length > 0 {
		ptr := l.Head
		for i := 0; i < l.Length; i++ {
			fmt.Print(ptr.Value, " ")
			ptr = ptr.Next
		}
	}
}

func main() {
	myList := NewLinkedList()
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	myList.PushBack(4)
	myList.PushBack(5)
	myList.Remove(*myList.Last().Prev)
	fmt.Println("List length: ", myList.Len())
	fmt.Println("Head ", myList.First())
	fmt.Println("Tail ", myList.Last())
	myList.Print()
}
