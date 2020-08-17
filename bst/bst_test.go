package bst

import (
	"go-datastructures/model"
	"reflect"
	"testing"
)

func TestBST_Add(t *testing.T) {
	type fields struct {
		Root *Node
	}
	type args struct {
		obj model.Object
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		expectedLeft  *model.Object
		expectedRight *model.Object
	}{
		{
			name: "first call to add creates a root if it is not present",
			args: args{
				obj: model.Object{Value: "first"},
			},
		},
		{
			name: "call to Add() with a non-nil root saves the value in Root",
			fields: fields{
				Root: &Node{},
			},
			args: args{
				obj: model.Object{Value: "first"},
			},
		},
		{
			name: "call to Add() places left node correctly",
			fields: fields{
				Root: &Node{
					Value: model.Object{Value: "first"},
				},
			},
			args: args{
				obj: model.Object{Value: "two"},
			},
			expectedLeft: &model.Object{Value: "two"},
		},
		{
			name: "call to Add() places right node correctly",
			fields: fields{
				Root: &Node{
					Value: model.Object{Value: "first"},
				},
			},
			args: args{
				obj: model.Object{Value: "quaternary"},
			},
			expectedRight: &model.Object{Value: "quaternary"},
		},
		{
			name: "call to Add() places right node correctly",
			fields: fields{
				Root: &Node{
					Value: model.Object{Value: "primary"},
				},
			},
			args: args{
				obj: model.Object{Value: "two"},
			},
			expectedLeft: &model.Object{Value: "two"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BST{
				Root: tt.fields.Root,
			}
			b.Add(tt.args.obj)
			if node, found := b.Find(tt.args.obj); !found {
				t.Errorf("Add() failed to find object in BST after Add() call")
			} else {
				if node.Value.Value != tt.args.obj.Value {
					t.Errorf("Add() item added to BST doesn't match added value, got: %s", node.Value.Value)
				}
			}
			if tt.expectedRight == nil && b.Root.Right != nil {
				t.Errorf("Add() unexpected Right leaf")
			}
			if (tt.expectedRight != nil && b.Root.Right != nil) && b.Root.Right.Value.Value != tt.expectedRight.Value {
				t.Errorf("Add() Right value does not match expected, got %s, want %s", b.Root.Right.Value.Value, tt.expectedRight.Value)
			}
			if tt.expectedLeft == nil && b.Root.Left != nil {
				t.Errorf("Add() unexpected Left leaf")
			}
			if (tt.expectedLeft != nil && b.Root.Left != nil) && b.Root.Left.Value.Value != tt.expectedLeft.Value {
				t.Errorf("Add() Left value does not match expected, got %s, want %s", b.Root.Left.Value.Value, tt.expectedLeft.Value)
			}
		})
	}
}

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
			got, got1 := n.Find(tt.args.obj)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.Find() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Node.Find() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBST_Remove(t *testing.T) {
	type fields struct {
		Root *Node
	}
	type args struct {
		obj model.Object
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "value matches root",
			args: args{
				obj: model.Object{Value: "root"},
			},
			fields: fields{
				Root: &Node{Value: model.Object{"root"}},
			},
			want: true,
		},
		{
			name: "value doesn't match root",
			args: args{
				obj: model.Object{Value: "root"},
			},
			fields: fields{
				Root: &Node{Value: model.Object{"notRoot"}},
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BST{
				Root: tt.fields.Root,
			}
			got, err := b.Remove(tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("BST.Remove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BST.Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_Remove(t *testing.T) {
	rootVal := model.Object{Value: "root"}
	rightVal := model.Object{Value: "right"}
	leftVal := model.Object{Value: "le"}
	type fields struct {
		Value model.Object
		Left  *Node
		Right *Node
	}
	type args struct {
		parent *Node
		side   int
		obj    model.Object
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      bool
		wantChild string
	}{
		{
			name: "root is removed",
			fields: fields{
				Value: rootVal,
			},
			args: args{
				parent: &Node{Value: rootVal},
				side:   root,
				obj:    rootVal,
			},
			want: true,
		},
		{
			name: "right child is removed",
			fields: fields{
				Value: rootVal,
				Right: &Node{Value: rightVal},
			},
			args: args{
				parent: &Node{Value: rightVal},
				side:   root,
				obj:    rightVal,
			},
			want: true,
		},
		{
			name: "right child is removed and tree is correctly re-built",
			fields: fields{
				Value: rootVal,
				Right: &Node{
					Value: rightVal,
					Left:  &Node{Value: model.Object{Value: "righ"}},
					Right: &Node{Value: model.Object{Value: "righter"}},
				},
			},
			args: args{
				parent: &Node{Value: rightVal},
				side:   root,
				obj:    rightVal,
			},
			want: true,
		},
		{
			name: "right child is removed and right child is promoted",
			fields: fields{
				Value: rootVal,
				Right: &Node{
					Value: rightVal,
					Right: &Node{Value: model.Object{Value: "righter"}},
				},
			},
			args: args{
				parent: &Node{Value: rightVal},
				side:   root,
				obj:    rightVal,
			},
			want:      true,
			wantChild: "righter",
		},
		{
			name: "right child is removed and left child is promoted",
			fields: fields{
				Value: rootVal,
				Right: &Node{
					Value: rightVal,
					Left:  &Node{Value: model.Object{Value: "righ"}},
				},
			},
			args: args{
				parent: &Node{Value: rightVal},
				side:   root,
				obj:    rightVal,
			},
			want:      true,
			wantChild: "righ",
		},
		{
			name: "left child is removed",
			fields: fields{
				Value: rootVal,
				Left:  &Node{Value: leftVal},
			},
			args: args{
				parent: &Node{Value: leftVal},
				side:   root,
				obj:    leftVal,
			},
			want: true,
		},
		{
			name: "left child is removed and tree is correctly re-built",
			fields: fields{
				Value: rootVal,
				Left: &Node{
					Value: leftVal,
					Left:  &Node{Value: model.Object{Value: "l"}},
					Right: &Node{Value: model.Object{Value: "lef"}},
				},
			},
			args: args{
				parent: &Node{Value: leftVal},
				side:   root,
				obj:    leftVal,
			},
			want: true,
		},
		{
			name: "left child is removed and right child is promoted",
			fields: fields{
				Value: rootVal,
				Left: &Node{
					Value: leftVal,
					Right: &Node{Value: model.Object{Value: "lef"}},
				},
			},
			args: args{
				parent: &Node{Value: leftVal},
				side:   root,
				obj:    leftVal,
			},
			want:      true,
			wantChild: "lef",
		},
		{
			name: "left child is removed and left child is promoted",
			fields: fields{
				Value: rootVal,
				Left: &Node{
					Value: leftVal,
					Left:  &Node{Value: model.Object{Value: "l"}},
				},
			},
			args: args{
				parent: &Node{Value: leftVal},
				side:   root,
				obj:    leftVal,
			},
			want:      true,
			wantChild: "l",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Node{
				Value: tt.fields.Value,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := n.Remove(tt.args.parent, tt.args.side, tt.args.obj); got != tt.want {
				t.Errorf("Node.Remove() = %v, want %v", got, tt.want)
			}
			_, found := n.Find(tt.args.obj)
			if tt.want == true && found {
				t.Errorf("Node.Remove() value still found in tree after Remove()")
			}
			if tt.wantChild != "" {
				_, foundChild := n.Find(model.Object{Value: tt.wantChild})
				if !foundChild {
					t.Errorf("Node.Remove() expected child not found in tree after Remove()")
				}
			}
		})
	}
}
