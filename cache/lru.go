package cache

import (
	"caching/lrueviction"
)

type LRU struct {
	items    map[interface{}]*lrueviction.Element
	capacity int
	list     *lrueviction.List
}

func NewLRU(capacity int) Operations {
	storage := make(map[interface{}]*lrueviction.Element, capacity)
	return &LRU{
		items:    storage,
		capacity: capacity,
		list:     lrueviction.NewList(),
	}
}

func (l *LRU) Get(key interface{}) interface{} {
	var result interface{}
	if node, found := l.items[key]; found {
		l.list.Remove(node)
		l.list.Add(node)
		result = node.Value
	}
	return result
}

func (l *LRU) Set(key, value interface{}) {
	if node, found := l.items[key]; found {
		l.list.Remove(node)
		node.Value = value
		l.list.Add(node)
	} else {
		if len(l.items) == l.capacity {
			delete(l.items, l.list.Tail.Prev.Key)
			l.list.Remove(l.list.Tail.Prev)
		}
		newElement := &lrueviction.Element{}
		newElement.Key = key
		newElement.Value = value
		l.items[key] = newElement
		l.list.Add(newElement)
	}
}

func (l *LRU) Flush() {
	l.items = nil
	l.list = nil
}
