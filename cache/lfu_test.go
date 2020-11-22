package cache

import (
	"container/list"
	"reflect"
	"testing"
)

func TestLFU(t *testing.T) {
	tests := []struct {
		name  string
		setup func() *LFU
	}{
		{
			name: "LRU implementations",
			setup: func() *LFU {
				lfu := LFU{
					items:         make(map[interface{}]*Item),
					capacity:      2,
					frequencyList: list.New(),
				}
				lfu.Set("1", "5")
				lfu.Set("5", "12")
				return &lfu
			},
		},
	}
	for _, tt := range tests {
		lfu := tt.setup()
		fetch5 := lfu.Get("5")
		if !reflect.DeepEqual(fetch5.(string), "12") {
			t.Errorf("Get()=%v, wanted %v", fetch5.(string), "12")
		}
		fetch1 := lfu.Get("1")
		if !reflect.DeepEqual(fetch1.(string), "5") {
			t.Errorf("Get()=%v, wanted %v", fetch1.(string), "5")
		}
		fetch10 := lfu.Get("10")
		if !reflect.DeepEqual(fetch10, nil) {
			t.Errorf("Get()=%v, wanted %v", fetch10, nil)
		}
		lfu.Set("6", "14") // this pushed out key = 5 as Cache is full
		fetch5Again := lfu.Get("5")
		if !reflect.DeepEqual(fetch5Again, nil) {
			t.Errorf("Get()=%v, wanted %v", fetch5Again, nil)
		}
	}
}
