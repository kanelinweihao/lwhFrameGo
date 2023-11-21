package cacheWriter

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/cache/cacheConnector"
)

func MakeEntityCacheWriter(entityCacheConnector *cacheConnector.EntityCacheConnector, boxForCache base.BoxData) (entityCacheWriter *EntityCacheWriter) {
	entityCacheWriter = new(EntityCacheWriter)
	entityCacheWriter.Init(entityCacheConnector, boxForCache)
	return entityCacheWriter
}

func (self *EntityCacheWriter) Init(entityCacheConnector *cacheConnector.EntityCacheConnector, boxForCache base.BoxData) *EntityCacheWriter {
	self.setParamsIn(entityCacheConnector, boxForCache).setParamsMore()
	return self
}

func (self *EntityCacheWriter) setParamsIn(entityCacheConnector *cacheConnector.EntityCacheConnector, boxForCache base.BoxData) *EntityCacheWriter {
	self.EntityCacheConnector = entityCacheConnector
	self.BoxForCache = boxForCache
	return self
}

func (self *EntityCacheWriter) setParamsMore() *EntityCacheWriter {
	self.setAttrEntityCacheDataWrite()
	return self
}

func (self *EntityCacheWriter) setAttrEntityCacheDataWrite() *EntityCacheWriter {
	cacheRedis := self.EntityCacheConnector.CacheRedis
	boxForCache := self.BoxForCache
	attrEntityCacheData := make(map[string]*EntityCacheDataWrite)
	for cacheKey, attrT1ForCacheToAssign := range boxForCache {
		attrT1ForCache := attrT1ForCacheToAssign.(base.AttrT1)
		entityCacheData := MakeEntityCacheData(cacheRedis, attrT1ForCache)
		attrEntityCacheData[cacheKey] = entityCacheData
	}
	self.AttrEntityCacheDataWrite = attrEntityCacheData
	return self
}
