package commons

import "sync"

// Cache is a simple in-memory cache.
type Cache struct {
	data  map[string]interface{}
	mutex sync.RWMutex
}

// NewCache creates a new Cache.
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

// Set adds a key-value pair to the cache.
func (c *Cache) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = value
}

// Get retrieves a value from the cache by key.
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	value, found := c.data[key]
	return value, found
}

// Delete removes a key-value pair from the cache by key.
func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.data, key)
}

// Has checks if a key exists in the cache.
func (c *Cache) Has(key string) bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	_, found := c.data[key]
	return found
}
