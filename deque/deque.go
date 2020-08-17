package deque

import (
	"errors"
	"go-datastructures/linkedlist"
	"go-datastructures/model"
)

// Deque :: struct :: FILO collection
type Deque struct {
	List *linkedlist.DoublyLinkedList
}

// New :: func :: Returns pointer to a new Deque
func New(values ...string) *Deque {
	l := linkedlist.NewDoublyLinked(values...)
	return &Deque{
		List: l,
	}
}

// Dequeue :: func :: returns the first value in the Queue,
// and removes that value from the embedded LinkedList
func (d *Deque) Dequeue() (model.Object, error) {
	if d.List.Head != nil {
		val := d.List.Head.Value
		return val, d.List.Remove(val)
	}
	return model.Object{}, errors.New("deque is empty")
}

// AddFirst :: func :: Adds a value to the Deque in first position
func (d *Deque) AddFirst(obj model.Object) {
	d.List.AddHead(obj)
}

// AddLast :: func :: Adds a value to the Deque in last position
func (d *Deque) AddLast(obj model.Object) {
	d.List.AddTail(obj)
}

// Remove :: func :: Removes a value from the Queue
func (d *Deque) Remove(obj model.Object) error {
	return d.List.Remove(obj)
}

// PeekFirst :: func :: Returns the Deque's current value
func (d *Deque) PeekFirst() model.Object {
	return d.List.Current.Value
}
