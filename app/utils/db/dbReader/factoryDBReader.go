package dbReader

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/db/dbConnector"
)

func InitEntityDBReader(entityDBConnector *dbConnector.EntityDBConnector, arrSqlName []string, attrArgsForQuery base.AttrS1) (entityDBReader *EntityDBReader) {
	entityDBReader = new(EntityDBReader)
	entityDBReader.Init(entityDBConnector, arrSqlName, attrArgsForQuery)
	return entityDBReader
}

func (self *EntityDBReader) Init(entityDBConnector *dbConnector.EntityDBConnector, arrSqlName []string, attrArgsForQuery base.AttrS1) *EntityDBReader {
	self.setPropertiesIn(entityDBConnector, arrSqlName, attrArgsForQuery).setPropertiesMore()
	return self
}

func (self *EntityDBReader) setPropertiesIn(entityDBConnector *dbConnector.EntityDBConnector, arrSqlName []string, attrArgsForQuery base.AttrS1) *EntityDBReader {
	self.EntityDBConnector = entityDBConnector
	self.ArrSQLName = arrSqlName
	self.AttrArgsForQuery = attrArgsForQuery
	return self
}

func (self *EntityDBReader) setPropertiesMore() *EntityDBReader {
	self.setAttrEntityDBReader()
	return self
}

func (self *EntityDBReader) setAttrEntityDBReader() *EntityDBReader {
	dbSqlx := self.EntityDBConnector.DBSqlx
	arrSQLName := self.ArrSQLName
	attrArgsForQuery := self.AttrArgsForQuery
	attrEntityDBReader := make(map[string]*EntityDBDataRead)
	for _, sqlName := range arrSQLName {
		entityDBDataRead := InitEntityDBDataRead(dbSqlx, sqlName, attrArgsForQuery)
		attrEntityDBReader[sqlName] = entityDBDataRead
	}
	self.AttrEntityDBDataRead = attrEntityDBReader
	return self
}
