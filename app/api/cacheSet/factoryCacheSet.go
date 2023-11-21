package cacheSet

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	ttt "github.com/kanelinweihao/lwhFrameGo/app/utils/time"
)

func MakeEntityCacheSet(arrPathFileExcel []string) (entityCacheSet *EntityCacheSet) {
	entityCacheSet = new(EntityCacheSet)
	entityCacheSet.Init(arrPathFileExcel)
	return entityCacheSet
}

func (self *EntityCacheSet) Init(arrPathFileExcel []string) *EntityCacheSet {
	self.setParamsIn(arrPathFileExcel).setParamsMore()
	return self
}

func (self *EntityCacheSet) setParamsIn(arrPathFileExcel []string) *EntityCacheSet {
	self.ArrPathFileExcel = arrPathFileExcel
	return self
}

func (self *EntityCacheSet) setParamsMore() *EntityCacheSet {
	self.setCacheKey().setCacheValue().setTTL().setAttrEntityForCache().setBoxForCache()
	return self
}

func (self *EntityCacheSet) setCacheKey() *EntityCacheSet {
	timeSuffix := ttt.GetTimeSuffix()
	cacheKeyPrefix := conf.GetProjectNameEN()
	cacheKey := fmt.Sprintf(
		"%s:%s",
		cacheKeyPrefix,
		timeSuffix)
	self.CacheKey = cacheKey
	return self
}

func (self *EntityCacheSet) setCacheValue() *EntityCacheSet {
	arrPathFileExcel := self.ArrPathFileExcel
	cacheValue := conv.ToJsonFromSlice(arrPathFileExcel)
	self.CacheValue = cacheValue
	return self
}

func (self *EntityCacheSet) setTTL() *EntityCacheSet {
	self.TTL = conf.TTL60s
	return self
}

func (self *EntityCacheSet) setAttrEntityForCache() *EntityCacheSet {
	entityForCache := new(EntityForCache)
	entityForCache.CacheKey = self.CacheKey
	entityForCache.CacheValue = self.CacheValue
	entityForCache.TTL = self.TTL
	attrEntityForCache := make(map[string]*EntityForCache)
	attrEntityForCache[self.CacheKey] = entityForCache
	self.AttrEntityForCache = attrEntityForCache
	return self
}

func (self *EntityCacheSet) setBoxForCache() *EntityCacheSet {
	attrEntityForCache := self.AttrEntityForCache
	boxForCache := make(base.BoxData)
	for cacheKey, entityForCache := range attrEntityForCache {
		attrT1ForCache := conv.ToAttrFromEntity(entityForCache)
		boxForCache[cacheKey] = attrT1ForCache
	}
	self.BoxForCache = boxForCache
	return self
}
