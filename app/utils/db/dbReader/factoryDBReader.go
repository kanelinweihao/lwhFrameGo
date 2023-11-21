package dbReader

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/db/dbConnector"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/rfl"
)

func MakeEntityDBReader(entityDBConnector *dbConnector.EntityDBConnector, boxForDB base.BoxData) (entityDBReader *EntityDBReader) {
	entityDBReader = new(EntityDBReader)
	entityDBReader.Init(entityDBConnector, boxForDB)
	return entityDBReader
}

func (self *EntityDBReader) Init(entityDBConnector *dbConnector.EntityDBConnector, boxForDB base.BoxData) *EntityDBReader {
	self.setParamsIn(entityDBConnector, boxForDB).setParamsMore()
	return self
}

func (self *EntityDBReader) setParamsIn(entityDBConnector *dbConnector.EntityDBConnector, boxForDB base.BoxData) *EntityDBReader {
	self.EntityDBConnector = entityDBConnector
	self.BoxForDB = boxForDB
	return self
}

func (self *EntityDBReader) setParamsMore() *EntityDBReader {
	self.setFromBoxForDB().setAttrEntityDBReader()
	return self
}

func (self *EntityDBReader) setFromBoxForDB() *EntityDBReader {
	boxForDB := self.BoxForDB
	arrSQLName, ok := boxForDB["ArrSQLName"].([]string)
	if !ok {
		rfl.ErrPanicFormat(arrSQLName, "arrSQLName", "[]string")
	}
	attrArgsForQuery, ok := boxForDB["AttrArgsForQuery"].(base.AttrS1)
	if !ok {
		rfl.ErrPanicFormat(attrArgsForQuery, "attrArgsForQuery", "base.AttrS1")
	}
	self.ArrSQLName = arrSQLName
	self.AttrArgsForQuery = attrArgsForQuery
	return self
}

func (self *EntityDBReader) setAttrEntityDBReader() *EntityDBReader {
	dbSqlx := self.EntityDBConnector.DBSqlx
	arrSQLName := self.ArrSQLName
	attrArgsForQuery := self.AttrArgsForQuery
	attrEntityDBReader := make(map[string]*EntityDBDataRead)
	for _, sqlName := range arrSQLName {
		entityDBDataRead := MakeEntityDBDataRead(dbSqlx, sqlName, attrArgsForQuery)
		attrEntityDBReader[sqlName] = entityDBDataRead
	}
	self.AttrEntityDBDataRead = attrEntityDBReader
	return self
}
