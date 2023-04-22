package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cacheKey := "mykey"
	cacheValue := "myvalue"

	SetCache(cacheKey, cacheValue, 2*time.Second)

	value := GetCache(cacheKey)

	assert.Equal(t, cacheValue, value)

	time.Sleep(3 * time.Second)

	expiredValue := GetCache(cacheKey)

	assert.Nil(t, expiredValue)
}
