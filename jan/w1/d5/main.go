package d5

import "time"

type Cache struct {
	Data map[string]*CacheEntry
}

type CacheEntry struct {
	Value     any
	ExpiresAt time.Time
}

func (c Cache) Set(key string, value any) {
	c.Data[key] = &CacheEntry{
		Value:     value,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
}

// this is perfect because we are only reading from the cache. in go, a map is a pointer to a struct
// so ideally, you have access to the underlying map
func (c Cache) Get(key string) (any, bool) {
	value, exists := c.Data[key]

	if !exists || time.Now().After(value.ExpiresAt) {
		return nil, false
	}
	return value.Value, true
}

func (c Cache) Prune() {
	for k, v := range c.Data {
		if time.Now().After(v.ExpiresAt) {
			delete(c.Data, k)
		}
	}
}

func (c *Cache) SetAsPointer(key string, value any) {
	c.Data[key] = &CacheEntry{
		Value:     value,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
}

func (c Cache) GetAsPointer(key string) (any, bool) {
	value := c.Data[key]

	if value != nil {
		return value.Value, true
	}
	return nil, false
}

func (c Cache) PruneAsPointer() {
	for k, v := range c.Data {
		if time.Now().After(v.ExpiresAt) {
			delete(c.Data, k)
		}
	}
}
