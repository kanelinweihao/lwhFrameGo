package cacheWriter

import (
	"context"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/goroutine"
	"github.com/redis/go-redis/v9"
	"time"
)

type EntityCacheDataWrite struct {
	CacheRedis *redis.Client
	CTX        context.Context
	CacheKey   string
	CacheValue string
	TTL        time.Duration
}

func (self *EntityCacheDataWrite) WriteToChannelOfWriteCacheData(entityChannel *goroutine.EntityChannel) {
	self.setCache()
	cacheKey := self.CacheKey
	entityChannel.WriteOnce(cacheKey)
	return
}

func (self *EntityCacheDataWrite) setCache() {
	r := self.CacheRedis
	ctx := self.CTX
	cacheKey := self.CacheKey
	cacheValue := self.CacheValue
	ttl := self.TTL
	errRedisSet := r.Set(ctx, cacheKey, cacheValue, ttl).Err()
	err.ErrCheck(errRedisSet)
	return
}
