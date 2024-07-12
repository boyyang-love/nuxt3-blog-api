package helper

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"time"
)

type Cache struct {
	BigCache *bigcache.BigCache
}

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) Init() (err error) {
	c.BigCache, err = bigcache.New(context.Background(), bigcache.DefaultConfig(5*time.Minute))
	return err
}

func (c *Cache) Get(key string) (interface{}, bool) {
	value, err := c.BigCache.Get(key)
	if err != nil {
		return nil, false
	}

	return value, true
}

func (c *Cache) Set(key string, b []byte) error {
	err := c.BigCache.Set(key, b)
	if err != nil {
		return err
	}

	return nil
}
