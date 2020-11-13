package cache

import (
	"reflect"
	"testing"

	"caching/fifoeviction"
)

func TestFIFO(t *testing.T) {
	tests := []struct{
		name string
		setup func() *FIFO
	} {
		{
			name:"FIFO implementations",
			setup: func() *FIFO {
				fifo := FIFO{
					items:    make(map[interface{}]interface{}),
					capacity: 4,
					queue:    fifoeviction.NewQueue(),
				}
				fifo.Set("A", "A")
				fifo.Set("B", "B")
				fifo.Set("C", "C")
				fifo.Set("D", "D")
				return &fifo
			},
		},
	}
	for _, tt := range tests {
		fifo := tt.setup()
		fetch := fifo.Get("D")
		if !reflect.DeepEqual(fetch.(string), "D") {
			t.Errorf("Dequeue()=%v, wanted %v", fetch.(string), "D")
		}
		fifo.Set("E", "E")
		fetchedFirst := fifo.Get("A")
		if !reflect.DeepEqual(fetchedFirst, nil) {
			t.Errorf("Dequeue()=%v, wanted %v", fetchedFirst, nil)
		}
		fetchLatest := fifo.Get("E")
		if !reflect.DeepEqual(fetchLatest, "E") {
			t.Errorf("Dequeue()=%v, wanted %v", fetchLatest, "E")
		}
	}
}
