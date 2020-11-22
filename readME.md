# Caching

Supports various types of mem caching evictions algorithm. Currently, support

- LRU (Least recently used)
- LFU (Least frequently used)
- FIFO/LILO (First in First out/Last in Last out)

#### LRU

- Used Doubly LinkList and Map to get, set the data in O(1).

#### LFU

- Used 2 LinkList and Map to get, set the data in O(1).

#### FIFO

- Used Queue to perform all operations in O(1).


#### Steps to run

The code expects two args:

- Capacity - Maximum capacity of cache.
- CachingEvictionAlgo - (fifo/lilo, lfu. lru).
