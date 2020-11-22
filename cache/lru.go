package cache

import (
	"caching/lrueviction"
)

type LRU struct {
	items    map[interface{}]*lrueviction.Element // items stores values by key
	capacity int                                  // capacity of cache
	list     *lrueviction.List                    // list maintains more recently used in the end of DLL and least recently used in front of DLL
}

// Get retrieve value provided the key
func (l *LRU) Get(key interface{}) interface{} {
	var result interface{}
	if node, found := l.items[key]; found {
		l.list.Remove(node)
		l.list.Add(node)
		result = node.Value
	}
	return result
}

// Set add value to cache to key
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

// Flush clears the entire cache
func (l *LRU) Flush() {
	l.items = nil
	l.list = nil
}

// NewLRU initializes LRU cache
func NewLRU(capacity int) Operations {
	storage := make(map[interface{}]*lrueviction.Element)
	return &LRU{
		items:    storage,
		capacity: capacity,
		list:     lrueviction.NewList(),
	}
}
