package cacheSet

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/cache"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"time"
)

type EntityCacheSet struct {
	ParamsOut  base.AttrT1
	CacheKey   string
	CacheValue string
	TTL        time.Duration
}

func (self *EntityCacheSet) SetToRedis() {
	cacheKey := self.CacheKey
	cacheValue := self.CacheValue
	ttl := self.TTL
	entityCache := cache.MakeEntityOfCache()
	defer entityCache.CloseCache()
	entityCache.SetToCache(
		cacheKey,
		cacheValue,
		ttl)
	return
}

func (self *EntityCacheSet) GetFromRedis() (cacheValue string) {
	cacheKey := self.CacheKey
	entityCache := cache.MakeEntityOfCache()
	defer entityCache.CloseCache()
	cacheValue = entityCache.GetFromCache(cacheKey)
	return cacheValue
}
