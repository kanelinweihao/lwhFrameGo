package dbReader

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/db/dbConnector"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
)

type typeChanData = typeMap.AttrT2

type EntityDBReader struct {
	AttrT3DBData         typeMap.AttrT3
	EntityDBConnector    *dbConnector.EntityDBConnector
	ArrSQLName           []string
	AttrArgsForQuery     typeMap.AttrS1
	AttrEntityDBDataRead map[string]*EntityDBDataRead
	AttrChan             map[string]chan typeChanData
}

func (self *EntityDBReader) BatchReadDB() (attrT3DBData typeMap.AttrT3) {
	self.writeAttrChan().readAttrChan()
	attrT3DBData = self.AttrT3DBData
	return attrT3DBData
}

func (self *EntityDBReader) writeAttrChan() *EntityDBReader {
	attrEntityDBDataRead := self.AttrEntityDBDataRead
	attrChan := make(map[string]chan typeChanData, len(attrEntityDBDataRead))
	for sqlName, entityDBDataRead := range attrEntityDBDataRead {
		chanBase := make(chan typeChanData, 1)
		attrChan[sqlName] = chanBase
		go entityDBDataRead.writeToChannelOfReadDBData(chanBase)
	}
	self.AttrChan = attrChan
	return self
}

func (self *EntityDBReader) readAttrChan() *EntityDBReader {
	attrChan := self.AttrChan
	attrT3DBData := make(typeMap.AttrT3)
	for sqlName, chanBase := range attrChan {
		attrT2DBData := readFromChannelOfReadDBData(chanBase)
		attrT3DBData[sqlName] = attrT2DBData
	}
	self.AttrT3DBData = attrT3DBData
	return self
}
