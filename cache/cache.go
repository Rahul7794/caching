package cache

// EvictionPolicy stores key as cache mechanism and its relevant object as value
var EvictionPolicy = map[string]func(capacity int) Operations{
	"lru":  NewLRU,
	"lfu":  NewLFU,
	"fifo": NewFIFO,
}

// NewCache initializes cache mechanism object
func NewCache(evictionMechanism string, capacity int) Operations {
	if _, ok := EvictionPolicy[evictionMechanism]; ok {
		return EvictionPolicy[evictionMechanism](capacity)
	}
	return nil
}
