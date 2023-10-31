package cacheSet

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
)

func SetCache(paramsOut base.AttrT1) {
	entityCacheSet := MakeEntityOfCacheSet(paramsOut)
	entityCacheSet.SetToRedis()
	entityCacheSet.GetFromRedis()
	return
}
