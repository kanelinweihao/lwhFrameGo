package cacheSet

import (
	"fmt"
	"go.lwh.com/linweihao/lwhFrameGo/app/conf"
	"go.lwh.com/linweihao/lwhFrameGo/app/factory/factoryOfCache"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/base"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/cache"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/conv"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
	ttt "go.lwh.com/linweihao/lwhFrameGo/app/utils/time"
	"time"
)

func SetCache(paramsOut base.AttrT1) {
	// entityCache := cache.InitCache()
	entityCache := factoryOfCache.MakeEntityOfCache()
	defer entityCache.CloseCache()
	setToCache(
		entityCache,
		paramsOut)
	ttt.ShowTimeAndMsg("Cache set success")
	return
}

func setToCache(entityCache *cache.EntityCache, paramsOut base.AttrT1) {
	cacheKey := getCacheKey()
	cacheValue := getCacheValue(paramsOut)
	ttl := time.Second * time.Duration(60)
	entityCache.SetToCache(
		cacheKey,
		cacheValue,
		ttl)
	return
}

func getCacheKey() (cacheKey string) {
	timeSuffix := ttt.GetTimeSuffix()
	moduleNameEN := conf.GetModuleNameEN()
	cacheKey = fmt.Sprintf(
		"%s:%s:%s",
		"log",
		moduleNameEN,
		timeSuffix)
	return cacheKey
}

func getCacheValue(paramsOut base.AttrT1) (cacheValue string) {
	cacheValue = conv.ToJsonFromAttr(paramsOut)
	// dd.DD(cacheValue)
	return cacheValue
}
