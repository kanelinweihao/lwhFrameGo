package repoCache

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/base/repoBase/repoCacheBase"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

func InitEntityRepoCache(arrPathFileExcel []string) (entityRepoCache typeStruct.EntityRepoCache) {
	entityRepoCache = new(EntityRepoCache)
	entityRepoCacheBase := new(repoCacheBase.EntityRepoCacheBase)
	entityRepoCache.Load(entityRepoCacheBase).Init(arrPathFileExcel)
	return entityRepoCache
}

func (self *EntityRepoCache) Load(entityRepoCacheBase typeStruct.EntityRepoCache) typeStruct.EntityRepoCache {
	self.EntityRepoCache = entityRepoCacheBase
	entityRepoCacheBase.Load(self)
	return self
}
