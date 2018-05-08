package data

import "errors"

var (
	errNotFound = errors.New("not found")
)

type node struct {
	Value uint32
	Left  *node
	Right *node
	Attch string
}

type BSTree struct {
	root *node
}

func (bst *BSTree) Insert(v uint32, attch string) {
	n := &node{Value: v, Attch: attch}
	if bst.root == nil {
		bst.root = n
		return
	}
	p, fa := bst.root, bst.root
	for p != nil {
		if p.Value == n.Value {
			return
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
}
func (bst *BSTree) Search(v uint32) (s string, err error) {
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
