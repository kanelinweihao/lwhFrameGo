package cacheWriter

import (
	"context"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
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

func (self *EntityCacheDataWrite) WriteToChannelOfWriteCacheData(chanWrite chan<- typeChanData) {
	self.setCache().writeChan(chanWrite)
	return
}

func (self *EntityCacheDataWrite) setCache() *EntityCacheDataWrite {
	r := self.CacheRedis
	ctx := self.CTX
	cacheKey := self.CacheKey
	cacheValue := self.CacheValue
	ttl := self.TTL
	errRedisSet := r.Set(ctx, cacheKey, cacheValue, ttl).Err()
	err.ErrCheck(errRedisSet)
	return self
}

func (self *EntityCacheDataWrite) writeChan(chanWrite chan<- typeChanData) *EntityCacheDataWrite {
	cacheKey := self.CacheKey
	chanWrite <- cacheKey
	close(chanWrite)
	return self
}

func readFromChannelOfWriteCacheData(chanRead <-chan typeChanData) (cacheKey string) {
	cacheKey = <-chanRead
	return cacheKey
}
