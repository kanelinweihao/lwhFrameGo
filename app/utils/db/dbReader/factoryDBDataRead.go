package dbReader

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/sqlInfo"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/pack"
	"strings"
)

func MakeEntityDBDataRead(dbSqlx *sqlx.DB, sqlName string, attrArgsForQuery base.AttrS1) (entityDBDataRead *EntityDBDataRead) {
	entityDBDataRead = new(EntityDBDataRead)
	entityDBDataRead.Init(dbSqlx, sqlName, attrArgsForQuery)
	return entityDBDataRead
}

func (self *EntityDBDataRead) Init(dbSqlx *sqlx.DB, sqlName string, attrArgsForQuery base.AttrS1) *EntityDBDataRead {
	self.setParamsIn(dbSqlx, sqlName, attrArgsForQuery).setParamsMore()
	return self
}

func (self *EntityDBDataRead) setParamsIn(dbSqlx *sqlx.DB, sqlName string, attrArgsForQuery base.AttrS1) *EntityDBDataRead {
	self.DBSqlx = dbSqlx
	self.SQLName = sqlName
	self.AttrArgsForQuery = attrArgsForQuery
	return self
}

func (self *EntityDBDataRead) setParamsMore() *EntityDBDataRead {
	self.setPathDirSQL().setQueryOriginal().setArrArgsForQuery().setQueryWithArgs().setEntityDBData()
	return self
}

func (self *EntityDBDataRead) setPathDirSQL() *EntityDBDataRead {
	sqlName := self.SQLName
	pathDirSQL := conf.GetPathDirSQL(sqlName)
	self.PathDirSQL = pathDirSQL
	return self
}

func (self *EntityDBDataRead) setQueryOriginal() *EntityDBDataRead {
	pathDirSQL := self.PathDirSQL
	fileNameQuery := conf.FileNameQuery
	pathQuery := pathDirSQL + fileNameQuery
	queryOriginal := pack.ReadFileEmbedAsString(pathQuery)
	queryTrimEnter := strings.Replace(queryOriginal, "\n", " ", -1)
	queryTrimSemicolon := strings.Replace(queryTrimEnter, ";", "", -1)
	queryFormated := strings.TrimSpace(queryTrimSemicolon)
	self.QueryWithoutArgs = queryFormated
	return self
}

func (self *EntityDBDataRead) setArrArgsForQuery() *EntityDBDataRead {
	sqlName := self.SQLName
	arrArgName := sqlInfo.GetArrArgNameBySQLName(sqlName)
	attrArgsForQuery := self.AttrArgsForQuery
	arrArgsForQuery := make([]string, 0)
	for _, argName := range arrArgName {
		argValue, ok := attrArgsForQuery[argName]
		if !ok {
			msgError := fmt.Sprintf(
				"The argForQuery |%s| is required",
				argName)
			err.ErrPanic(msgError)
		}
		arrArgsForQuery = append(arrArgsForQuery, argValue)
	}
	self.ArrArgsForQuery = arrArgsForQuery
	return self
}

func (self *EntityDBDataRead) setQueryWithArgs() *EntityDBDataRead {
	queryOriginal := self.QueryWithoutArgs
	arrArgsForQuery := self.ArrArgsForQuery
	arrArgsForQueryAny := make([]interface{}, len(arrArgsForQuery))
	for k, v := range arrArgsForQuery {
		arrArgsForQueryAny[k] = v
	}
	queryWithArgs := fmt.Sprintf(
		queryOriginal,
		arrArgsForQueryAny...)
	self.QueryWithArgs = queryWithArgs
	return self
}

func (self *EntityDBDataRead) setEntityDBData() *EntityDBDataRead {
	sqlName := self.SQLName
	entityDBData := sqlInfo.GetEntityDBDataBySQLName(sqlName)
	self.EntityDBData = entityDBData
	return self
}
