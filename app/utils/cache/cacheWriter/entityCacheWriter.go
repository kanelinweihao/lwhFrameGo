package cacheWriter

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/cache/cacheConnector"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
)

type typeChanData = string

type EntityCacheWriter struct {
	ArrCacheKey              []string
	BoxToCache               typeMap.BoxData
	EntityCacheConnector     *cacheConnector.EntityCacheConnector
	AttrEntityCacheDataWrite map[string]*EntityCacheDataWrite
	AttrChan                 map[string]chan typeChanData
}

func (self *EntityCacheWriter) BatchSetCache() (arrCacheKey []string) {
	self.writeAttrChan().readAttrChan()
	arrCacheKey = self.ArrCacheKey
	return arrCacheKey
}

func (self *EntityCacheWriter) writeAttrChan() *EntityCacheWriter {
	attrEntityCacheDataWrite := self.AttrEntityCacheDataWrite
	attrChan := make(map[string]chan typeChanData, len(attrEntityCacheDataWrite))
	for cacheKey, entityCacheDataWrite := range attrEntityCacheDataWrite {
		chanBase := make(chan typeChanData, 1)
		attrChan[cacheKey] = chanBase
		go entityCacheDataWrite.WriteToChannelOfWriteCacheData(chanBase)
	}
	self.AttrChan = attrChan
	return self
}

func (self *EntityCacheWriter) readAttrChan() *EntityCacheWriter {
	attrChan := self.AttrChan
	arrCacheKey := make([]string, 0, len(attrChan))
	for _, chanBase := range attrChan {
		cacheKey := readFromChannelOfWriteCacheData(chanBase)
		arrCacheKey = append(arrCacheKey, cacheKey)
	}
	self.ArrCacheKey = arrCacheKey
	return self
}
