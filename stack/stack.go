package stack

import (
	"errors"
	"go-datastructures/linkedlist"
	"go-datastructures/model"
)

// Stack :: struct :: FIFO collection
type Stack struct {
	List *linkedlist.SinglyLinkedList
}

// Pop :: func :: returns the first value in the Stack,
// and removes that value from the embedded LinkedList
func (s *Stack) Pop() (model.Object, error) {
	if s.List.Head != nil {
		val := s.List.Head.Value
		return val, s.List.Remove(val)
	}
	return model.Object{}, errors.New("list is empty!")
}

// Add :: func :: Adds a value to the Stack in first position
func (s *Stack) Add(obj model.Object) {
	s.List.Add(obj)
}
