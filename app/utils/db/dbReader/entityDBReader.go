package dbReader

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/db/dbConnector"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/goroutine"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/rfl"
)

type EntityDBReader struct {
	BoxFromDB            base.AttrS3
	EntityDBConnector    *dbConnector.EntityDBConnector
	BoxForDB             base.BoxData
	ArrSQLName           []string
	AttrArgsForQuery     base.AttrS1
	AttrEntityDBDataRead map[string]*EntityDBDataRead
	AttrEntityChannel    base.AttrT1
}

func (self *EntityDBReader) BatchReadDB() (boxFromDB base.AttrS3) {
	self.writeAttrEntityChannel().readAttrEntityChannel()
	boxFromDB = self.BoxFromDB
	return boxFromDB
}

func (self *EntityDBReader) writeAttrEntityChannel() *EntityDBReader {
	attrEntityDBDataRead := self.AttrEntityDBDataRead
	attrEntityChannel := make(base.AttrT1)
	for sqlName, entityDBDataRead := range attrEntityDBDataRead {
		entityChannel := goroutine.MakeEntityChannel()
		attrEntityChannel[sqlName] = entityChannel
		go entityDBDataRead.writeToChannelOfReadDBData(entityChannel)
	}
	self.AttrEntityChannel = attrEntityChannel
	return self
}

func (self *EntityDBReader) readAttrEntityChannel() *EntityDBReader {
	attrEntityChannel := self.AttrEntityChannel
	boxFromDB := make(base.AttrS3)
	for sqlName, entityChannelToAssign := range attrEntityChannel {
		entityChannel := entityChannelToAssign.(*goroutine.EntityChannel)
		attrS2DBData := readFromChannelOfReadDBData(entityChannel)
		boxFromDB[sqlName] = attrS2DBData
	}
	self.BoxFromDB = boxFromDB
	return self
}

func readFromChannelOfReadDBData(entityChannel *goroutine.EntityChannel) (attrS2DBData base.AttrS2) {
	dataOnce := entityChannel.ReadOnce()
	attrS2DBData, ok := dataOnce.(base.AttrS2)
	if !ok {
		rfl.ErrPanicFormat(attrS2DBData, "attrS2DBData", "base.AttrS2")
	}
	return attrS2DBData
}
