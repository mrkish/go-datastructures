package linkedlist

import (
	"errors"
	"go-datastructures/model"
)

// SinglyLinkedList - struct - Singly-Linked LinkedList
type SinglyLinkedList struct {
	Current *Node
	Head    *Node
}

// Node :: struct :: Container struct for *model.Object nodes
type Node struct {
	Value *model.Object
	Next  *Node
}

// Add :: func :: Adds a new node to the LinkedList at the Head
// Previous Head still points towards it's own Next
// New Node being added to the front has the Next point towards the previous Head (in constructor)
// List then accepts the new node as the current Head
func (l *SinglyLinkedList) Add(obj *model.Object) error {
	newItem := &Node{
		Value: obj,
		Next:  l.Head,
	}
	l.Head = newItem
	return nil
}

// Find -- func -- find an object in the list
func (l *SinglyLinkedList) Find(obj *model.Object) bool {
	current = l.Head
	for current != nil {
		if current.Value == obj {
			return true
		}
		current = current.Next
	}
}

// DoublyLinkedList - struct - DoublyLinkedList
type DoublyLinkedList struct {
	Head    *DoubleNode
	Tail    *DoubleNode
	Current *DoubleNode
}

// DoubleNode :: struct :: Container struct for *model.Object nodes
type DoubleNode struct {
	Value    *model.Object
	Next     *DoubleNode
	Previous *DoubleNode
}

// AddHead :: func :: Adds a new node to the LinkedList
// Previous Head still points towards it's own Next
// Previous Head.Previous points to new Head
// New Node being added to the front has the Next point towards the previous Head
// List then accepts the new node as the current Head
func (l *DoublyLinkedList) AddHead(obj *model.Object) error {
	newItem := &DoubleNode{
		Value: obj,
		Next:  l.Head,
	}
	l.Head.Previous = newItem
	l.Head = newItem
	return nil
}

// AddTail :: func :: Adds a new node to the LinkedList
// Previous Tail's Next is updated to new Node (in constructor)
// New Node.Previous points back to the old Tail
// LinkedList updates current tail as the new Node
func (l *DoublyLinkedList) AddTail(obj *model.Object) error {
	newItem := &DoubleNode{
		Value:    obj,
		Previous: l.Tail,
	}
	l.Tail.Next = newItem
	l.Tail = newItem
	return nil
}

// Find -- func -- find an object in the list
func (l *SinglyLinkedList) Find(obj *model.Object) bool {
	current = l.Head
	for current != nil {
		if current.Value == obj {
			return true
		}
		current = current.Next
	}
}

// Find -- func -- find an object in the list
func (l *DoublyLinkedList) Find(obj *model.Object) bool {
	current = l.Head
	for current != nil {
		if current.Value == obj {
			return true
		}
		current = current.Next
	}
}

// Remove -- func -- find an object in the list
func (l *DoublyLinkedList) Remove(obj *model.Object) error {
	if !l.Find(obj) {
		return errors.New("object not in list")
	}
	current = l.Head
	for current != nil {
		if current.Value == obj {
			return true
		}
		current = current.Next
	}
}
