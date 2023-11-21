package dbReader

import (
	"github.com/jmoiron/sqlx"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/goroutine"
)

type EntityDBDataRead struct {
	AttrS2DBData     base.AttrS2
	DBSqlx           *sqlx.DB
	SQLName          string
	AttrArgsForQuery base.AttrS1
	ArrArgsForQuery  []string
	PathDirSQL       string
	QueryWithoutArgs string
	QueryWithArgs    string
	EntityDBData     base.EntityDBData
}

func (self *EntityDBDataRead) writeToChannelOfReadDBData(entityChannel *goroutine.EntityChannel) *EntityDBDataRead {
	self.ReadDB()
	attrS2DBData := self.AttrS2DBData
	entityChannel.WriteOnce(attrS2DBData)
	return self
}

func (self *EntityDBDataRead) ReadDB() *EntityDBDataRead {
	d := self.DBSqlx
	queryWithArgs := self.QueryWithArgs
	entityDBData := self.EntityDBData
	ptrRows, errQueryx := d.Queryx(queryWithArgs)
	err.ErrCheck(errQueryx)
	attrS2DBData := make(base.AttrS2)
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
	self.AttrS2DBData = attrS2DBData
	return self
}
