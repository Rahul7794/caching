# Caching

Supports various types of mem caching evictions algorithm. Current support

- LRU (Least recently used) [Uses doubly link list and map makes it ]
- FIFO/LILO (First in First out/Last in Last out)

####LRU

- Used Doubly LinkList and Map to get, set the data in O(1).

####FIFO

- Used Queue to perform all operations in O(1).

To be implemented next:

- LFU (Least Frequently used)

#### Steps to run

The code expects two args:

- Capacity - Maximum capacity of cache.
- CachingEvictionAlgo - (fifo/lru, lfu).

**NOTE** : No inbuilt libraries used, self implementation of algo 