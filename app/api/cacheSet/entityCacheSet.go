package cacheSet

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/cache"
	"time"
)

type EntityCacheSet struct {
	BoxForCache        base.BoxData
	ArrPathFileExcel   []string
	CacheKey           string
	CacheValue         string
	TTL                time.Duration
	AttrEntityForCache map[string]*EntityForCache
}

type EntityForCache struct {
	CacheKey   string
	CacheValue string
	TTL        time.Duration
}

func (self *EntityCacheSet) SetCache() (boxFromCache base.BoxData) {
	boxForCache := self.BoxForCache
	entityCache := cache.MakeEntityCache()
	defer entityCache.CloseCache()
	arrCacheKey := entityCache.BatchSetDataToCache(boxForCache)
	boxFromCache = entityCache.BatchGetDataFromCache(arrCacheKey)
	return boxFromCache
}
