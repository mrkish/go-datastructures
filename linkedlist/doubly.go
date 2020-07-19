package linkedlist

import (
	"errors"
	"reflect"

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
func (l *DoublyLinkedList) AddHead(obj model.Object) {
	newItem := &DoubleNode{
		Value: obj,
		Next:  l.Head,
	}
	oldHead := l.Head
	if oldHead != nil {
		l.Head.Previous = newItem
		if oldHead.Next == l.Tail {
			l.Tail = l.Head
		}
	}
	l.Head = newItem
	// If this is the only item, it's also the tail
	if l.Tail == nil {
		l.Tail = newItem
	}
}

// AddTail :: func :: Adds a new node to the LinkedList
// - Previous Tail's Next is updated to new Node (in constructor)
// - New Node.Previous points back to the old Tail
// - LinkedList updates current tail as the new Node
func (l *DoublyLinkedList) AddTail(obj model.Object) {
	newItem := &DoubleNode{
		Value:    obj,
		Previous: l.Tail,
	}
	// Looks weird, but updating the Next reference to to previous tail
	// before the List's tail is actually updated.
	if l.Tail != nil {
		l.Tail.Next = newItem
	}
	// Update the List's Tail to be the new Node
	l.Tail = newItem
}

// Find :: func :: find an object in the list
func (l *DoublyLinkedList) Find(obj model.Object) (model.Object, bool) {
	current := l.Head
	for current != nil {
		if reflect.DeepEqual(current.Value, obj) {
			return current.Value, true
		}
		current = current.Next
	}
	return model.Object{}, false
}

// Remove :: func :: find an object in the list
// This implementation gets to be simpler because the reference to the Previous
// is kept in the DoubleNode struct.
func (l *DoublyLinkedList) Remove(obj model.Object) error {
	_, found := l.Find(obj)
	if !found {
		return errors.New("object not in list")
	}
	node := l.Current
	node.Previous.Next = node.Next
	node.Next.Previous = node.Previous
	return nil
}

// HasNext :: func :: returns true if the next Node is not nil
func (l *DoublyLinkedList) HasNext() bool {
	// Check Current/Head to verify the list has Nodes
	// Check if Current isn't set
	if l.Current == nil {
		// If true, then check if Head is nil
		if l.Head != nil {
			l.Current = l.Head
		} else {
			// No Head or Current, list must be empty
			return false
		}
	}
	return l.Current.Next != nil
}

// HasPrevious :: func :: returns true if the previous Node is not nil
func (l *DoublyLinkedList) HasPrevious() bool {
	// Check Current/Tail to verify the list has Nodes
	// Check if Current isn't set
	if l.Current == nil {
		// If true, then check if Head is nil
		if l.Tail != nil {
			l.Current = l.Tail
		} else {
			// No Tail or Current, list must be empty
			return false
		}
	}
	return l.Current.Previous != nil
}

// Helper function to build list or add new nodes to existing list
func (l *DoublyLinkedList) addNode(n ...*DoubleNode) {
	// Determine position in list before iterating
	if l.Current == nil && l.Head != nil {
		l.Current = l.Head
	} else if l.Head != nil && l.Current == nil {
		l.Head = l.Current
	} else if l.Head == nil && l.Tail != nil {
		l.Head = l.Tail
		l.Current = l.Head
	} else if l.Head != nil && l.Tail != nil {
		// Start at the end of the list
		l.Current = l.Tail
	}
	for i, node := range n {
		if i == 0 && l.Current == nil {
			// Set current if list is empty
			l.Current = node
			continue
		} else {
			node.Previous = l.Current
			l.Current.Next = node
			l.Current = node
		}
		// Set Head if first node and Head is unset
		if i == 0 && l.Head == nil {
			l.Head = node
		}
		// Set Tail if last node in the list
		// We're not adding anywhere at the end;
		// what we've added must be the Tail
		if i == len(n)-1 {
			l.Tail = node
		}
	}
}
