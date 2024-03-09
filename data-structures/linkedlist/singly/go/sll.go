package main

import (
	"fmt"

	"github.com/pkg/errors"
)

// SllNode is an element in the linked list
type SllNode struct {
	element interface{}
	next    *SllNode
}

func (n *SllNode) String() string {
	if n == nil {
		return "nil"
	}
	return fmt.Sprintf("%s", n.element)
}

// NewSllNode returns a new node
func NewSllNode(element interface{}) *SllNode {
	return &SllNode{
		element: element,
		next:    nil,
	}
}

// SinglyLinkedList is a unidirectional linked list
type SinglyLinkedList struct {
	head *SllNode
	tail *SllNode
	size int
}

// NewSinglyLinkedList returns a new singly-linked list
func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

// AddFirst adds an element to the front of the list so that it becomes the new head node
func (l *SinglyLinkedList) AddFirst(element interface{}) {
	newHead := NewSllNode(element)
	newHead.next = l.head
	l.head = newHead

	if l.tail == nil {
		l.tail = l.head
	}

	l.size++
}

// AddLast adds an element to the back of the list so that it becomes the new tail node
func (l *SinglyLinkedList) AddLast(element interface{}) {
	node := l.tail

	node.next = NewSllNode(element)
	l.tail = node.next

	l.size++
}

// AddAtPosition adds an element to a specified position in the list
func (l *SinglyLinkedList) AddAtPosition(n int, element interface{}) error {
	if n < 1 || n > l.size+1 {
		return errors.New("no such position")
	}

	if n == 1 {
		l.AddFirst(element)
		return nil
	}

	prev, node := l.head, l.head
	for i := 1; i < n; i++ {
		prev = node
		node = node.next
	}

	newNode := NewSllNode(element)
	newNode.next = node
	prev.next = newNode

	l.size++
	return nil
}

// RemoveFirst removes the head of the list and returns it
func (l *SinglyLinkedList) RemoveFirst() (interface{}, error) {
	if l.head == nil {
		return nil, errors.New("no such element")
	}

	node := l.head
	if node.next == nil {
		l.head = nil
		l.tail = nil
		l.size = 0
		return node.element, nil
	}

	l.head = l.head.next
	l.size--
	return node.element, nil
}

// RemoveLast removes the tail of the list and returns it
func (l *SinglyLinkedList) RemoveLast() (interface{}, error) {
	if l.tail == nil {
		return nil, errors.New("no such element")
	}

	if l.head.next == nil {
		node := l.tail
		l.head = nil
		l.tail = nil
		l.size = 0
		return node.element, nil
	}

	prev := l.head
	node := l.head
	for node.next != nil {
		prev = node
		node = node.next
	}

	prev.next = nil
	l.tail = prev
	l.size--
	return node.element, nil
}

// RemoveAtPosition removes the element at the specified position in the list and returns it
func (l *SinglyLinkedList) RemoveAtPosition(n int) (interface{}, error) {
	if n < 1 || n > l.size+1 {
		return nil, errors.New("no such position")
	}

	if n == 1 {
		return l.RemoveFirst()
	}

	if n == l.size {
		return l.RemoveLast()
	}

	prev := l.head
	node := l.head
	for i := 1; i < n; i++ {
		prev = node
		node = node.next
	}

	prev.next = node.next
	l.size--
	return node.element, nil
}

func (l SinglyLinkedList) String() string {
	var str string
	curr := l.head
	str += fmt.Sprint(curr)
	if curr != nil {
		for curr.next != nil {
			curr = curr.next
			str += fmt.Sprintf("->%s", curr)
		}
	}
	return str
}

func main() {
	list := NewSinglyLinkedList()
	fmt.Println(list)

	list.AddFirst("1")
	fmt.Println(list)
	list.AddLast("2")
	fmt.Println(list)
	list.AddAtPosition(2, "3")
	fmt.Println(list)

	list.RemoveAtPosition(2)
	fmt.Println(list)
	list.RemoveFirst()
	fmt.Println(list)
	list.RemoveLast()
	fmt.Println(list)

	// nil
	// 1
	// 1->2
	// 1->3->2
	// 1->2
	// 2
	// nil
}
