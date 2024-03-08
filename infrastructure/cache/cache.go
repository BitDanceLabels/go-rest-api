// /infrastructure/cache/cache.go
package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var CacheClient *cache.Cache

func InitCache() {
	CacheClient = cache.New(5*time.Minute, 10*time.Minute)
}
