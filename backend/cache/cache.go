package cache

import (
	"container/list"
	"sync"
	"time"
)

type CacheEntry struct {
	Key       string
	Value     interface{}
	ExpiresAt time.Time
}

type LRUCache struct {
	maxKeys  int
	cache    map[string]*list.Element
	eviction *list.List
	mutex    sync.Mutex
}

func NewLRUCache(maxKeys int) *LRUCache {
	return &LRUCache{
		maxKeys:  maxKeys,
		cache:    make(map[string]*list.Element),
		eviction: list.New(),
	}
}

func (c *LRUCache) GetAll() []*CacheEntry {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	cacheEntries := make([]*CacheEntry, 0, c.eviction.Len())
	for elem := c.eviction.Front(); elem != nil; elem = elem.Next() {
		if time.Now().After(elem.Value.(*CacheEntry).ExpiresAt) {
			c.evictEntry(elem)
		}
		cacheEntries = append(cacheEntries, elem.Value.(*CacheEntry))
	}

	return cacheEntries
}

func (c *LRUCache) Set(key string, value interface{}, expiration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if elem, ok := c.cache[key]; ok {
		elem.Value.(*CacheEntry).Value = value
		elem.Value.(*CacheEntry).ExpiresAt = time.Now().Add(expiration)
		c.eviction.MoveToFront(elem)
	} else {
		entry := &CacheEntry{Key: key, Value: value, ExpiresAt: time.Now().Add(expiration)}
		elem := c.eviction.PushFront(entry)
		c.cache[key] = elem

		if c.eviction.Len() > c.maxKeys {
			c.evictLRU()
		}
	}
}

func (c *LRUCache) Get(key string) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if elem, ok := c.cache[key]; ok {
		if time.Now().After(elem.Value.(*CacheEntry).ExpiresAt) {
			// Evict expired entry
			c.evictEntry(elem)
			return nil, false
		}

		// Move the element to front (most recently used)
		c.eviction.MoveToFront(elem)
		return elem.Value.(*CacheEntry).Value, true
	}

	return nil, false
}

// evictLRU evicts the least recently used key from the cache
func (c *LRUCache) evictLRU() {
	elem := c.eviction.Back()
	if elem != nil {
		c.evictEntry(elem)
	}
}

// evictEntry removes the given entry from the cache
func (c *LRUCache) evictEntry(elem *list.Element) {
	delete(c.cache, elem.Value.(*CacheEntry).Key)
	c.eviction.Remove(elem)
}
