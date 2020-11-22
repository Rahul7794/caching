package lrueviction

type Element struct {
	Key        interface{} // Key of a Node
	Value      interface{} // Value of a Node
	Prev, Next *Element    // Next and Prev pointer of a node
}

type List struct {
	Head, Tail *Element    // Head and Tail pointer of a doubly link list
}

// Add element to doubly link list
func (l *List) Add(node *Element) {
	headNext := l.Head.Next
	node.Next = headNext
	headNext.Prev = node
	l.Head.Next = node
	node.Prev = l.Head
}

// Remove element from doubly link list
func (l *List) Remove(node *Element) {
	next := node.Next
	prev := node.Prev
	next.Prev = prev
	prev.Next = next
}

// NewList return new List
func NewList() *List {
	list := &List{
		Head: &Element{},
		Tail: &Element{},
	}
	list.Head.Next = list.Tail
	list.Tail.Prev = list.Head
	return list
}
