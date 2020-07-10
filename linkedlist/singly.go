package linkedlist

import (
	"errors"
	"reflect"

	"go-datastructures/model"
)

// SinglyLinkedList - struct - Singly-Linked LinkedList
type SinglyLinkedList struct {
	Current *Node
	Head    *Node
}

// Node :: struct :: Container struct for model.Object nodes
type Node struct {
	Value model.Object
	Next  *Node
}

// Add :: func :: Adds a new node to the LinkedList at the Head
// - Previous Head still points towards it's own Next
// - New Node being added to the front has the Next point towards the previous Head (in constructor)
// - List then accepts the new node as the current Head
func (l *SinglyLinkedList) Add(obj model.Object) error {
	newItem := &Node{
		Value: obj,
		Next:  l.Head,
	}
	l.Head = newItem
	return nil
}

// Find :: func :: find an object in the list
func (l *SinglyLinkedList) Find(obj model.Object) bool {
	current := l.Head
	for current != nil {
		if reflect.DeepEqual(current.Value, obj) {
			return true
		}
		current = current.Next
	}
	return false
}

// Remove :: func :: find an object in the list
// Unfortunately since the SinglyLinkedList doesn't have Previous references,
// the code for the Remove method must re-implement the logic for Find to keep
// the previous node so that reference can be updated to the found node's Next.
func (l *SinglyLinkedList) Remove(obj model.Object) error {
	current := l.Head
	var previous *Node
	for current != nil {
		if reflect.DeepEqual(current.Value, obj) {
			// Set the previous node to point ahead of the current position
			previous.Next = current.Next
			return nil
		}
		previous = current
		current = current.Next
	}
	return errors.New("object not found")
}
