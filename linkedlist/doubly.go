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

// NewDoublyLinked :: func :: Returns a pointer to a new DoublyLinkedList
func NewDoublyLinked(values ...string) (l *DoublyLinkedList) {
	l = &DoublyLinkedList{}
	if values != nil || len(values) > 0 {
		l.AddNode(BuildDoubleNodes(values)...)
	}
	return l
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
func (l *DoublyLinkedList) Find(obj model.Object) bool {
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

// Remove :: func :: find an object in the list
// This implementation gets to be simpler because the reference to the Previous
// is kept in the DoubleNode struct.
func (l *DoublyLinkedList) Remove(obj model.Object) error {
	if found := l.Find(obj); !found {
		return errors.New("object not in list")
	}
	node := l.Current
	if node == l.Head {
		// Removing the Head
		l.Head = node.Next
		if l.Head != nil {
			l.Head.Previous = nil
		} else {
			// This was the last item in the list
			l.Tail = nil
		}
	} else if node == l.Tail {
		// Removing the Tail
		l.Tail = node.Previous
		l.Tail.Next = nil
	} else {
		// Removing a link
		node.Previous.Next = node.Next
		node.Next.Previous = node.Previous
	}
	l.Current = nil
	return nil
}

// HasNext :: func :: returns true if the next Node is not nil
func (l *DoublyLinkedList) HasNext() bool {
	if l.Current == nil {
		l.Current = l.Head
		if l.Current != nil {
			return l.Current.Next != nil
		}
	}
	current := l.Current
	// Advance Current if Next isn't nil
	if l.Current.Next != nil {
		l.Current = l.Current.Next
	}
	return current.Next != nil
}

// HasPrevious :: func :: returns true if the previous Node is not nil
func (l *DoublyLinkedList) HasPrevious() bool {
	// Check Current/Tail to verify the list has Nodes
	// Check if Current isn't set
	if l.Current == nil {
		// If true, then check if Head is nil
		if l.Tail != nil {
			l.Current = l.Tail
		} else if l.Head != nil {
			l.Current = l.Head
			l.Tail = l.Head
		} else {
			// No Tail or Current, list must be empty
			return false
		}
	}
	return l.Current.Previous != nil
}

// AddNode :: func :: Helper function to build list or add new nodes to existing list
func (l *DoublyLinkedList) AddNode(n ...*DoubleNode) {
	// Determine position in list before iterating
	if l.Head != nil && l.Tail != nil {
		l.checkHeadTail()
		l.Current = l.Tail
		// If Head and Tail are set but not connected
	} else if l.Current == nil && l.Head != nil {
		l.checkHeadTail()
		// Advance to last item in list
		for l.HasNext() {
		}
		if l.Tail == nil {
			l.Tail = l.Current
			l.Current = l.Tail
		}
	} else if l.Head == nil && l.Tail != nil {
		l.Head = l.Tail
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
			node.Previous = l.Current
			l.Current.Next = node
			l.Current = node
		}
	}
	l.Tail = l.Current
}

// BuildDoubleNodes :: func :: Helper function to wrap values into Nodes
func BuildDoubleNodes(in []string) []*DoubleNode {
	var out []*DoubleNode
	for _, val := range in {
		out = append(out, &DoubleNode{Value: model.Object{Value: val}})
	}
	return out
}

func (l DoublyLinkedList) checkHeadTail() {
	if l.Head != nil && l.Head.Next == nil && l.Tail != nil {
		l.Head.Next = l.Tail
		l.Tail.Previous = l.Head
	}
}
