package cache

import (
	"time"

	"github.com/keshareenv/go-base/data/cache/inmem"
)

type Cache interface {
	// Set adds item to cache replacing existing one
	Set(k string, val interface{})

	// Add adds item to cache only if the item doesn't exist or
	// the key has expired. It won't remove an active existing value
	Add(k string, val interface{}) error

	// Replace an item if it exists
	Replace(k string, val interface{}) error

	// SetWithDuration sets the key with a value for a time period
	SetWithDuration(k string, val interface{}, expiration time.Duration)

	// Get returns the value for the key from the cache and sets found flag as
	// true or it returns false if the value is not found
	Get(k string) (val interface{}, found bool)

	// Delete deletes the key from the cache, and doesn't do anything
	// if key is not found
	Delete(k string)
}

func NewInMemoryCache(
	expiry time.Duration,
	eviction time.Duration,
	options ...inmem.Option,
) (Cache, error) {
	return inmem.New(expiry, eviction, options...), nil
}
