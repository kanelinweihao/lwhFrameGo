package db

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/db/dbConnector"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/db/dbReader"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
)

type EntityDB struct {
	AttrT3DBData      typeMap.AttrT3
	ArrSQLName        []string
	AttrArgsForQuery  typeMap.AttrS1
	EntityDBConnector *dbConnector.EntityDBConnector
	EntityDBReader    *dbReader.EntityDBReader
}

func (self *EntityDB) BatchGetDataFromDB(arrSqlName []string, attrArgsForQuery typeMap.AttrS1) (attrT3DBData typeMap.AttrT3) {
	self.ArrSQLName = arrSqlName
	self.AttrArgsForQuery = attrArgsForQuery
	entityDBConnector := self.EntityDBConnector
	entityDBReader := dbReader.InitEntityDBReader(entityDBConnector, arrSqlName, attrArgsForQuery)
	self.EntityDBReader = entityDBReader
	attrT3DBData = entityDBReader.BatchReadDB()
	self.AttrT3DBData = attrT3DBData
	return attrT3DBData
}
