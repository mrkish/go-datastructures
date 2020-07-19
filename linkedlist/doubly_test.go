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

func TestDoublyLinkedList_AddHead(t *testing.T) {
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
	}{
		{
			name: "head added successfully",
			fields: fields{
				Head: &DoubleNode{Value: model.Object{Value: "first"}},
			},
			value: "newFirst",
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
			var err error
			l.AddHead(obj)
			if err == nil && !reflect.DeepEqual(l.Head.Value, obj) {
				t.Errorf("DoublyLinkedList.AddHead() Head doens't match expected | head: %v | expected: %v", l.Head.Value, obj)
			}
		})
	}
}

func TestDoublyLinkedList_AddTail(t *testing.T) {
	type fields struct {
		Current *DoubleNode
		Head    *DoubleNode
		Tail    *DoubleNode
	}
	type args struct {
		obj model.Object
	}
	tests := []struct {
		name      string
		nodes     []*DoubleNode
		nodesHead string
		nodesTail string
		wantHead  string
		wantTail  string
		wantNodes []*DoubleNode
		fields    fields
		value     string
	}{
		{
			name: "tail added successfully",
			fields: fields{
				Tail: &DoubleNode{Value: model.Object{Value: "last"}},
			},
			value: "newLast",
		},
		{
			name: "tail replaces existing successfully, links stay connected",
			nodes: buildNodes([]string{
				"first",
				"last",
			}),
			wantNodes: buildNodes([]string{
				"first",
				"second",
				"third",
			}),
			value: "newLast",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &DoublyLinkedList{
				Current: tt.fields.Current,
				Head:    tt.fields.Head,
				Tail:    tt.fields.Tail,
			}
			if tt.nodes != nil {
				l.addNode(tt.nodes...)
			}

			obj := model.Object{Value: tt.value}
			l.AddTail(obj)
			if !reflect.DeepEqual(l.Tail.Value, obj) {
				t.Errorf("DoublyLinkedList.AddTail() Tail doens't match expected | tail: %v | expected: %v", l.Tail.Value, obj)
			}
			if tt.nodes != nil {
				wantList := &DoublyLinkedList{}
				wantList.addNode(tt.wantNodes...)

				if reflect.DeepEqual(l, wantList) {
					t.Errorf("AddTail() failed to maintain links in list as expected, got: %v want: %v", l, wantList)
				}
			}
		})
	}
}

func TestDoublyLinkedList_addNode(t *testing.T) {
	type fields struct {
		Current *DoubleNode
		Head    *DoubleNode
		Tail    *DoubleNode
	}
	tests := []struct {
		name   string
		fields fields
		values []string
	}{
		{
			name: "list is built with all nodes connected",
			values: []string{
				"first",
				"second",
				"third",
				"fourth",
			},
		},
		{
			name: "existing list without current set",
			fields: fields{
				Current: nil,
				Head: &DoubleNode{
					Value: model.Object{Value: "first"},
				},
			},
			values: []string{
				"second",
				"third",
				"fourth",
			},
		},
		{
			name: "existing list with set head and tail",
			fields: fields{
				Current: nil,
				Head: &DoubleNode{
					Value: model.Object{Value: "first"},
				},
				Tail: &DoubleNode{
					Value: model.Object{Value: "second"},
				},
			},
			values: []string{
				"third",
				"fourth",
			},
		},
		{
			name: "existing list with unset head and tail",
			fields: fields{
				Current: nil,
				Tail: &DoubleNode{
					Value: model.Object{Value: "first"},
				},
			},
			values: []string{
				"second",
				"third",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &DoublyLinkedList{
				Current: tt.fields.Current,
				Head:    tt.fields.Head,
				Tail:    tt.fields.Tail,
			}
			l.addNode(buildNodes(tt.values)...)
			i := 0
			for l.HasNext() {
				var previous model.Object
				var next model.Object
				if l.Current.Previous != nil {
					previous = l.Current.Previous.Value
					wantPrevious := tt.values[i-1]
					if previous.Value != wantPrevious {
						t.Errorf("addNode() mismatched values: previous: %v, expected: %v", previous.Value, wantPrevious)
					}
				}
				if l.Current.Next != nil {
					next = l.Current.Next.Value
					var wantNext interface{}
					if i+1 > len(tt.values) {
						wantNext = tt.values[i+1]
					}
					if next.Value == wantNext {
						t.Errorf("addNode() mismatched values: next: %v, expected: %v", next.Value, wantNext)
					}
				}
				i++
			}
		})
	}
}

func buildNodes(in []string) []*DoubleNode {
	var out []*DoubleNode
	for _, val := range in {
		out = append(out, &DoubleNode{Value: model.Object{Value: val}})
	}
	return out
}
