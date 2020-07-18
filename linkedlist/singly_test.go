package linkedlist

import (
	"go-datastructures/model"
	"reflect"
	"testing"
)

// With all of these test cases passing we leverage the method for other tests
// instead of doing weirdness to account for Nodes being pointers.
func TestSinglyLinkedList_Find(t *testing.T) {
	type fields struct {
		Current *Node
		Head    *Node
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		value  string
		found  bool
	}{
		{
			name: "value not found",
			fields: fields{
				Head: &Node{
					Value: model.Object{Value: "first"},
				},
			},
			value: "last",
			found: false,
		},
		{
			name: "value is found",
			fields: fields{
				Head: &Node{
					Value: model.Object{Value: "first"},
				},
			},
			value: "first",
			found: true,
		},
		{
			name: "value is found in list with multiple nodes",
			fields: fields{
				Head: &Node{
					Value: model.Object{Value: "first"},
					Next: &Node{
						Value: model.Object{Value: "second"},
					},
				},
			},
			value: "second",
			found: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &SinglyLinkedList{
				Current: tt.fields.Current,
				Head:    tt.fields.Head,
			}
			// Instead of duplicating the literal struct in all the test cases
			obj := model.Object{Value: tt.value}
			var got model.Object
			var found bool
			if got, found = l.Find(obj); found != tt.found {
				t.Errorf("SinglyLinkedList.Find() = %v, want %v", got, tt.found)
			}
			if !reflect.DeepEqual(got, obj) {
				t.Errorf("SinglyLinkedList.Find() = %v, want %v", got, obj)
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
			_, found := l.Find(tt.args.obj)
			if !found {
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
			_, found := l.Find(obj)
			if found {
				t.Errorf("Remove() did not remove item expected; found in list")
			}
			if tt.nextValue != "" {
				_, foundNext := l.Find(nextObj)
				if !foundNext {
					t.Errorf("Remove() broke the links in the chain; missing next: %v", nextObj)
				}
			}
		})
	}
}
