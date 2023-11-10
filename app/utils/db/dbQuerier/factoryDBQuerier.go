package dbQuerier

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/sqlInfo"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/pack"
	"strings"
)

var fileNameQuery string = "query.sql"

/*
Init
*/

func MakeEntityOfDBQuerier(arrSQLName []string, attrArgsForQuery base.AttrS1) (entityDBQuerier *EntityDBQuerier) {
	entityDBQuerier = initEntityDBQuerier(arrSQLName, attrArgsForQuery)
	return entityDBQuerier
}

func initEntityDBQuerier(arrSQLName []string, attrArgsForQuery base.AttrS1) (entityDBQuerier *EntityDBQuerier) {
	entityDBQuerier = new(EntityDBQuerier)
	entityDBQuerier.Init(arrSQLName, attrArgsForQuery)
	return entityDBQuerier
}

func (self *EntityDBQuerier) Init(arrSQLName []string, attrArgsForQuery base.AttrS1) *EntityDBQuerier {
	self.setArrSQLName(arrSQLName).setAttrArgsForQuery(attrArgsForQuery).setAttrEntityDBQuerier()
	return self
}

func (self *EntityDBQuerier) setArrSQLName(arrSQLName []string) *EntityDBQuerier {
	self.ArrSQLName = arrSQLName
	return self
}

func (self *EntityDBQuerier) setAttrArgsForQuery(attrArgsForQuery base.AttrS1) *EntityDBQuerier {
	self.AttrArgsForQuery = attrArgsForQuery
	return self
}

func (self *EntityDBQuerier) setAttrEntityDBQuerier() *EntityDBQuerier {
	arrSQLName := self.ArrSQLName
	attrArgsForQuery := self.AttrArgsForQuery
	attrEntityDBQuerier := make(map[string]*EntityDBQuery)
	attrT2DBQuerier := make(base.AttrT2)
	for _, sqlName := range arrSQLName {
		entityDBQuery := makeEntityOfDBQuery(sqlName, attrArgsForQuery)
		attrT1DBQuery := make(base.AttrT1)
		attrT1DBQuery["entityDBData"] = entityDBQuery.EntityDBData
		attrT1DBQuery["queryWithArgs"] = entityDBQuery.QueryWithArgs
		attrEntityDBQuerier[sqlName] = entityDBQuery
		attrT2DBQuerier[sqlName] = attrT1DBQuery
	}
	self.AttrEntityDBQuery = attrEntityDBQuerier
	self.AttrT2DBQuery = attrT2DBQuerier
	return self
}

/*
InitOne
*/

func makeEntityOfDBQuery(sqlName string, attrArgsForQuery base.AttrS1) (entityDBQuery *EntityDBQuery) {
	entityDBQuery = initEntityDBQuery(sqlName, attrArgsForQuery)
	return entityDBQuery
}

func initEntityDBQuery(sqlName string, attrArgsForQuery base.AttrS1) (entityDBQuery *EntityDBQuery) {
	entityDBQuery = new(EntityDBQuery)
	entityDBQuery.setSQLName(sqlName).setPathDirSQL().setQueryOriginal().setArrArgsForQuery(attrArgsForQuery).setQueryWithArgs().setEntityDBData()
	return entityDBQuery
}

func (self *EntityDBQuery) setSQLName(sqlName string) *EntityDBQuery {
	self.SQLName = sqlName
	return self
}

func (self *EntityDBQuery) setPathDirSQL() *EntityDBQuery {
	sqlName := self.SQLName
	pathDirSQL := conf.GetPathDirSQL(sqlName)
	self.PathDirSQL = pathDirSQL
	return self
}

func (self *EntityDBQuery) setQueryOriginal() *EntityDBQuery {
	pathDirSQL := self.PathDirSQL
	pathQuery := pathDirSQL + fileNameQuery
	queryOriginal := pack.ReadFileEmbedAsString(pathQuery)
	queryTrimEnter := strings.Replace(queryOriginal, "\n", " ", -1)
	queryTrimSemicolon := strings.Replace(queryTrimEnter, ";", "", -1)
	queryFormated := strings.TrimSpace(queryTrimSemicolon)
	self.QueryWithoutArgs = queryFormated
	return self
}

func (self *EntityDBQuery) setArrArgsForQuery(attrArgsForQuery base.AttrS1) *EntityDBQuery {
	sqlName := self.SQLName
	arrArgName := sqlInfo.GetArrArgNameBySQLName(sqlName)
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

func (self *EntityDBQuery) setQueryWithArgs() *EntityDBQuery {
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

func (self *EntityDBQuery) setEntityDBData() *EntityDBQuery {
	sqlName := self.SQLName
	entityDBData := sqlInfo.GetEntityDBDataBySQLName(sqlName)
	self.EntityDBData = entityDBData
	return self
}
