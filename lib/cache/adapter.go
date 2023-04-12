package cache

import (
	"time"
	"github.com/patrickmn/go-cache"
)

var caching = cache.New(24*time.Hour, 24*time.Hour)

func GetCache(cacheKey string) (i interface{}) {
	if val, found := caching.Get(cacheKey); found {
		return val
	}
	return nil
}

func SetCache(cacheKey string, value interface{}, expireTime time.Duration) {
	if expireTime == 0 {
		expireTime = cache.DefaultExpiration
	}

	caching.Set(cacheKey, value, expireTime)
}
