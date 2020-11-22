package cache

import (
	"reflect"
	"testing"

	"caching/lrueviction"
)

func TestLRU(t *testing.T) {
	tests := []struct{
		name string
		setup func() *LRU
	} {
		{
			name: "LRU implementations",
			setup: func() *LRU {
				lru := LRU{
					items:    make(map[interface{}]*lrueviction.Element),
					capacity: 2,
					list:     lrueviction.NewList(),
				}
				lru.Set("1", "10")
				lru.Set("5", "12")
				return &lru
			},
		},
	}
	for _, tt := range tests {
		lru := tt.setup()
		fetch5 := lru.Get("5")
		if !reflect.DeepEqual(fetch5.(string), "12") {
			t.Errorf("Get()=%v, wanted %v", fetch5.(string), "12")
		}
		fetch1 := lru.Get("1")
		if !reflect.DeepEqual(fetch1.(string), "10") {
			t.Errorf("Get()=%v, wanted %v", fetch1.(string), "10")
		}
		fetch10 := lru.Get("10")
		if !reflect.DeepEqual(fetch10, nil) {
			t.Errorf("Get()=%v, wanted %v", fetch10, nil)
		}
		lru.Set("6", "14") // this pushed out key = 5 as LRU is full
		fetch5Again := lru.Get("5")
		if !reflect.DeepEqual(fetch5Again, nil) {
			t.Errorf("Get()=%v, wanted %v", fetch5Again, nil)
		}
	}
}
