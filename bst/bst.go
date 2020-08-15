package bst

import (
	"errors"
	"go-datastructures/model"
)

const (
	root  = 0
	left  = 1
	right = 2
)

// BST :: struct :: Gives Root a home, and a place
// for calling the various iteration methods from.
type BST struct {
	Root *Node
}

// Add :: func :: This will add in strings, with the left/right
// placement of new nodes being decided by the length of the
// incoming string. Nice and arbitrary.
func (b *BST) Add(obj model.Object) {
	if b.Root != nil {
		b.Root.Add(obj)
		return
	} else {
		b.Root = &Node{
			Value: obj,
		}
	}
}

func (b BST) Remove(obj model.Object) (bool, error) {
	removed := b.Root.Remove(b.Root, root, obj)
	if !removed {
		return removed, errors.New("object not found in list")
	}
	return removed, nil
}

func (b BST) Find(obj model.Object) (*Node, bool) {
	return b.Root.Find(b.Root, obj)
}

// NodeFunc :: func :: Some function that takes in  model.Object
// and does an operation on the stored value, with no return.
type NodeFunc func(obj model.Object)

// PreOrder :: func :: Processes current node, then left, the right
func (b BST) PreOrder(f NodeFunc) {
	if b.Root == nil {
		return
	}
	f(b.Root.Value)
	b.Root.GoLeft(f)
	b.Root.GoRight(f)
}

// InOrder :: func :: Processes left, current node, the right
func (b BST) InOrder(f NodeFunc) {
	if b.Root == nil {
		return
	}
	b.Root.GoLeft(f)
	f(b.Root.Value)
	b.Root.GoRight(f)
}

// PostOrder :: func :: Processes left, the right, current node
func (b BST) PostOrder(f NodeFunc) {
	if b.Root == nil {
		return
	}
	b.Root.GoLeft(f)
	b.Root.GoRight(f)
	f(b.Root.Value)
}

type Node struct {
	Value model.Object
	Left  *Node
	Right *Node
}

func (n Node) Find(parent *Node, obj model.Object) (*Node, bool) {
	match := n.Value.Value == obj.Value
	if match {
		return &n, true
	}
	if less(obj, n.Value) && n.Left != nil {
		return n.Left.Find(&n, obj)
	} else if n.Right != nil {
		return n.Right.Find(&n, obj)
	}
	return nil, false
}

// Add :: func :: Adds a new node
func (n *Node) Add(obj model.Object) {
	if n.Value.Value == "" {
		n.Value = obj
		return
	}
	if less(n.Value, obj) {
		if n.Right == nil {
			n.Right = &Node{Value: obj}
			return
		}
		n.Right.Add(obj)
		return
	}
	if n.Left == nil {
		n.Left = &Node{Value: obj}
		return
	}
	n.Left.Add(obj)
}

// Remove :: func :: Removes any matching nodes
func (n Node) Remove(parent *Node, side int, obj model.Object) bool {
	switch side {
	case root:
		match := parent.Value.Value == n.Value.Value
		if !match {
			return false
		}
		n.Value.Value = ""
	case left:
		if match := n.Value.Value == obj.Value; !match {
			return false
		}
	case right:
		if match := n.Value.Value == obj.Value; !match {
			return false
		}
	default:
	}
	return true
}

// Methods used for the various processing Methods on BST
func (n *Node) GoLeft(f NodeFunc) *Node {
	if n.Left != nil {
		f(n.Left.Value)
		return n.Left
	}
	return n
}

func (n *Node) GoRight(f NodeFunc) *Node {
	if n.Right != nil {
		f(n.Right.Value)
		return n.Right
	}
	return n
}

func less(o1, o2 model.Object) bool {
	return len(o1.Value) < len(o2.Value)
}
