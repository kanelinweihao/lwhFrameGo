package cache

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/cache/cacheConnector"
)

/*
Init
*/

func MakeEntityCache() (entityCache *EntityCache) {
	entityCache = new(EntityCache)
	entityCache.Init()
	return entityCache
}

func (self *EntityCache) Init() *EntityCache {
	self.setEntityEmailConnector()
	return self
}

func (self *EntityCache) setEntityEmailConnector() *EntityCache {
	entityCacheConnector := cacheConnector.MakeEntityCacheConnector()
	self.EntityCacheConnector = entityCacheConnector
	return self
}

/*
Close
*/

func (self *EntityCache) CloseCache() {
	entityCacheConnector := self.EntityCacheConnector
	entityCacheConnector.CloseCacheConnector()
	self.EntityCacheConnector = nil
	return
}
