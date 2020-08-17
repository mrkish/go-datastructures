package queue

import (
	"errors"
	"go-datastructures/linkedlist"
	"go-datastructures/model"
)

// Queue :: struct :: FILO collection
type Queue struct {
	List *linkedlist.DoublyLinkedList
}

// New :: func :: Returns pointer to a new Queue
func New(values ...string) *Queue {
	l := linkedlist.NewDoublyLinked(values...)
	return &Queue{
		List: l,
	}
}

// Dequeue :: func :: returns the first value in the Queue,
// and removes that value from the embedded LinkedList
func (q *Queue) Dequeue() (model.Object, error) {
	if q.List.Head != nil {
		val := q.List.Head.Value
		return val, q.List.Remove(val)
	}
	return model.Object{}, errors.New("queue is empty")
}

// Add :: func :: Adds a value to the Queue in last position
func (q *Queue) Add(obj model.Object) {
	q.List.AddTail(obj)
}

// Remove :: func :: Removes a value from the Queue
func (q *Queue) Remove(obj model.Object) error {
	return q.List.Remove(obj)
}

// Peek :: func :: Returns the Queue's current value
func (q *Queue) Peek() model.Object {
	return q.List.Current.Value
}
