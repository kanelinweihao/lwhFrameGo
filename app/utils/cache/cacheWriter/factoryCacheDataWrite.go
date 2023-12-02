package cacheWriter

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/redis/go-redis/v9"
)

func InitEntityCacheData(cacheRedis *redis.Client, attrT1ForCache base.AttrT1) (entityCacheData *EntityCacheDataWrite) {
	entityCacheData = new(EntityCacheDataWrite)
	entityCacheData.Init(cacheRedis, attrT1ForCache)
	return entityCacheData
}

func (self *EntityCacheDataWrite) Init(cacheRedis *redis.Client, attrT1ForCache base.AttrT1) *EntityCacheDataWrite {
	self.setPropertiesIn(cacheRedis, attrT1ForCache).setPropertiesMore()
	return self
}

func (self *EntityCacheDataWrite) setPropertiesIn(cacheRedis *redis.Client, attrT1ForCache base.AttrT1) *EntityCacheDataWrite {
	conv.ToEntityFromAttr(attrT1ForCache, self)
	self.CacheRedis = cacheRedis
	return self
}

func (self *EntityCacheDataWrite) setPropertiesMore() *EntityCacheDataWrite {
	self.setCTX()
	return self
}

func (self *EntityCacheDataWrite) setCTX() *EntityCacheDataWrite {
	self.CTX = conf.CTX
	return self
}
