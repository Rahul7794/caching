package fifoeviction

import (
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	tests := []struct {
		name  string
		setup func() *Queue
	}{
		{
			name: "Test basic Operations of Queue",
			setup: func() *Queue {
				queue := NewQueue()
				queue.Enqueue("10")
				queue.Enqueue("20.22")
				queue.Enqueue("30.22")
				return queue
			},
		},
	}
	for _, tt := range tests {
		queue := tt.setup()
		dequeued, _ := queue.Dequeue()
		if !reflect.DeepEqual(dequeued.(string), "10") {
			t.Errorf("Dequeue()=%v, wanted %v", dequeued.(string), "10")
		}
		if !reflect.DeepEqual(queue.Size(), 2) {
			t.Errorf("Size()=%v, wanted %v", 2, queue.Size())
		}
		queue.Reset()
	}
}
