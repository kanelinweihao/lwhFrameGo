package cacheWriter

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/cache/cacheConnector"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/goroutine"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/rfl"
)

type EntityCacheWriter struct {
	ArrCacheKey              []string
	BoxToCache               base.BoxData
	EntityCacheConnector     *cacheConnector.EntityCacheConnector
	AttrEntityCacheDataWrite map[string]*EntityCacheDataWrite
	AttrEntityChannel        base.AttrT1
}

func (self *EntityCacheWriter) BatchSetCache() (arrCacheKey []string) {
	self.writeAttrEntityChannel().readAttrEntityChannel()
	arrCacheKey = self.ArrCacheKey
	return arrCacheKey
}

func (self *EntityCacheWriter) writeAttrEntityChannel() *EntityCacheWriter {
	attrEntityCacheDataWrite := self.AttrEntityCacheDataWrite
	attrEntityChannel := make(base.AttrT1)
	for cacheKey, entityCacheDataWrite := range attrEntityCacheDataWrite {
		entityChannel := goroutine.MakeEntityChannel()
		attrEntityChannel[cacheKey] = entityChannel
		go entityCacheDataWrite.WriteToChannelOfWriteCacheData(entityChannel)
	}
	self.AttrEntityChannel = attrEntityChannel
	return self
}

func (self *EntityCacheWriter) readAttrEntityChannel() *EntityCacheWriter {
	attrEntityChannel := self.AttrEntityChannel
	arrCacheKey := make([]string, 0, len(attrEntityChannel))
	for _, entityChannelToAssign := range attrEntityChannel {
		entityChannel := entityChannelToAssign.(*goroutine.EntityChannel)
		cacheKey := readFromChannelOfWriteCacheData(entityChannel)
		arrCacheKey = append(arrCacheKey, cacheKey)
	}
	self.ArrCacheKey = arrCacheKey
	return self
}

func readFromChannelOfWriteCacheData(entityChannel *goroutine.EntityChannel) (cacheKey string) {
	dataOnce := entityChannel.ReadOnce()
	cacheKey, ok := dataOnce.(string)
	if !ok {
		rfl.ErrPanicFormat(cacheKey, "cacheKey", "string")
	}
	return cacheKey
}
