package dbQuerier

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

type EntityDBQuerier struct {
	AttrT2DBQuery     base.AttrT2
	AttrEntityDBQuery map[string]*EntityDBQuery
	ArrSQLName        []string
	AttrArgsForQuery  base.AttrS1
}
type EntityDBQuery struct {
	EntityDBData     base.EntityDBData
	QueryWithArgs    string
	SQLName          string
	PathDirSQL       string
	QueryWithoutArgs string
	ArrArgsForQuery  []string
}
