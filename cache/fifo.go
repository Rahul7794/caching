package cache

import (
	"caching/fifoeviction"
)

type FIFO struct {
	items    map[interface{}]interface{}
	capacity int
	queue    *fifoeviction.Queue
}

func NewFIFO(capacity int) Operations {
	storage := make(map[interface{}]interface{}, capacity)
	return &FIFO{
		items:    storage,
		capacity: capacity,
		queue:    fifoeviction.NewQueue(),
	}
}

func (f *FIFO) Get(key interface{}) interface{} {
	var result interface{}
	if value, found := f.items[key]; found {
		result = value
	}
	return result
}

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

func (f *FIFO) Flush() {
	f.queue.Reset()
	f.items = nil
}
