package repoCacheBase

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/times"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

func (self *EntityRepoCacheBase) Load(entityRepoCacheDerived typeStruct.EntityRepoCache) typeStruct.EntityRepoCache {
	self.Derived = entityRepoCacheDerived
	return self.Derived
}

func (self *EntityRepoCacheBase) Init(arrPathFileExcel []string) typeStruct.EntityRepoCache {
	self.setPropertiesIn(arrPathFileExcel).setPropertiesMore()
	return self.Derived
}

func (self *EntityRepoCacheBase) setPropertiesIn(arrPathFileExcel []string) *EntityRepoCacheBase {
	self.ArrPathFileExcel = arrPathFileExcel
	return self
}

func (self *EntityRepoCacheBase) setPropertiesMore() *EntityRepoCacheBase {
	self.setCacheKey().setCacheValue().setTTL().setAttrEntityForCache().setBoxToCache()
	return self
}

func (self *EntityRepoCacheBase) setCacheKey() *EntityRepoCacheBase {
	timeSuffix := times.GetTimeSuffix()
	cacheKeyPrefix := conf.GetProjectNameEN()
	cacheKey := fmt.Sprintf(
		"%s:%s",
		cacheKeyPrefix,
		timeSuffix)
	self.CacheKey = cacheKey
	return self
}

func (self *EntityRepoCacheBase) setCacheValue() *EntityRepoCacheBase {
	arrPathFileExcel := self.ArrPathFileExcel
	cacheValue := conv.ToJsonFromSlice(arrPathFileExcel)
	self.CacheValue = cacheValue
	return self
}

func (self *EntityRepoCacheBase) setTTL() *EntityRepoCacheBase {
	self.TTL = conf.TTL60s
	return self
}

func (self *EntityRepoCacheBase) setAttrEntityForCache() *EntityRepoCacheBase {
	entityForCache := new(EntityForCache)
	entityForCache.CacheKey = self.CacheKey
	entityForCache.CacheValue = self.CacheValue
	entityForCache.TTL = self.TTL
	attrEntityForCache := make(map[string]*EntityForCache)
	attrEntityForCache[self.CacheKey] = entityForCache
	self.AttrEntityForCache = attrEntityForCache
	return self
}

func (self *EntityRepoCacheBase) setBoxToCache() *EntityRepoCacheBase {
	attrEntityForCache := self.AttrEntityForCache
	boxToCache := make(typeMap.BoxData)
	for cacheKey, entityForCache := range attrEntityForCache {
		attrT1ForCache := conv.ToAttrFromEntity(entityForCache)
		boxToCache[cacheKey] = attrT1ForCache
	}
	self.BoxToCache = boxToCache
	return self
}
