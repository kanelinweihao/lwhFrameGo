package cacheSet

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	ttt "github.com/kanelinweihao/lwhFrameGo/app/utils/time"
	"time"
)

var cacheKeyPrefix = "log"
var ttl60s = time.Second * time.Duration(60)

func MakeEntityOfCacheSet(paramsOut base.AttrT1) (entityCacheSet *EntityCacheSet) {
	entityCacheSet = initCacheSet(paramsOut)
	return entityCacheSet
}

func initCacheSet(paramsOut base.AttrT1) (entityCacheSet *EntityCacheSet) {
	entityCacheSet = new(EntityCacheSet)
	entityCacheSet.Init(paramsOut)
	return entityCacheSet
}

func (self *EntityCacheSet) Init(paramsOut base.AttrT1) *EntityCacheSet {
	self.setParamsOut(paramsOut).setCacheKey().setCacheValue().setTTL()
	return self
}

func (self *EntityCacheSet) setParamsOut(paramsOut base.AttrT1) *EntityCacheSet {
	self.ParamsOut = paramsOut
	return self
}

func (self *EntityCacheSet) setCacheKey() *EntityCacheSet {
	timeSuffix := ttt.GetTimeSuffix()
	moduleNameEN := conf.GetModuleNameEN()
	cacheKey := fmt.Sprintf(
		"%s:%s:%s",
		cacheKeyPrefix,
		moduleNameEN,
		timeSuffix)
	self.CacheKey = cacheKey
	return self
}

func (self *EntityCacheSet) setCacheValue() *EntityCacheSet {
	paramsOut := self.ParamsOut
	cacheValue := conv.ToJsonFromAttr(paramsOut)
	self.CacheValue = cacheValue
	return self
}

func (self *EntityCacheSet) setTTL() *EntityCacheSet {
	self.TTL = ttl60s
	return self
}
