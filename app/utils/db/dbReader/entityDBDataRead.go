package dbReader

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/goroutine"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/time"
)

type EntityDBDataRead struct {
	AttrT2DBData     base.AttrT2
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
	attrT2DBData := self.AttrT2DBData
	entityChannel.WriteOnce(attrT2DBData)
	return self
}

func (self *EntityDBDataRead) ReadDB() *EntityDBDataRead {
	d := self.DBSqlx
	queryWithArgs := self.QueryWithArgs
	entityDBData := self.EntityDBData
	ptrRows, errQueryx := d.Queryx(queryWithArgs)
	err.ErrCheck(errQueryx)
	attrT2DBData := make(base.AttrT2)
	num := 0
	for ptrRows.Next() {
		errScan := ptrRows.StructScan(entityDBData)
		err.ErrCheck(errScan)
		attrT1DBData := conv.ToAttrFromEntity(entityDBData)
		numStr := conv.ToStrFromInt(num)
		attrT2DBData[numStr] = attrT1DBData
		num++
	}
	if len(attrT2DBData) == 0 {
		msgEmptySQL := fmt.Sprintf(
			"The sql is invalid :\n|%s|\n",
			queryWithArgs)
		time.ShowTimeAndMsg(msgEmptySQL)
	}
	self.AttrT2DBData = attrT2DBData
	return self
}
