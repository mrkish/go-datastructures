package stack

import (
	"errors"
	"go-datastructures/linkedlist"
	"go-datastructures/model"
)

type Stack struct {
	List *linkedlist.SinglyLinkedList
}

func (s *Stack) Pop() (model.Object, error) {
	out := s.List.Head.Value
	err := s.List.Remove(out)
	if err != nil {
		return model.Object{}, errors.New("List is empty!")
	}
	return out, nil
}

func (s *Stack) Add(obj model.Object) {
	s.List.Add(obj)
}
