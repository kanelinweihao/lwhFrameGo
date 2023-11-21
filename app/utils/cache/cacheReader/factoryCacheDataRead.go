package cacheReader

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/redis/go-redis/v9"
)

func MakeEntityCacheDataRead(cacheRedis *redis.Client, cacheKey string) (entityCacheDataRead *EntityCacheDataRead) {
	entityCacheDataRead = new(EntityCacheDataRead)
	entityCacheDataRead.Init(cacheRedis, cacheKey)
	return entityCacheDataRead
}

func (self *EntityCacheDataRead) Init(cacheRedis *redis.Client, cacheKey string) *EntityCacheDataRead {
	self.setParamsIn(cacheRedis, cacheKey).setParamsMore()
	return self
}

func (self *EntityCacheDataRead) setParamsIn(cacheRedis *redis.Client, cacheKey string) *EntityCacheDataRead {
	self.CacheRedis = cacheRedis
	self.CacheKey = cacheKey
	return self
}

func (self *EntityCacheDataRead) setParamsMore() *EntityCacheDataRead {
	self.setCTX()
	return self
}

func (self *EntityCacheDataRead) setCTX() *EntityCacheDataRead {
	self.CTX = conf.CTX
	return self
}
