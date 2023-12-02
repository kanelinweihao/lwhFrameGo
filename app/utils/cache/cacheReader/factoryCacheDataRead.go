package cacheReader

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/redis/go-redis/v9"
)

func InitEntityCacheDataRead(cacheRedis *redis.Client, cacheKey string) (entityCacheDataRead *EntityCacheDataRead) {
	entityCacheDataRead = new(EntityCacheDataRead)
	entityCacheDataRead.Init(cacheRedis, cacheKey)
	return entityCacheDataRead
}

func (self *EntityCacheDataRead) Init(cacheRedis *redis.Client, cacheKey string) *EntityCacheDataRead {
	self.setPropertiesIn(cacheRedis, cacheKey).setPropertiesMore()
	return self
}

func (self *EntityCacheDataRead) setPropertiesIn(cacheRedis *redis.Client, cacheKey string) *EntityCacheDataRead {
	self.CacheRedis = cacheRedis
	self.CacheKey = cacheKey
	return self
}

func (self *EntityCacheDataRead) setPropertiesMore() *EntityCacheDataRead {
	self.setCTX()
	return self
}

func (self *EntityCacheDataRead) setCTX() *EntityCacheDataRead {
	self.CTX = conf.CTX
	return self
}
