package bst

import (
	"go-datastructures/model"
	"reflect"
	"testing"
)

// func TestBST_Add(t *testing.T) {
// 	type fields struct {
// 		Root *Node
// 	}
// 	type args struct {
// 		obj model.Object
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		{
// 			name: "sucessful add",
// 			args: args{
// 				obj: model.Object{Value: "first"},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			b := BST{
// 				Root: tt.fields.Root,
// 			}
// 			b.Add(tt.args.obj)
// 			if node, found := b.Find(tt.args.obj); !found {
// 				t.Errorf("Add() failed to find object in BST after Add() call")
// 			} else {
// 				if node.Value.Value != tt.args.obj.Value {
// 					t.Errorf("Add() item added to BST doesn't match added value, got: %s", node.Value.Value)
// 				}
// 			}
// 		})
// 	}
// }

func TestBST_Find(t *testing.T) {
	type fields struct {
		Root *Node
	}
	type args struct {
		obj model.Object
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *Node
		wantFound bool
	}{
		{
			name: "searching for a value that doesn't exist",
			fields: fields{
				Root: &Node{
					Value: model.Object{
						Value: "first",
					},
				},
			},
			args: args{
				model.Object{Value: "second"},
			},
		},
		{
			name: "bst with value at root returns true",
			fields: fields{
				Root: &Node{
					Value: model.Object{
						Value: "first",
					},
				},
			},
			args: args{
				model.Object{Value: "first"},
			},
			want:      &Node{Value: model.Object{Value: "first"}},
			wantFound: true,
		},
		{
			name: "bst with value at the right returns true",
			fields: fields{
				Root: &Node{
					Value: model.Object{
						Value: "first",
					},
					Right: &Node{
						Value: model.Object{Value: "second"},
					},
				},
			},
			args: args{
				model.Object{Value: "second"},
			},
			want:      &Node{Value: model.Object{Value: "second"}},
			wantFound: true,
		},
		{
			name: "bst with value at the left returns true",
			fields: fields{
				Root: &Node{
					Value: model.Object{
						Value: "first",
					},
					Left: &Node{
						Value: model.Object{Value: "two"},
					},
				},
			},
			args: args{
				model.Object{Value: "two"},
			},
			want:      &Node{Value: model.Object{Value: "two"}},
			wantFound: true,
		},
		{
			name: "find goes through multiple levels to find expected match",
			fields: fields{
				Root: &Node{
					Value: model.Object{
						Value: "first",
					},
					Left: &Node{
						Value: model.Object{Value: "two"},
					},
					Right: &Node{
						Value: model.Object{Value: "secondary"},
						Left: &Node{
							Value: model.Object{Value: "seconda"},
						},
					},
				},
			},
			args: args{
				model.Object{Value: "seconda"},
			},
			want:      &Node{Value: model.Object{Value: "seconda"}},
			wantFound: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BST{
				Root: tt.fields.Root,
			}
			got, found := b.Find(tt.args.obj)
			if !reflect.DeepEqual(&got, &tt.want) {
				t.Errorf("BST.Find() got = %v, want %v", got, tt.want)
			}
			if found != tt.wantFound {
				t.Errorf("BST.Find() found = %v, want %v", found, tt.wantFound)
			}
		})
	}
}

func TestNode_Find(t *testing.T) {
	type fields struct {
		Value model.Object
		Left  *Node
		Right *Node
	}
	type args struct {
		parent *Node
		obj    model.Object
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Node{
				Value: tt.fields.Value,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			got, got1 := n.Find(tt.args.parent, tt.args.obj)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.Find() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Node.Find() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
