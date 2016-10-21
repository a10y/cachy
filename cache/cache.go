// Implements super dumb LRU cache ~30 LOC
package cache

import (
	"container/list"
	"fmt"
	"strings"
)

// Type for a cache
type cacheEntry struct {
	key   string
	value interface{} // value can be anything
}

// A linked list of cache entries and i mostly want to see how long this line can be before it tells me to move it to the next line, or just does it automatically
type Cache struct {
	entries *list.List // List of cache entries tracked locally, sorted by key
}

func New() *Cache {
	return &Cache{list.New()}
}

// Adds the given mapping into the cache, no expiration
func (c *Cache) Add(key string, value interface{}) {
	c.entries.PushFront(cacheEntry{key, value})
}

// Attempts to lookup the given key, returns error if not found
func (c *Cache) Lookup(key string) (interface{}, error) {
	// entries
	entriesPtr := c.entries.Front()
	for ; entriesPtr != nil; entriesPtr = entriesPtr.Next() {
		next := entriesPtr.Value.(cacheEntry)
		if strings.EqualFold(next.key, key) {
			// return this one
			// Move to the front of the list
			c.entries.MoveToFront(entriesPtr)
			return next.value, nil
		}
	}
	return nil, fmt.Errorf("key \"%v\" not found", key)
}

// Evicts the least-recently accessed item in the cache
func (c *Cache) Evict() {
	if back := c.entries.Back(); back != nil {
		c.entries.Remove(back)
	}
}
