package fifoeviction

import (
	"fmt"
	"sync"
)

// Queue is implementation of Queue DataStructure in Go.
type Queue struct {
	items []interface{} // Elements are stored in items in FIFO order.
	mutex *sync.Mutex   // mutex is introduced for thread safety.
}

// Enqueue adds an element to the Queue.
func (queue *Queue) Enqueue(item interface{}) {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	queue.items = append(queue.items, item)
}

// Dequeue remove and return the first element in Queue.
func (queue *Queue) Dequeue() (interface{}, error) {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	// handle the case where Queue is empty
	if len(queue.items) == 0 {
		return nil, fmt.Errorf("the queue is empty")
	}

	lastItem := queue.items[0]
	queue.items = queue.items[1:]

	return lastItem, nil
}

// Size returns the size of Queue
func (queue *Queue) Size() int {
	return len(queue.items)
}

// Reset flushes the Queue.
func (queue *Queue) Reset() {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	queue.items = nil
}

// NewQueue return Queue object
func NewQueue() *Queue {
	queue := &Queue{}
	queue.items = make([]interface{}, 0)
	queue.mutex = new(sync.Mutex)
	return queue
}
