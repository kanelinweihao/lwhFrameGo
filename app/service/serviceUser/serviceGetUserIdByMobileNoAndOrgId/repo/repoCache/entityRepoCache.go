package repoCache

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/cache"
	"time"
)

type EntityRepoCache struct {
	BoxToCache         base.BoxData
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

func (self *EntityRepoCache) SetCache() (boxFromCache base.BoxData) {
	boxToCache := self.BoxToCache
	entityCache := cache.InitEntityCache()
	defer entityCache.CloseCache()
	arrCacheKey := entityCache.BatchSetDataToCache(boxToCache)
	boxFromCache = entityCache.BatchGetDataFromCache(arrCacheKey)
	return boxFromCache
}
