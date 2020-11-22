package cache

import (
	"caching/fifoeviction"
)


type FIFO struct {
	items    map[interface{}]interface{}
	capacity int
	queue    *fifoeviction.Queue
}

// Get retrieve value provided the key
func (f *FIFO) Get(key interface{}) interface{} {
	var result interface{}
	if value, found := f.items[key]; found {
		result = value
	}
	return result
}

// Set add value to cache to key
func (f *FIFO) Set(key, value interface{}) {
	if _, found := f.items[key]; found {
		f.items[key] = value
		return
	} else {
		if len(f.items) == f.capacity {
			e, _ := f.queue.Dequeue()
			delete(f.items, e)
		}
		f.items[key] = value
		f.queue.Enqueue(key)
	}
}

// Flush clears the entire cache
func (f *FIFO) Flush() {
	f.queue.Reset()
	f.items = nil
}

// NewFIFO initializes FIFO cache
func NewFIFO(capacity int) Operations {
	storage := make(map[interface{}]interface{}, capacity)
	return &FIFO{
		items:    storage,
		capacity: capacity,
		queue:    fifoeviction.NewQueue(),
	}
}
