package db

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/db/dbConnector"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/db/dbQuerier"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/db/dbReader"
)

type EntityDB struct {
	BoxData           base.AttrS3
	EntityDBConnector *dbConnector.EntityDBConnector
	EntityDBQuerier   *dbQuerier.EntityDBQuerier
	EntityDBReader    *dbReader.EntityDBReader
}

func (self *EntityDB) GetDBData(arrSQLName []string, attrArgsForQuery base.AttrS1) (boxData base.AttrS3) {
	self.setEntityDBQuerier(
		arrSQLName,
		attrArgsForQuery)
	self.setEntityDBReader(
		self.EntityDBConnector.DBSqlx,
		self.EntityDBQuerier.AttrT2DBQuery)
	self.setBoxData()
	boxData = self.getBoxData()
	return boxData
}

func (self *EntityDB) setBoxData() *EntityDB {
	boxData := self.EntityDBReader.SetBoxData().GetBoxData()
	self.BoxData = boxData
	return self
}

func (self *EntityDB) getBoxData() (boxData base.AttrS3) {
	boxData = self.BoxData
	return boxData
}
