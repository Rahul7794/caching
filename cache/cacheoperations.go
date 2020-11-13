package cache

// Operations to be performed on cache
type Operations interface {
	Get(key interface{}) interface{} // Get value(always +ve) of the key, if the key exists in the cache else returns -1
	Set(key, value interface{})      // Set or insert the value if the key is not already present. When the cache reaches its capacity.
	Flush()                          // Flush removes all the element in cache
}
