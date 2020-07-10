package stack

import (
	"errors"
	"go-datastructures/model"
)

type Stack struct {
	List *SinglyLinkedList
}

func (s *Stack) Pop() model.Object {
	out := s.List.Head.Value
	err := s.List.Remove(out)
	if err != nil {
		return errors.New("List is empty!")
	}
	return out
}

func (s *Stack) Add(obj model.Object) {
	s.List.Add(obj)
}
