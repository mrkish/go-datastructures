package linkedlist

import (
	"errors"

	"go-datastructures/model"
)

// DoublyLinkedList - struct - DoublyLinkedList
type DoublyLinkedList struct {
	Current *DoubleNode
	Head    *DoubleNode
	Tail    *DoubleNode
}

// DoubleNode :: struct :: Container struct for model.Object nodes
type DoubleNode struct {
	Value    model.Object
	Next     *DoubleNode
	Previous *DoubleNode
}

// AddHead :: func :: Adds a new node to the LinkedList
// - Previous Head still points towards it's own Next
// - Previous Head.Previous points to new Head
// - New Node being added to the front has the Next point towards the previous Head
// - List then accepts the new node as the current Head
func (l *DoublyLinkedList) AddHead(obj model.Object) error {
	newItem := &DoubleNode{
		Value: obj,
		Next:  l.Head,
	}
	l.Head.Previous = newItem
	l.Head = newItem
	return nil
}

// AddTail :: func :: Adds a new node to the LinkedList
// - Previous Tail's Next is updated to new Node (in constructor)
// - New Node.Previous points back to the old Tail
// - LinkedList updates current tail as the new Node
func (l *DoublyLinkedList) AddTail(obj model.Object) error {
	newItem := &DoubleNode{
		Value:    obj,
		Previous: l.Tail,
	}
	// Looks weird, but updating the Next reference to to previous tail
	// before the List's tail is actually updated.
	l.Tail.Next = newItem
	// Update the List's Tail to be the new Node
	l.Tail = newItem
	return nil
}

// Find :: func :: find an object in the list
func (l *DoublyLinkedList) Find(obj model.Object) (*DoubleNode, bool) {
	current := l.Head
	for current != nil {
		if current.Value == obj {
			return current, true
		}
		current = current.Next
	}
	return nil, false
}

// Remove :: func :: find an object in the list
// This implementation gets to be simpler because the reference to the Previous
// is kept in the DoubleNode struct.
func (l *DoublyLinkedList) Remove(obj model.Object) error {
	if node, found := l.Find(obj); !found {
		return errors.New("object not in list")
	} else {
		node.Previous.Next = node.Next
		node.Next.Previous = node.Previous
	}
	return nil
}
