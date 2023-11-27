package repoCache

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	ttt "github.com/kanelinweihao/lwhFrameGo/app/utils/time"
)

func MakeEntityRepoCache(arrPathFileExcel []string) (entityRepoCache *EntityRepoCache) {
	entityRepoCache = new(EntityRepoCache)
	entityRepoCache.Init(arrPathFileExcel)
	return entityRepoCache
}

func (self *EntityRepoCache) Init(arrPathFileExcel []string) *EntityRepoCache {
	self.setPropertiesIn(arrPathFileExcel).setPropertiesMore()
	return self
}

func (self *EntityRepoCache) setPropertiesIn(arrPathFileExcel []string) *EntityRepoCache {
	self.ArrPathFileExcel = arrPathFileExcel
	return self
}

func (self *EntityRepoCache) setPropertiesMore() *EntityRepoCache {
	self.setCacheKey().setCacheValue().setTTL().setAttrEntityForCache().setBoxToCache()
	return self
}

func (self *EntityRepoCache) setCacheKey() *EntityRepoCache {
	timeSuffix := ttt.GetTimeSuffix()
	cacheKeyPrefix := conf.GetProjectNameEN()
	cacheKey := fmt.Sprintf(
		"%s:%s",
		cacheKeyPrefix,
		timeSuffix)
	self.CacheKey = cacheKey
	return self
}

func (self *EntityRepoCache) setCacheValue() *EntityRepoCache {
	arrPathFileExcel := self.ArrPathFileExcel
	cacheValue := conv.ToJsonFromSlice(arrPathFileExcel)
	self.CacheValue = cacheValue
	return self
}

func (self *EntityRepoCache) setTTL() *EntityRepoCache {
	self.TTL = conf.TTL60s
	return self
}

func (self *EntityRepoCache) setAttrEntityForCache() *EntityRepoCache {
	entityForCache := new(EntityForCache)
	entityForCache.CacheKey = self.CacheKey
	entityForCache.CacheValue = self.CacheValue
	entityForCache.TTL = self.TTL
	attrEntityForCache := make(map[string]*EntityForCache)
	attrEntityForCache[self.CacheKey] = entityForCache
	self.AttrEntityForCache = attrEntityForCache
	return self
}

func (self *EntityRepoCache) setBoxToCache() *EntityRepoCache {
	attrEntityForCache := self.AttrEntityForCache
	boxToCache := make(base.BoxData)
	for cacheKey, entityForCache := range attrEntityForCache {
		attrT1ForCache := conv.ToAttrFromEntity(entityForCache)
		boxToCache[cacheKey] = attrT1ForCache
	}
	self.BoxToCache = boxToCache
	return self
}
