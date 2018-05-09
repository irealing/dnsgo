package data

import (
	"errors"
)

var (
	errNotFound       = errors.New("not found")
	errAlreadyExisted = errors.New("element already existed")
)

type TraverMethod func(*Index)
type IndexTree interface {
	Insert(index *Index) error
	Search(v uint32) (*Index, error)
	Traverse(method TraverMethod)
	TraverseLeft(method TraverMethod)
	TraverseRight(method TraverMethod)
	Root() *Index
	Empty() bool
}

type node struct {
	Value uint32
	Left  *node
	Right *node
	Attch *Index
}

func (n node) Empty() bool {
	return n.Right == nil && n.Left == nil
}
func NewIndexTree() IndexTree {
	return new(bsTree)
}

type bsTree struct {
	root *node
}

func (bst *bsTree) Insert(ele *Index) error {
	n := &node{Value: ele.Index(), Attch: ele}
	if bst.root == nil {
		bst.root = n
		return errAlreadyExisted
	}
	p, fa := bst.root, bst.root
	for p != nil {
		if p.Value == n.Value {
			return errAlreadyExisted
		}
		fa = p
		if p.Value > n.Value {
			p = p.Left
		} else {
			p = p.Right
		}
	}
	if fa.Value > n.Value {
		fa.Left = n
	} else {
		fa.Right = n
	}
	return nil
}
func (bst *bsTree) Search(v uint32) (s *Index, err error) {
	p := bst.root
	if p == nil {
		return nil, errNotFound
	}
	for p != nil {
		if p.Value == v {
			break
		}
		if p.Value > v {
			p = p.Left
		} else {
			p = p.Right
		}
	}
	if p != nil {
		s = p.Attch
	} else {
		err = errNotFound
	}
	return
}
func traverse(n *node, callback func(*node)) {
	p := n
	ns := newStack()
	for p != nil || !ns.Empty() {
		for p != nil {
			ns.Push(p)
			p = p.Left
		}
		if !ns.Empty() {
			p = ns.Pop()
			callback(p)
			p = p.Right
		}
	}
}

func (bst *bsTree) Traverse(method TraverMethod) {
	if bst.root == nil {
		return
	}
	traverse(bst.root, func(i *node) {
		method(i.Attch)
	})
}
func (bst *bsTree) TraverseLeft(method TraverMethod) {
	if bst.root == nil || bst.root.Left == nil {
		return
	}
	traverse(bst.root.Left, func(i *node) {
		method(i.Attch)
	})
}
func (bst *bsTree) TraverseRight(method TraverMethod) {
	if bst.root == nil || bst.root.Right == nil {
		return
	}
	traverse(bst.root.Right, func(i *node) {
		method(i.Attch)
	})
}
func (bst *bsTree) Root() *Index {
	if bst.root != nil {
		return bst.root.Attch
	}
	return nil
}
func (bst *bsTree) Empty() bool {
	return bst.root == nil
}
