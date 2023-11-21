package cacheReader

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/cache/cacheConnector"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/goroutine"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/rfl"
)

type EntityCacheReader struct {
	BoxFromCache            base.BoxData
	EntityCacheConnector    *cacheConnector.EntityCacheConnector
	ArrCacheKey             []string
	AttrEntityCacheDataRead map[string]*EntityCacheDataRead
	AttrEntityChannel       base.AttrT1
}

func (self *EntityCacheReader) BatchGetCache() (boxFromCache base.BoxData) {
	self.writeAttrEntityChannel().readAttrEntityChannel()
	boxFromCache = self.BoxFromCache
	return boxFromCache
}

func (self *EntityCacheReader) writeAttrEntityChannel() *EntityCacheReader {
	attrEntityCacheData := self.AttrEntityCacheDataRead
	attrEntityChannel := make(base.AttrT1)
	for cacheKey, entityCacheDataRead := range attrEntityCacheData {
		entityChannel := goroutine.MakeEntityChannel()
		attrEntityChannel[cacheKey] = entityChannel
		go entityCacheDataRead.WriteToChannelOfReadCacheData(entityChannel)
	}
	self.AttrEntityChannel = attrEntityChannel
	return self
}

func (self *EntityCacheReader) readAttrEntityChannel() *EntityCacheReader {
	attrEntityChannel := self.AttrEntityChannel
	boxFromCache := make(base.BoxData)
	for cacheKey, entityChannelToAssign := range attrEntityChannel {
		entityChannel := entityChannelToAssign.(*goroutine.EntityChannel)
		cacheValue := readFromChannelOfReadCacheData(entityChannel)
		boxFromCache[cacheKey] = cacheValue
	}
	self.BoxFromCache = boxFromCache
	return self
}

func readFromChannelOfReadCacheData(entityChannel *goroutine.EntityChannel) (cacheValue string) {
	dataOnce := entityChannel.ReadOnce()
	cacheValue, ok := dataOnce.(string)
	if !ok {
		rfl.ErrPanicFormat(cacheValue, "cacheValue", "string")
	}
	return cacheValue
}
