package avl

import (
	"errors"
	"go-datastructures/model"
)

const (
	root = iota
	left
	right
)

// AVL :: struct :: Self-balancing BST.
type AVL struct {
	Root    *Node
	Height  int
	LHeight int
	RHeight int
}

// Add :: func :: This will add in strings, with the left/right
// placement of new nodes being decided by the length of the
// incoming string. Nice and arbitrary.
func (a *AVL) Add(obj model.Object) {
	if a.Root != nil {
		switch a.Root.add(obj) {
		case root:
			a.Height++
		case left:
			a.Height++
			a.LHeight++
		case right:
			a.Height++
			a.RHeight++
		}
		return
	} else {
		a.Root = &Node{
			Value: obj,
		}
		a.Height++
	}
}

func (a AVL) Balance() {
	// TODO
}

func (a AVL) LeftRotate() {
	// TODO
}

func (a AVL) LeftRightRotate() {
	// TODO
}

func (a AVL) RightRotate() {
	// TODO
}

func (a AVL) RightLeftRotate() {
	// TODO
}

// Remove :: func :: Removes a object/value from the AVL. Returns an error if the value is not in the AVL.
func (a AVL) Remove(obj model.Object) (bool, error) {
	removed := a.Root.remove(a.Root, root, obj)
	if !removed {
		return removed, errors.New("object not found in list")
	}
	return removed, nil
}

func (a AVL) Find(obj model.Object) (*Node, bool) {
	return a.Root.find(obj)
}

// NodeFunc :: func :: Some function that takes in  model.Object
// and does an operation on the stored value, with no return.
type NodeFunc func(obj model.Object)

// PreOrder :: func :: Processes current, left, right
func (a AVL) PreOrder(f NodeFunc) {
	if a.Root == nil {
		return
	}
	a.Root.preOrder(f)
}

// InOrder :: func :: Processes left, current, right
// Items in the list will be processed in Sort Order
func (a AVL) InOrder(f NodeFunc) {
	if a.Root == nil {
		return
	}
	a.Root.inOrder(f)
}

// PostOrder :: func :: Processes left, right, current
// Root will be processed last -- Deletion of the entire tree could be a use case
func (a AVL) PostOrder(f NodeFunc) {
	if a.Root == nil {
		return
	}
	a.Root.postOrder(f)
}

// Node :: struct :: Node holds the values for the elements of the AVL, and any pointers to child values
type Node struct {
	Value  model.Object
	Parent *Node
	Left   *Node
	Right  *Node
}

func (n Node) preOrder(f NodeFunc) {
	f(n.Value)
	if n.Left != nil {
		n.Left.preOrder(f)
	}
	if n.Right != nil {
		n.Right.preOrder(f)
	}
}

func (n Node) inOrder(f NodeFunc) {
	if n.Left != nil {
		n.Left.inOrder(f)
	}
	f(n.Value)
	if n.Right != nil {
		n.Right.inOrder(f)
	}
}

func (n Node) postOrder(f NodeFunc) {
	if n.Left != nil {
		n.Left.postOrder(f)
	}
	if n.Right != nil {
		n.Right.postOrder(f)
	}
	f(n.Value)
}

func (n Node) find(obj model.Object) (*Node, bool) {
	match := n.Value.Value == obj.Value
	if match {
		return &n, true
	}
	if less(obj, n.Value) && n.Left != nil {
		return n.Left.find(obj)
	} else if n.Right != nil {
		return n.Right.find(obj)
	}
	return nil, false
}

// add :: func :: adds a new node
func (n *Node) add(obj model.Object) int {
	if n.Value.Value == "" {
		n.Value = obj
		return root
	}
	if less(n.Value, obj) {
		if n.Right == nil {
			n.Right = &Node{Value: obj}
		} else {
			n.Right.add(obj)
		}
		return right
	}
	if n.Left == nil {
		n.Left = &Node{Value: obj}
	} else {
		n.Left.add(obj)
	}
	return left
}

// remove :: func :: removes any matching nodes
func (n *Node) remove(parent *Node, side int, obj model.Object) bool {
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

func less(o1, o2 model.Object) bool {
	return len(o1.Value) < len(o2.Value)
}
