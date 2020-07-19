package linkedlist

import (
	"go-datastructures/model"
	"testing"
)

// With all of these test cases passing we leverage the method for other tests
// instead of doing weirdness to account for Nodes being pointers.
func TestSinglyLinkedList_Find(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		search string
		values []string
		found  bool
	}{
		{
			name: "value not found",
			values: []string{
				"first",
			},
			search: "last",
			found:  false,
		},
		{
			name: "value is found",
			values: []string{
				"first",
			},
			search: "first",
			found:  true,
		},
		{
			name: "value is found in list with multiple nodes",
			values: []string{
				"first",
				"second",
			},
			search: "second",
			found:  true,
		},
		{
			name: "value is found in list with multiple nodes and set current",
			values: []string{
				"first",
				"second",
			},
			search: "second",
			found:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &SinglyLinkedList{}
			l.addNode(buildSingleNodes(tt.values)...)
			obj := model.Object{Value: tt.search}
			var found bool
			if found = l.Find(obj); found != tt.found {
				t.Errorf("SinglyLinkedList.Find() object not found in list")
			}
		})
	}
}

func TestSinglyLinkedList_Add(t *testing.T) {
	type fields struct {
		Current *Node
		Head    *Node
	}
	type args struct {
		obj model.Object
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "successful add - new head",
			fields: fields{
				Current: nil,
				Head:    nil,
			},
			args: args{
				obj: model.Object{Value: "thing"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &SinglyLinkedList{
				Current: tt.fields.Current,
				Head:    tt.fields.Head,
			}
			l.Add(tt.args.obj)
			if !l.Find(tt.args.obj) {
				t.Errorf("SinglyLinkedList.Add() failure = %v not found after call to Add()", tt.args.obj.Value)
			}
		})
	}
}

func TestSinglyLinkedList_Remove(t *testing.T) {
	type fields struct {
		Current *Node
		Head    *Node
	}
	type args struct {
		obj model.Object
	}
	tests := []struct {
		name      string
		fields    fields
		value     string
		nextValue string
		wantErr   bool
	}{
		{
			name: "node not found: error",
			fields: fields{
				Head: &Node{
					Value: model.Object{Value: "first"},
					Next: &Node{
						Value: model.Object{Value: "second"},
						Next:  nil,
					},
				},
			},
			value:   "third",
			wantErr: true,
		},
		{
			name: "node is removed",
			fields: fields{
				Head: &Node{
					Value: model.Object{Value: "first"},
					Next: &Node{
						Value: model.Object{Value: "second"},
						Next:  nil,
					},
				},
			},
			value:   "second",
			wantErr: false,
		},
		{
			name: "node middle node is removed and links are preserved",
			fields: fields{
				Head: &Node{
					Value: model.Object{Value: "first"},
					Next: &Node{
						Value: model.Object{Value: "second"},
						// It's thirdles all the way down
						Next: &Node{
							Value: model.Object{Value: "third"},
							Next:  nil,
						},
					},
				},
			},
			value:   "second",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &SinglyLinkedList{
				Current: tt.fields.Current,
				Head:    tt.fields.Head,
			}
			obj := model.Object{Value: tt.value}
			nextObj := model.Object{Value: tt.nextValue}
			if err := l.Remove(obj); (err != nil) != tt.wantErr {
				t.Errorf("SinglyLinkedList.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
			if l.Find(obj) {
				t.Errorf("Remove() did not remove item expected; found in list")
			}
			if tt.nextValue != "" {
				if l.Find(nextObj) {
					t.Errorf("Remove() broke the links in the chain; missing next: %v", nextObj)
				}
			}
		})
	}
}

func TestSinglyLinkedList_HasNext(t *testing.T) {
	type fields struct {
		Current *Node
		Head    *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &SinglyLinkedList{
				Current: tt.fields.Current,
				Head:    tt.fields.Head,
			}
			if got := l.HasNext(); got != tt.want {
				t.Errorf("SinglyLinkedList.HasNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func buildSingleNodes(in []string) []*Node {
	var out []*Node
	for _, val := range in {
		out = append(out, &Node{Value: model.Object{Value: val}})
	}
	return out
}
