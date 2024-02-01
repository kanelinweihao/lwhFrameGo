package repoCacheBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/cache"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
	"time"
)

type EntityRepoCacheBase struct {
	Derived    typeInterface.EntityRepoCache
	BoxToCache typeMap.BoxData
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

func (self *EntityRepoCacheBase) ToAttr() (paramsOut typeMap.AttrT1) {
	paramsOut = conv.ToAttrFromEntity(self)
	return paramsOut
}

func (self *EntityRepoCacheBase) SetCache() (boxFromCache typeMap.BoxData) {
	boxToCache := self.BoxToCache
	entityCache := cache.InitEntityCache()
	defer entityCache.CloseCache()
	arrCacheKey := entityCache.BatchSetDataToCache(boxToCache)
	boxFromCache = entityCache.BatchGetDataFromCache(arrCacheKey)
	return boxFromCache
}
