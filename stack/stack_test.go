package stack

import (
	"go-datastructures/linkedlist"
	"go-datastructures/model"
	"reflect"
	"testing"
)

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name     string
		values   []string
		expected []string
		want     model.Object
		wantErr  bool
	}{
		{
			name:    "error due to empty list",
			wantErr: true,
		},
		{
			name: "first item is removed",
			values: []string{
				"first",
				"second",
			},
			expected: []string{
				"second",
			},
			want: model.Object{Value: "first"},
		},
		{
			name: "single item is popped",
			values: []string{
				"first",
			},
			want: model.Object{Value: "first"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &linkedlist.SinglyLinkedList{}
			l.AddNode(linkedlist.BuildSingleNodes(tt.values)...)
			expected := &linkedlist.SinglyLinkedList{}
			expected.AddNode(linkedlist.BuildSingleNodes(tt.expected)...)
			s := &Stack{
				List: l,
			}
			got, err := s.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Stack.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Pop() = %v, want %v", got, tt.want)
			}
			if tt.want.Value != "" {
				if s.List.Find(tt.want) {
					t.Error("Stack.Pop() popped item not removed from stack")
				}
			}
		})
	}
}

func TestStack_Add(t *testing.T) {
	type args struct {
		obj model.Object
	}
	tests := []struct {
		name     string
		values   []string
		expected []string
		args     args
	}{
		{
			name: "item is added to existing Stack",
			values: []string{
				"first",
				"second",
			},
			expected: []string{
				"newFirst",
				"first",
				"second",
			},
			args: args{
				model.Object{Value: "newFirst"},
			},
		},
		{
			name: "item is added to a new stack",
			expected: []string{
				"first",
			},
			args: args{
				model.Object{Value: "first"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &linkedlist.SinglyLinkedList{}
			l.AddNode(linkedlist.BuildSingleNodes(tt.values)...)
			expected := &linkedlist.SinglyLinkedList{}
			expected.AddNode(linkedlist.BuildSingleNodes(tt.expected)...)
			s := &Stack{
				List: l,
			}
			s.Add(tt.args.obj)
			if !reflect.DeepEqual(s.List.Head.Value, tt.args.obj) {
				t.Error("Add() did not add values as expected")
			}
			// Check that the links are intact
			if s.List.Head.Next != nil {
				if !reflect.DeepEqual(s.List.Head.Next.Value, model.Object{Value: tt.expected[1]}) {
					t.Error("Add() did not add values as expected")
				}
			}
		})
	}
}
