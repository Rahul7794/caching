package cache

import (
	"container/list"
)

type LFU struct {
	items         map[interface{}]*Item // items stores values by key
	frequencyList *list.List            // frequencyList maintains set of nodes that have same access frequency
	capacity      int                   // capacity of cache
}

type Item struct {
	key             interface{}   // key of item
	value           interface{}   // value of item
	frequencyParent *list.Element // frequencyParent is a pointer to parent node(frequency)
}

type FrequencyItem struct {
	entries   map[*Item]byte // set of entries
	frequency int            // frequency is no of times an item is accessed.
}

// Get retrieve value provided the key
func (l *LFU) Get(key interface{}) interface{} {
	var result interface{}
	if item, found := l.items[key]; found {
		l.add(item)
		result = item.value
	}
	return result
}

// Set add value to cache to key
func (l *LFU) Set(key, value interface{}) {
	if item, found := l.items[key]; found {
		item.value = value
		l.add(item)
	} else {
		// if capacity limit reached, the least frequently used item is evicted, if there is more than one item as same
		// frequency access count then removes the most oldest item accessed
		if len(l.items) == l.capacity {
			l.evict()
		}
		newElement := new(Item)
		newElement.key = key
		newElement.value = value
		l.items[key] = newElement
		l.add(newElement)
	}
}

// add item to the frequency list
func (l *LFU) add(item *Item) {
	currentFrequencyNode := item.frequencyParent
	var nextFrequency int
	var nextFrequencyNode *list.Element

	// Get the nextFrequencyNode and NextFrequency of Item
	if currentFrequencyNode == nil {
		nextFrequency = 1
		nextFrequencyNode = l.frequencyList.Front()
	} else {
		nextFrequency = currentFrequencyNode.Value.(*FrequencyItem).frequency + 1
		nextFrequencyNode = currentFrequencyNode.Next()
	}
	// Once the item is accessed again, increment the frequency by 1 and check if next node in the frequency list exist,
	// if dont exist, or next node is nil, create a new node in frequency list(if list is empty push to front or push
	// after the current node (to maintain the order of frequency)
	if nextFrequencyNode == nil || nextFrequencyNode.Value.(*FrequencyItem).frequency != nextFrequency {
		newFrequencyItem := new(FrequencyItem)
		newFrequencyItem.frequency = nextFrequency
		newFrequencyItem.entries = make(map[*Item]byte)
		if currentFrequencyNode == nil {
			nextFrequencyNode = l.frequencyList.PushFront(newFrequencyItem)
		} else {
			nextFrequencyNode = l.frequencyList.InsertAfter(newFrequencyItem, currentFrequencyNode)
		}
	}
	// if exist change the parent pointer to the next node
	item.frequencyParent = nextFrequencyNode
	nextFrequencyNode.Value.(*FrequencyItem).entries[item] = 1
	// After the attaching the item to new node, remove the item from old node and if the node becomes empty after the
	// removal, then delete the node itself.
	if currentFrequencyNode != nil {
		l.delete(currentFrequencyNode, item)
	}
}

// delete item from frequency node, if node is empty after deletion, delete frequency node too.
func (l *LFU) delete(list *list.Element, item *Item) {
	frequencyItem := list.Value.(*FrequencyItem)
	delete(frequencyItem.entries, item)
	if len(frequencyItem.entries) == 0 {
		l.frequencyList.Remove(list)
	}
}

// evict the least frequently item used
func (l *LFU) evict() {
	if item := l.frequencyList.Front(); item != nil {
		for entry, _ := range item.Value.(*FrequencyItem).entries {
			delete(l.items, entry.key)
			l.delete(item, entry)
			break
		}
	}
}

// Flush clears the entire cache
func (l *LFU) Flush() {
	l.items = nil
	l.frequencyList = nil
}

// NewLFU return new object of LFU cache
func NewLFU(capacity int) Operations {
	storage := make(map[interface{}]*Item)
	return &LFU{
		items:         storage,
		capacity:      capacity,
		frequencyList: list.New(),
	}
}
