package data

type element struct {
	Value *node
	Pre   *element
}

type nodeStack struct {
	tail *element
}

func newStack() *nodeStack {
	return new(nodeStack)
}
func (ns *nodeStack) Push(n *node) {
	e := &element{Value: n}
	if ns.tail == nil {
		ns.tail = e
	} else {
		e.Pre = ns.tail
		ns.tail = e
	}
}
func (ns *nodeStack) Pop() *node {
	if ns.tail == nil {
		return nil
	}
	e := ns.tail
	ns.tail = e.Pre
	e.Pre = nil
	return e.Value
}
func (ns *nodeStack) Empty() bool {
	return ns.tail == nil
}
