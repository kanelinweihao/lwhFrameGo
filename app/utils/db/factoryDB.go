package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/db/dbConnector"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/db/dbQuerier"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/db/dbReader"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
)

/*
Init
*/

func MakeEntityOfDB() (entityDB *EntityDB) {
	entityDB = initEntityDB()
	return entityDB
}

func initEntityDB() (entityDB *EntityDB) {
	entityDB = new(EntityDB)
	entityDB.Init()
	return entityDB
}

func (self *EntityDB) Init() *EntityDB {
	self.setEntityDBConnector()
	return self
}

func (self *EntityDB) setEntityDBConnector() *EntityDB {
	self.EntityDBConnector = dbConnector.MakeEntityOfDBConnector()
	return self
}

func (self *EntityDB) setEntityDBQuerier(arrSQLName []string, attrArgsForQuery base.AttrS1) *EntityDB {
	self.EntityDBQuerier = dbQuerier.MakeEntityOfDBQuerier(arrSQLName, attrArgsForQuery)
	return self
}

func (self *EntityDB) setEntityDBReader(dbSqlx *sqlx.DB, attrT2DBQuery base.AttrT2) *EntityDB {
	self.EntityDBReader = dbReader.MakeEntityOfDBReader(dbSqlx, attrT2DBQuery)
	return self
}

/*
Close
*/

func (self *EntityDB) CloseDB() {
	self.closeDBConnector()
	return
}

func (self *EntityDB) closeDBConnector() *EntityDB {
	if self.EntityDBConnector == nil {
		return self
	}
	self.EntityDBConnector.CloseDBConnector()
	self.EntityDBConnector = nil
	return self
}
