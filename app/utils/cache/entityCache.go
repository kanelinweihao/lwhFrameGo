package cache

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/cache/cacheConnector"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/cache/cacheReader"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/cache/cacheWriter"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
)

type EntityCache struct {
	BoxFromCache         typeMap.BoxData
	EntityCacheConnector *cacheConnector.EntityCacheConnector
	EntityCacheWriter    *cacheWriter.EntityCacheWriter
	EntityCacheReader    *cacheReader.EntityCacheReader
	BoxToCache           typeMap.BoxData
}

func (self *EntityCache) BatchSetDataToCache(boxToCache typeMap.BoxData) (arrCacheKey []string) {
	self.BoxToCache = boxToCache
	entityCacheConnector := self.EntityCacheConnector
	entityCacheWriter := cacheWriter.InitEntityCacheWriter(entityCacheConnector, boxToCache)
	self.EntityCacheWriter = entityCacheWriter
	arrCacheKey = entityCacheWriter.BatchSetCache()
	return arrCacheKey
}

func (self *EntityCache) BatchGetDataFromCache(arrCacheKey []string) (boxFromCache typeMap.BoxData) {
	entityCacheConnector := self.EntityCacheConnector
	entityCacheReader := cacheReader.InitEntityCacheReader(entityCacheConnector, arrCacheKey)
	self.EntityCacheReader = entityCacheReader
	boxFromCache = entityCacheReader.BatchGetCache()
	self.BoxFromCache = boxFromCache
	return boxFromCache
}
