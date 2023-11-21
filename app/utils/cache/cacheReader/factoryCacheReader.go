package cacheReader

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/cache/cacheConnector"
)

func MakeEntityCacheReader(entityCacheConnector *cacheConnector.EntityCacheConnector, arrCacheKey []string) (entityCacheReader *EntityCacheReader) {
	entityCacheReader = new(EntityCacheReader)
	entityCacheReader.Init(entityCacheConnector, arrCacheKey)
	return entityCacheReader
}

func (self *EntityCacheReader) Init(entityCacheConnector *cacheConnector.EntityCacheConnector, arrCacheKey []string) *EntityCacheReader {
	self.setParamsIn(entityCacheConnector, arrCacheKey).setParamsMore()
	return self
}

func (self *EntityCacheReader) setParamsIn(entityCacheConnector *cacheConnector.EntityCacheConnector, arrCacheKey []string) *EntityCacheReader {
	self.EntityCacheConnector = entityCacheConnector
	self.ArrCacheKey = arrCacheKey
	return self
}

func (self *EntityCacheReader) setParamsMore() *EntityCacheReader {
	self.setAttrEntityCacheDataRead()
	return self
}

func (self *EntityCacheReader) setAttrEntityCacheDataRead() *EntityCacheReader {
	cacheRedis := self.EntityCacheConnector.CacheRedis
	arrCacheKey := self.ArrCacheKey
	attrEntityCacheDataRead := make(map[string]*EntityCacheDataRead)
	for _, cacheKey := range arrCacheKey {
		entityCacheDataRead := MakeEntityCacheDataRead(cacheRedis, cacheKey)
		attrEntityCacheDataRead[cacheKey] = entityCacheDataRead
	}
	self.AttrEntityCacheDataRead = attrEntityCacheDataRead
	return self
}
