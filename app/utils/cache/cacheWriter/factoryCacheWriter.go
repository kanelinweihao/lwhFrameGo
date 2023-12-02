package cacheWriter

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/cache/cacheConnector"
)

func InitEntityCacheWriter(entityCacheConnector *cacheConnector.EntityCacheConnector, boxToCache base.BoxData) (entityCacheWriter *EntityCacheWriter) {
	entityCacheWriter = new(EntityCacheWriter)
	entityCacheWriter.Init(entityCacheConnector, boxToCache)
	return entityCacheWriter
}

func (self *EntityCacheWriter) Init(entityCacheConnector *cacheConnector.EntityCacheConnector, boxToCache base.BoxData) *EntityCacheWriter {
	self.setPropertiesIn(entityCacheConnector, boxToCache).setPropertiesMore()
	return self
}

func (self *EntityCacheWriter) setPropertiesIn(entityCacheConnector *cacheConnector.EntityCacheConnector, boxToCache base.BoxData) *EntityCacheWriter {
	self.EntityCacheConnector = entityCacheConnector
	self.BoxToCache = boxToCache
	return self
}

func (self *EntityCacheWriter) setPropertiesMore() *EntityCacheWriter {
	self.setAttrEntityCacheDataWrite()
	return self
}

func (self *EntityCacheWriter) setAttrEntityCacheDataWrite() *EntityCacheWriter {
	cacheRedis := self.EntityCacheConnector.CacheRedis
	boxToCache := self.BoxToCache
	attrEntityCacheData := make(map[string]*EntityCacheDataWrite)
	for cacheKey, attrT1ForCacheToAssign := range boxToCache {
		attrT1ForCache := attrT1ForCacheToAssign.(base.AttrT1)
		entityCacheData := InitEntityCacheData(cacheRedis, attrT1ForCache)
		attrEntityCacheData[cacheKey] = entityCacheData
	}
	self.AttrEntityCacheDataWrite = attrEntityCacheData
	return self
}
