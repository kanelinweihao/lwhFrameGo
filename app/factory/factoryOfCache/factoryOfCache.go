package factoryOfCache

import (
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/cache"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
)

func MakeEntityOfCache() (entityCache *cache.EntityCache) {
	entityCache = cache.InitCache()
	return entityCache
}
