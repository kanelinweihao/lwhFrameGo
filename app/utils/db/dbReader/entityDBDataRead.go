package dbReader

import (
	"github.com/jmoiron/sqlx"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

type EntityDBDataRead struct {
	AttrT2DBData     typeMap.AttrT2
	DBSqlx           *sqlx.DB
	SQLName          string
	AttrArgsForQuery typeMap.AttrS1
	ArrArgsForQuery  []string
	PathQuery        string
	QueryWithoutArgs string
	QueryWithArgs    string
	EntityDBData     typeStruct.EntityDBData
}

func (self *EntityDBDataRead) writeToChannelOfReadDBData(chanWrite chan<- typeChanData) *EntityDBDataRead {
	self.readDB().writeChan(chanWrite)
	return self
}

func (self *EntityDBDataRead) readDB() *EntityDBDataRead {
	d := self.DBSqlx
	queryWithArgs := self.QueryWithArgs
	entityDBData := self.EntityDBData
	ptrRows, errQueryx := d.Queryx(queryWithArgs)
	err.ErrCheck(errQueryx)
	attrT2DBData := make(typeMap.AttrT2)
	num := 0
	for ptrRows.Next() {
		errScan := ptrRows.StructScan(entityDBData)
		err.ErrCheck(errScan)
		attrT1DBData := conv.ToAttrFromEntity(entityDBData)
		numStr := conv.ToStrFromInt(num)
		attrT2DBData[numStr] = attrT1DBData
		num++
	}
	// if len(attrT2DBData) == 0 {
	// 	msgEmptySQL := fmt.Sprintf(
	// 		"The data of sql is empty :\n|%s|\n",
	// 		queryWithArgs)
	// 	times.ShowTimeAndMsg(msgEmptySQL)
	// }
	self.AttrT2DBData = attrT2DBData
	return self
}

func (self *EntityDBDataRead) writeChan(chanWrite chan<- typeChanData) *EntityDBDataRead {
	attrT2DBData := self.AttrT2DBData
	chanWrite <- attrT2DBData
	close(chanWrite)
	return self
}

func readFromChannelOfReadDBData(chanRead <-chan typeChanData) (attrT2DBData typeMap.AttrT2) {
	attrT2DBData = <-chanRead
	return attrT2DBData
}
