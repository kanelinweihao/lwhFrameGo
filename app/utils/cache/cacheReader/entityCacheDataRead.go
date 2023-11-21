package cacheReader

import (
	"context"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/goroutine"
	"github.com/redis/go-redis/v9"
	"time"
)

type EntityCacheDataRead struct {
	CacheRedis *redis.Client
	CTX        context.Context
	CacheKey   string
	CacheValue string
	TTL        time.Duration
}

func (self *EntityCacheDataRead) WriteToChannelOfReadCacheData(entityChannel *goroutine.EntityChannel) {
	cacheValue := self.getCache()
	entityChannel.WriteOnce(cacheValue)
	return
}

func (self *EntityCacheDataRead) getCache() (cacheValue string) {
	r := self.CacheRedis
	ctx := self.CTX
	cacheKey := self.CacheKey
	cacheValue, errRedisGet := r.Get(ctx, cacheKey).Result()
	err.ErrCheck(errRedisGet)
	self.CacheValue = cacheValue
	return cacheValue
}
