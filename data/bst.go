package data

import (
	"errors"
)

var (
	errNotFound       = errors.New("not found")
	errAlreadyExisted = errors.New("element already existed")
)

type node struct {
	Value uint32
	Left  *node
	Right *node
	Attch string
}

func (n node) Empty() bool {
	return n.Right == nil && n.Left == nil
}

type bsTree struct {
	root *node
}

func (bst *bsTree) Insert(v uint32, attch string) error {
	n := &node{Value: v, Attch: attch}
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
func (bst *bsTree) Search(v uint32) (s string, err error) {
	p := bst.root
	if p == nil {
		return "", errNotFound
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

func (bst *bsTree) Traverse(callback func(uint32, string)) {
	traverse(bst.root, func(i *node) {
		callback(i.Value, i.Attch)
	})
}
func (bst *bsTree) Root() *node {
	return bst.root
}
