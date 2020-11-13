package lrueviction

import (
	"reflect"
	"testing"
)

func TestList_Push(t *testing.T) {
	tests := []struct {
		name  string
		setup func() *List
	}{
		{
			name: "test push function",
			setup: func() *List {
				list := NewList()
				list.Add(&Element{
					Key:   1,
					Value: 5,
				})
				list.Add(&Element{
					Key:   2,
					Value: 3,
				})
				list.Add(&Element{
					Key:   7,
					Value: 4,
				})
				return list
			},
		},
	}
	for _, tt := range tests {
		list := tt.setup()
		headNextV := list.Head.Next.Value
		tailPrevV := list.Tail.Prev.Value
		if !reflect.DeepEqual(headNextV, 4) {
			t.Errorf("Add()=%v, wanted %v", headNextV, 4)
		}
		if !reflect.DeepEqual(tailPrevV, 5) {
			t.Errorf("Add()=%v, wanted %v", tailPrevV, 5)
		}
	}
}
