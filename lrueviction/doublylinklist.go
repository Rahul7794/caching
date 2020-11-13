package lrueviction

type Element struct {
	Key        interface{}
	Value      interface{}
	Prev, Next *Element
}

type List struct {
	Head, Tail *Element
}

func (l *List) Add(node *Element) {
	headNext := l.Head.Next
	node.Next = headNext
	headNext.Prev = node
	l.Head.Next = node
	node.Prev = l.Head
}

func (l *List) Remove(node *Element) {
	next := node.Next
	prev := node.Prev
	next.Prev = prev
	prev.Next = next
}

func NewList() *List {
	list := &List{
		Head: &Element{},
		Tail: &Element{},
	}
	list.Head.Next = list.Tail
	list.Tail.Prev = list.Head
	return list
}
