package cache

import (
	"time"
)

type CacheItem struct {
	expire time.Time
	data   []byte
}
