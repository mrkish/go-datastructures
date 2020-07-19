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

// Dequeue :: func :: returns the first value in the Queue,
// and removes that value from the embedded LinkedList
func (s *Queue) Dequeue() (model.Object, error) {
	if s.List.Head != nil {
		val := s.List.Head.Value
		return val, s.List.Remove(val)
	}
	return model.Object{}, errors.New("queue is empty!")
}

// Add :: func :: Adds a value to the Queue in last position
func (s *Queue) Add(obj model.Object) {
	s.List.AddTail(obj)
}
