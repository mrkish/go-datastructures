package queue

import (
	"go-datastructures/linkedlist"
	"go-datastructures/model"
	"reflect"
	"testing"
)

func TestQueue_Dequeue(t *testing.T) {
	tests := []struct {
		name     string
		values   []string
		expected []string
		want     model.Object
		wantErr  bool
	}{
		{
			name: "item is added to existing Queue",
			values: []string{
				"first",
				"second",
			},
			expected: []string{
				"first",
				"second",
				"last",
			},
			want: model.Object{Value: "first"},
		},
		{
			name: "first value is returned and queue is then empty",
			values: []string{
				"first",
			},
			want: model.Object{Value: "first"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &linkedlist.DoublyLinkedList{}
			l.AddNode(linkedlist.BuildDoubleNodes(tt.values)...)
			expected := &linkedlist.DoublyLinkedList{}
			expected.AddNode(linkedlist.BuildDoubleNodes(tt.expected)...)
			q := &Queue{
				List: l,
			}
			got, err := q.Dequeue()
			if (err != nil) != tt.wantErr {
				t.Errorf("Queue.Dequeue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Dequeue() = %v, want %v", got, tt.want)
			}
			if tt.want.Value != "" {
				if q.List.Find(tt.want) {
					t.Error("Stack.Pop() popped item not removed from stack")
				}
			}
		})
	}
}

func TestQueue_Add(t *testing.T) {
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
			name: "item is added to existing Queue",
			values: []string{
				"first",
				"second",
			},
			expected: []string{
				"first",
				"second",
				"last",
			},
			args: args{
				model.Object{Value: "last"},
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
			l := &linkedlist.DoublyLinkedList{}
			l.AddNode(linkedlist.BuildDoubleNodes(tt.values)...)
			expected := &linkedlist.DoublyLinkedList{}
			expected.AddNode(linkedlist.BuildDoubleNodes(tt.expected)...)
			q := &Queue{
				List: l,
			}
			q.Add(tt.args.obj)
			if !reflect.DeepEqual(q.List.Tail.Value, tt.args.obj) {
				t.Error("Add() did not add values as expected")
			}
			// Check that the links are intact
			if q.List.Tail.Next != nil {
				if !reflect.DeepEqual(q.List.Tail.Next.Value, model.Object{Value: tt.expected[len(tt.expected)-1]}) {
					t.Error("Add() did not add values as expected")
				}
			}
		})
	}
}
