package linkedlist

import (
	"go-datastructures/model"
	"reflect"
	"testing"
)

func TestDoublyLinkedList_Find(t *testing.T) {
	type args struct {
		obj model.Object
	}
	tests := []struct {
		name   string
		search string
		values []string
		found  bool
	}{
		{
			name:   "empty list returns false",
			search: "last",
			found:  false,
		},
		{
			name: "value not found",
			values: []string{
				"first",
			},
			search: "last",
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
			name: "first value is found in list with multiple nodes",
			values: []string{
				"first",
				"second",
			},
			search: "first",
			found:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &DoublyLinkedList{}
			l.AddNode(BuildDoubleNodes(tt.values)...)
			obj := model.Object{Value: tt.search}
			_, found := l.Find(obj)
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
			l.AddHead(obj)
			if !reflect.DeepEqual(l.Head.Value, obj) {
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
			nodes: BuildDoubleNodes([]string{
				"first",
				"last",
			}),
			wantNodes: BuildDoubleNodes([]string{
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
				l.AddNode(tt.nodes...)
			}

			obj := model.Object{Value: tt.value}
			l.AddTail(obj)
			if !reflect.DeepEqual(l.Tail.Value, obj) {
				t.Errorf("DoublyLinkedList.AddTail() Tail doens't match expected | tail: %v | expected: %v", l.Tail.Value, obj)
			}
			if tt.nodes != nil {
				wantList := &DoublyLinkedList{}
				wantList.AddNode(tt.wantNodes...)

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
			l.AddNode(BuildDoubleNodes(tt.values)...)
			i := 0
			for l.HasNext() {
				var previous model.Object
				var next model.Object
				if l.Current.Previous != nil {
					previous = l.Current.Previous.Value
					wantPrevious := tt.values[i-1]
					if previous.Value != wantPrevious {
						t.Errorf("AddNode() mismatched values: previous: %v, expected: %v", previous.Value, wantPrevious)
					}
				}
				if l.Current.Next != nil {
					next = l.Current.Next.Value
					var wantNext interface{}
					if i+1 > len(tt.values) {
						wantNext = tt.values[i+1]
					}
					if next.Value == wantNext {
						t.Errorf("AddNode() mismatched values: next: %v, expected: %v", next.Value, wantNext)
					}
				}
				i++
			}
		})
	}
}

func TestDoublyLinkedList_Remove(t *testing.T) {
	type fields struct {
		Current *DoubleNode
		Head    *DoubleNode
		Tail    *DoubleNode
	}
	type args struct {
		obj model.Object
	}
	tests := []struct {
		name         string
		values       []string
		expectedList []string
		fields       fields
		args         args
		wantErr      bool
	}{
		{
			name: "empty list returns error",
			args: args{
				obj: model.Object{Value: "second"},
			},
			wantErr: true,
		},
		{
			name: "middle link is removed",
			values: []string{
				"first",
				"second",
				"third",
			},
			expectedList: []string{
				"first",
				"third",
			},
			args: args{
				obj: model.Object{Value: "second"},
			},
		},
		{
			name: "first link is removed",
			values: []string{
				"first",
				"second",
				"third",
			},
			expectedList: []string{
				"second",
				"third",
			},
			args: args{
				obj: model.Object{Value: "first"},
			},
		},
		{
			name: "last link is removed",
			values: []string{
				"first",
				"second",
			},
			expectedList: []string{
				"first",
			},
			args: args{
				obj: model.Object{Value: "second"},
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
			expected := &DoublyLinkedList{}
			l.AddNode(BuildDoubleNodes(tt.values)...)
			expected.AddNode(BuildDoubleNodes(tt.expectedList)...)
			if err := l.Remove(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("DoublyLinkedList.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
			_, found := l.Find(tt.args.obj)
			if found {
				t.Errorf("DobulyLinkedList.Remove() item not removed from list!")
			}
			if !reflect.DeepEqual(l, expected) {
				t.Errorf("DobulyLinkedList.Remove() list contains unexpected values")
			}
		})
	}
}

func TestDoublyLinkedList_HasNext(t *testing.T) {
	type fields struct {
		Current *DoubleNode
		Head    *DoubleNode
		Tail    *DoubleNode
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "has next",
			fields: fields{
				Head: &DoubleNode{
					Value: model.Object{Value: "first"},
					Next: &DoubleNode{
						Value: model.Object{Value: "second"},
					},
				},
			},
			want: true,
		},
		{
			name: "no next",
			fields: fields{
				Head: &DoubleNode{
					Value: model.Object{Value: "first"},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &DoublyLinkedList{
				Current: tt.fields.Current,
				Head:    tt.fields.Head,
				Tail:    tt.fields.Tail,
			}
			if got := l.HasNext(); got != tt.want {
				t.Errorf("DoublyLinkedList.HasNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoublyLinkedList_HasPrevious(t *testing.T) {
	type fields struct {
		Current *DoubleNode
		Head    *DoubleNode
		Tail    *DoubleNode
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "has previous",
			fields: fields{
				Head: &DoubleNode{
					Value: model.Object{Value: "first"},
					Next: &DoubleNode{
						Value: model.Object{Value: "second"},
					},
				},
				Tail: &DoubleNode{
					Value: model.Object{Value: "second"},
					Previous: &DoubleNode{
						Value: model.Object{Value: "first"},
					},
				},
			},
			want: true,
		},
		{
			name: "no previous",
			fields: fields{
				Head: &DoubleNode{
					Value: model.Object{Value: "first"},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &DoublyLinkedList{
				Current: tt.fields.Current,
				Head:    tt.fields.Head,
				Tail:    tt.fields.Tail,
			}
			if got := l.HasPrevious(); got != tt.want {
				t.Errorf("DoublyLinkedList.HasPrevious() = %v, want %v", got, tt.want)
			}
		})
	}
}
