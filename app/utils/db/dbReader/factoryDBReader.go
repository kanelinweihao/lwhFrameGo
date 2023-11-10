package dbReader

import (
	"github.com/jmoiron/sqlx"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

/*
Init
*/

func MakeEntityOfDBReader(dbSqlx *sqlx.DB, attrT2DBQuery base.AttrT2) (entityDBReader *EntityDBReader) {
	entityDBReader = initEntityDBReader(dbSqlx, attrT2DBQuery)
	return entityDBReader
}

func initEntityDBReader(dbSqlx *sqlx.DB, attrT2DBQuery base.AttrT2) (entityDBReader *EntityDBReader) {
	entityDBReader = new(EntityDBReader)
	entityDBReader.Init(dbSqlx, attrT2DBQuery)
	return entityDBReader
}

func (self *EntityDBReader) Init(dbSqlx *sqlx.DB, attrT2DBQuery base.AttrT2) *EntityDBReader {
	self.setDBSqlx(dbSqlx).setAttrEntityDBQuery(attrT2DBQuery)
	return self
}

func (self *EntityDBReader) setDBSqlx(dbSqlx *sqlx.DB) *EntityDBReader {
	self.DBSqlx = dbSqlx
	return self
}

func (self *EntityDBReader) setAttrEntityDBQuery(attrT2DBQuery base.AttrT2) *EntityDBReader {
	self.AttrT2DBQuery = attrT2DBQuery
	return self
}
