package dbReader

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/db/dbConnector"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/goroutine"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/rfl"
)

type EntityDBReader struct {
	AttrT3DBData         base.AttrT3
	EntityDBConnector    *dbConnector.EntityDBConnector
	ArrSQLName           []string
	AttrArgsForQuery     base.AttrS1
	AttrEntityDBDataRead map[string]*EntityDBDataRead
	AttrEntityChannel    base.AttrT1
}

func (self *EntityDBReader) BatchReadDB() (attrT3DBData base.AttrT3) {
	self.writeAttrEntityChannel().readAttrEntityChannel()
	attrT3DBData = self.AttrT3DBData
	return attrT3DBData
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
	attrT3DBData := make(base.AttrT3)
	for sqlName, entityChannelToAssign := range attrEntityChannel {
		entityChannel := entityChannelToAssign.(*goroutine.EntityChannel)
		attrT2DBData := readFromChannelOfReadDBData(entityChannel)
		attrT3DBData[sqlName] = attrT2DBData
	}
	self.AttrT3DBData = attrT3DBData
	return self
}

func readFromChannelOfReadDBData(entityChannel *goroutine.EntityChannel) (attrT2DBData base.AttrT2) {
	dataOnce := entityChannel.ReadOnce()
	attrT2DBData, ok := dataOnce.(base.AttrT2)
	if !ok {
		rfl.ErrPanicFormat(attrT2DBData, "attrT2DBData", "base.AttrT2")
	}
	return attrT2DBData
}
