package linkedlist

import (
	"go-datastructures/model"
	"reflect"
	"testing"
)

func TestDoublyLinkedList_Find(t *testing.T) {
	type fields struct {
		Current *DoubleNode
		Head    *DoubleNode
		Tail    *DoubleNode
	}
	type args struct {
		obj model.Object
	}
	tests := []struct {
		name   string
		fields fields
		value  string
		want   string
		found  bool
	}{
		{
			name: "value not found",
			fields: fields{
				Head: &DoubleNode{
					Value: model.Object{Value: "first"},
				},
			},
			value: "last",
			found: false,
		},
		{
			name: "value is found",
			fields: fields{
				Head: &DoubleNode{
					Value: model.Object{Value: "first"},
				},
			},
			value: "first",
			found: true,
		},
		{
			name: "value is found in list with multiple nodes",
			fields: fields{
				Head: &DoubleNode{
					Value: model.Object{Value: "first"},
					Next: &DoubleNode{
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
			l := &DoublyLinkedList{
				Current: tt.fields.Current,
				Head:    tt.fields.Head,
				Tail:    tt.fields.Tail,
			}
			obj := model.Object{Value: tt.value}
			got, found := l.Find(obj)
			if tt.found && !reflect.DeepEqual(got, obj) {
				t.Errorf("DoublyLinkedList.Find() got = %v, want %v", got, obj)
			}
			if found != tt.found {
				t.Error("DoublyLinkedList.Find() item not found in list")
			}
		})
	}
}
