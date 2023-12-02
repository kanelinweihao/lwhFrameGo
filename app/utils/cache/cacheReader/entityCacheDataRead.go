package cacheReader

import (
	"context"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
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

func (self *EntityCacheDataRead) WriteToChannelOfReadCacheData(chanWrite chan<- typeChanData) {
	self.getCache().writeChan(chanWrite)
	return
}

func (self *EntityCacheDataRead) getCache() *EntityCacheDataRead {
	r := self.CacheRedis
	ctx := self.CTX
	cacheKey := self.CacheKey
	cacheValue, errRedisGet := r.Get(ctx, cacheKey).Result()
	err.ErrCheck(errRedisGet)
	self.CacheValue = cacheValue
	return self
}

func (self *EntityCacheDataRead) writeChan(chanWrite chan<- typeChanData) *EntityCacheDataRead {
	cacheValue := self.CacheValue
	chanWrite <- cacheValue
	close(chanWrite)
	return self
}

func readFromChannelOfReadCacheData(chanRead <-chan typeChanData) (cacheValue string) {
	cacheValue = <-chanRead
	return cacheValue
}
