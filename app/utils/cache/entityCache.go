package cache

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/cache/cacheConnector"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/cache/cacheReader"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/cache/cacheWriter"
)

var ctx = conf.CTX

type EntityCache struct {
	BoxFromCache         base.BoxData
	EntityCacheConnector *cacheConnector.EntityCacheConnector
	EntityCacheWriter    *cacheWriter.EntityCacheWriter
	EntityCacheReader    *cacheReader.EntityCacheReader
	BoxToCache           base.BoxData
}

func (self *EntityCache) BatchSetDataToCache(boxToCache base.BoxData) (arrCacheKey []string) {
	self.BoxToCache = boxToCache
	entityCacheConnector := self.EntityCacheConnector
	entityCacheWriter := cacheWriter.MakeEntityCacheWriter(entityCacheConnector, boxToCache)
	self.EntityCacheWriter = entityCacheWriter
	arrCacheKey = entityCacheWriter.BatchSetCache()
	return arrCacheKey
}

func (self *EntityCache) BatchGetDataFromCache(arrCacheKey []string) (boxFromCache base.BoxData) {
	entityCacheConnector := self.EntityCacheConnector
	entityCacheReader := cacheReader.MakeEntityCacheReader(entityCacheConnector, arrCacheKey)
	self.EntityCacheReader = entityCacheReader
	boxFromCache = entityCacheReader.BatchGetCache()
	self.BoxFromCache = boxFromCache
	return boxFromCache
}
