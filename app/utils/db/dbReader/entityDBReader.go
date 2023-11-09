package dbReader

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/goroutine"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/rfl"
)

type EntityDBReader struct {
	BoxData           base.AttrS3
	AttrEntityChannel base.AttrT1
	DBSqlx            *sqlx.DB
	AttrT2DBQuery     base.AttrT2
}

func (self *EntityDBReader) SetBoxData() *EntityDBReader {
	self.writeAttrEntityChannel().readAttrEntityChannel()
	return self
}

func (self *EntityDBReader) GetBoxData() (boxData base.AttrS3) {
	boxData = self.BoxData
	return boxData
}

/*
Write
*/

func (self *EntityDBReader) writeAttrEntityChannel() *EntityDBReader {
	d := self.DBSqlx
	attrT2DBQuery := self.AttrT2DBQuery
	attrEntityChannel := make(base.AttrT1)
	for sqlName, attrT1DBQuery := range attrT2DBQuery {
		entityDBData := attrT1DBQuery["entityDBData"].(base.EntityDBData)
		queryWithArgs := attrT1DBQuery["queryWithArgs"].(string)
		entityChannel := goroutine.MakeEntityOfGoroutine()
		attrEntityChannel[sqlName] = entityChannel
		go writeToChannelOfGetAttrS2DBData(
			entityChannel,
			d,
			entityDBData,
			queryWithArgs)
	}
	self.AttrEntityChannel = attrEntityChannel
	return self
}

func writeToChannelOfGetAttrS2DBData(entityChannel *goroutine.EntityChannel, d *sqlx.DB, entityDBData base.EntityDBData, queryWithArgs string) (attrS2DBData base.AttrS2) {
	attrS2DBData = getAttrS2DBData(
		d,
		entityDBData,
		queryWithArgs)
	entityChannel.WriteOnce(attrS2DBData)
	return
}

func getAttrS2DBData(d *sqlx.DB, entityDBData base.EntityDBData, queryWithArgs string) (attrS2DBData base.AttrS2) {
	ptrRows, errQueryx := d.Queryx(queryWithArgs)
	err.ErrCheck(errQueryx)
	attrS2DBData = make(base.AttrS2)
	num := 0
	for ptrRows.Next() {
		errScan := ptrRows.StructScan(entityDBData)
		err.ErrCheck(errScan)
		attrT1DBData := conv.ToAttrFromEntity(entityDBData)
		attrS1DBData := conv.ToAttrStrFromAttr(attrT1DBData)
		numStr := conv.ToStrFromInt(num)
		attrS2DBData[numStr] = attrS1DBData
		num++
	}
	return attrS2DBData
}

/*
Read
*/

func (self *EntityDBReader) readAttrEntityChannel() *EntityDBReader {
	attrEntityChannel := self.AttrEntityChannel
	boxData := make(base.AttrS3)
	for sqlName, entityChannelToAssign := range attrEntityChannel {
		entityChannel := entityChannelToAssign.(*goroutine.EntityChannel)
		attrS2DBData := readFromChannelOfGetAttrS2DBData(entityChannel)
		boxData[sqlName] = attrS2DBData
	}
	self.BoxData = boxData
	return self
}

func readFromChannelOfGetAttrS2DBData(entityChannel *goroutine.EntityChannel) (attrS2DBData base.AttrS2) {
	dataOnce := entityChannel.ReadOnce()
	attrS2DBData, ok := dataOnce.(base.AttrS2)
	if !ok {
		_, typeName, _ := rfl.GetTypeInfo(attrS2DBData)
		msgError := fmt.Sprintf(
			"The type of |%s| is |%s|, it should be |%s|",
			"attrS2DBData",
			typeName,
			"base.AttrS2")
		err.ErrPanic(msgError)
	}
	return attrS2DBData
}
