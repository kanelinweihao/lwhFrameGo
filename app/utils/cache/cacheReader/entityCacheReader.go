package cacheReader

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/cache/cacheConnector"
)

type typeChanData = string

type EntityCacheReader struct {
	BoxFromCache            base.BoxData
	EntityCacheConnector    *cacheConnector.EntityCacheConnector
	ArrCacheKey             []string
	AttrEntityCacheDataRead map[string]*EntityCacheDataRead
	AttrChan                map[string]chan typeChanData
}

func (self *EntityCacheReader) BatchGetCache() (boxFromCache base.BoxData) {
	self.writeAttrChan().readAttrChan()
	boxFromCache = self.BoxFromCache
	return boxFromCache
}

func (self *EntityCacheReader) writeAttrChan() *EntityCacheReader {
	attrEntityCacheData := self.AttrEntityCacheDataRead
	attrChan := make(map[string]chan typeChanData, len(attrEntityCacheData))
	for cacheKey, entityCacheDataRead := range attrEntityCacheData {
		chanBase := make(chan typeChanData, 1)
		attrChan[cacheKey] = chanBase
		go entityCacheDataRead.WriteToChannelOfReadCacheData(chanBase)
	}
	self.AttrChan = attrChan
	return self
}

func (self *EntityCacheReader) readAttrChan() *EntityCacheReader {
	attrChan := self.AttrChan
	boxFromCache := make(base.BoxData)
	for cacheKey, chanBase := range attrChan {
		cacheValue := readFromChannelOfReadCacheData(chanBase)
		boxFromCache[cacheKey] = cacheValue
	}
	self.BoxFromCache = boxFromCache
	return self
}
