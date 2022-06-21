package bst

import (
	"errors"
	"go-datastructures/model"
)

const (
	root = iota
	left
	right
)

// BST :: struct :: Basic Binary Search Tree implementation.
type BST[T any] struct {
	Root *Node[T]
}

// Add :: func :: This will add in strings, with the left/right
// placement of new nodes being decided by the length of the
// incoming string. Nice and arbitrary.
func (b *BST[T]) Add(t T) {
	if b.Root != nil {
		b.Root.add(t)
		return
	} else {
		b.Root = &Node[T]{
			Value: t,
		}
	}
}

// Remove :: func :: Removes a object/value from the BST. Returns an error if the value is not in the BST
func (b BST[T]) Remove(obj T) (bool, error) {
	removed := b.Root.remove(b.Root, root, obj)
	if !removed {
		return removed, errors.New("object not found in list")
	}
	return removed, nil
}

func (b BST[T]) Find(obj T) (*Node[T], bool) {
	return b.Root.find(obj)
}

// NodeFunc :: func :: Some function that takes in  model.Object
// and does an operation on the stored value, with no return.
type NodeFunc func(t any)

// PreOrder :: func :: Processes current, left, right
func (b BST[T]) PreOrder(f NodeFunc) {
	if b.Root == nil {
		return
	}
	b.Root.preOrder(f)
}

// InOrder :: func :: Processes left, current, right
// Items in the list will be processed in Sort Order
func (b BST[T]) InOrder(f NodeFunc) {
	if b.Root == nil {
		return
	}
	b.Root.inOrder(f)
}

// PostOrder :: func :: Processes left, right, current
// Root will be processed last -- Deletion of the entire tree could be a use case
func (b BST[T]) PostOrder(f NodeFunc) {
	if b.Root == nil {
		return
	}
	b.Root.postOrder(f)
}

// Node :: struct :: Node holds the values for the elements of the BST, and any pointers to child values
type Node[T any] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

func (n Node[T]) preOrder(f NodeFunc) {
	f(n.Value)
	if n.Left != nil {
		n.Left.preOrder(f)
	}
	if n.Right != nil {
		n.Right.preOrder(f)
	}
}

func (n Node[T]) inOrder(f NodeFunc) {
	if n.Left != nil {
		n.Left.inOrder(f)
	}
	f(n.Value)
	if n.Right != nil {
		n.Right.inOrder(f)
	}
}

func (n Node[T]) postOrder(f NodeFunc) {
	if n.Left != nil {
		n.Left.postOrder(f)
	}
	if n.Right != nil {
		n.Right.postOrder(f)
	}
	f(n.Value)
}

func (n Node[T]) find(t T) (*Node[T], bool) {
	match := n.Value.Value == t.Value
	if match {
		return &n, true
	}
	if less(t, n.Value) && n.Left != nil {
		return n.Left.find(t)
	} else if n.Right != nil {
		return n.Right.find(t)
	}
	return nil, false
}

// add :: func :: adds a new node
func (n *Node[T]) add(t T) {
	if n.Value == nil {
		n.Value = t
		return
	}
	if less(n.Value, t) {
		if n.Right == nil {
			n.Right = &Node[T]{Value: t}
			return
		}
		n.Right.add(t)
		return
	}
	if n.Left == nil {
		n.Left = &Node{Value: t}
		return
	}
	n.Left.add(t)
}

// remove :: func :: removes any matching nodes
func (n *Node[T]) remove(parent *Node[T], side int, obj T) bool {
	switch side {
	case root:
		match := parent.Value.Value == n.Value.Value
		if !match {
			less := less(n.Value, obj)
			if less {
				return n.Right != nil && n.Right.remove(n, right, obj)
			}
			return n.Left != nil && n.Left.remove(n, left, obj)
		}
		n.Value.Value = ""
		return true
	case left:
		if n.Value.Value == "" {
			return false
		}
		// If this is a match for the current node
		if match := n.Value.Value == obj.Value; match {
			if n.Left != nil {
				// Promote the left child
				parent.Left = n.Left
				n.Left.Right = n.Right
			} else {
				// Promote the right child
				parent.Left = n.Right
			}
			return true
		}
		// Not a match, so return removing any non-nil children
		return n.Right != nil && n.Right.remove(n, right, n.Right.Value) ||
			n.Left != nil && n.Left.remove(n, left, n.Left.Value)
	case right:
		if n.Value.Value == "" {
			return false
		}
		// If this is a match for the current node
		if match := n.Value.Value == obj.Value; match {
			if n.Right != nil {
				// Promote the right child
				parent.Right = n.Right
				n.Right.Left = n.Left
			} else {
				// Promote the left child
				parent.Right = n.Left
			}
			return true
		}
		// Not a match, so return removing any non-nil children
		return n.Right != nil && n.Right.remove(n, right, n.Right.Value) ||
			n.Left != nil && n.Left.remove(n, left, n.Left.Value)
	default:
	}
	return false
}

func less[T any](o1, o2 T) bool {
	// return len(o1.Value) < len(o2.Value)
    return false
}
