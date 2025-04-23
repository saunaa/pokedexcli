package pokecache

import ( "time"
		"sync"
)

type cacheEntry struct {
	createdAt 		time.Time
	val 			[]byte
}

type Cache struct {
	entries			map[string]cacheEntry
	mux 			sync.RWMutex


}


func NewCache(interval time.Duration) *Cache{
		cache := &Cache{
		entries: make(map[string]cacheEntry),
		}
		go cache.reapLoop(interval)
		return cache
	}
	


func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	
	entry := cacheEntry{
		createdAt:	time.Now(),
		val:		value,
	}

	c.entries[key] = entry
	


}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()

	entry, ok := c.entries[key] 
	if !ok {
		return nil, false
		}
	return entry.val, true

	

}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mux.Lock()
		for key, value := range c.entries {
			if time.Since(value.createdAt) > interval {
				delete(c.entries, key)
				}
		} 
		c.mux.Unlock()

	}
}
