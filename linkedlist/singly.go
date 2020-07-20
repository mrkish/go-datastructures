package linkedlist

import (
	"errors"
	"reflect"

	"go-datastructures/model"
)

// SinglyLinkedList :: struct :: Singly-Linked LinkedList
type SinglyLinkedList struct {
	Current *Node
	Head    *Node
}

// Node :: struct :: Container struct for model.Object nodes
type Node struct {
	Value model.Object
	Next  *Node
}

// NewSinglyLinked :: func :: Returns a pointer to a new SinglyLinkedList
func NewSinglyLinked(values ...string) (l *SinglyLinkedList) {
	l = &SinglyLinkedList{}
	if values != nil || len(values) > 0 {
		l.AddNode(BuildSingleNodes(values)...)
	}
	return l
}

// Add :: func :: Adds a new node to the LinkedList at the Head
// - Previous Head still points towards it's own Next
// - New Node being added to the front has the Next point towards the previous Head (in constructor)
// - List then accepts the new node as the current Head
func (l *SinglyLinkedList) Add(obj model.Object) {
	newItem := &Node{
		Value: obj,
		Next:  l.Head,
	}
	l.Head = newItem
	// l.Current = newItem
}

// Find :: func :: Find an object in the list
func (l *SinglyLinkedList) Find(obj model.Object) bool {
	l.Current = l.Head
	if l.Current == nil {
		return false
	}
	// Check first element manually, since HasNext will advance Current
	if firstMatch := reflect.DeepEqual(l.Current.Value, obj); firstMatch {
		return firstMatch
	}
	// Iterate through rest of list
	for l.HasNext() {
		if reflect.DeepEqual(l.Current.Value, obj) {
			return true
		}
	}
	return false
}

// Remove :: func :: Remove an object from the list
func (l *SinglyLinkedList) Remove(obj model.Object) error {
	l.Current = l.Head
	var previous *Node
	for l.Current != nil {
		if reflect.DeepEqual(l.Current.Value, obj) {
			if previous != nil {
				previous.Next = l.Current.Next
			} else {
				l.Head = nil
				l.Current = nil
			}
			return nil
		}
		previous = l.Current
		l.Current = l.Current.Next
	}
	return errors.New("object not found in list")
}

// HasNext :: func :: returns true if the next Node is not nil
// Since this is being use to iterate over lists, it also
// advances the Current marker.
func (l *SinglyLinkedList) HasNext() bool {
	// Check if Current isn't set
	if l.Current == nil {
		l.Current = l.Head
	}
	current := l.Current
	// Advance Current if Next isn't nil
	if l.Current != nil && l.Current.Next != nil {
		l.Current = l.Current.Next
	}
	return current.Next != nil
}

// AddNode :: func :: Helper function to build list or add new nodes to existing list
func (l *SinglyLinkedList) AddNode(n ...*Node) {
	// Determine position in list before iterating
	if l.Current == nil {
		l.Current = l.Head
	}
	if l.Current != nil {
		// Advance to last link
		for l.HasNext() {

		}
	}
	// Current has been set, iterate to add Nodes
	for i, node := range n {
		if i == 0 && l.Current == nil {
			// Set current if list is empty
			l.Current = node
			l.Head = node
			continue
		} else {
			l.Current.Next = node
			l.Current = node
		}
	}
}

// BuildSingleNodes :: func :: Helper function to build Nodes
func BuildSingleNodes(in []string) []*Node {
	var out []*Node
	for _, val := range in {
		out = append(out, &Node{Value: model.Object{Value: val}})
	}
	return out
}
