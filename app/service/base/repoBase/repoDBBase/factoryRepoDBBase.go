package repoDBBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

func (self *EntityRepoDBBase) Load(entityRepoDBDerived typeInterface.EntityRepoDB) typeInterface.EntityRepoDB {
	self.Derived = entityRepoDBDerived
	return self.Derived
}

func (self *EntityRepoDBBase) Init(paramsRepoDB typeMap.AttrT1) typeInterface.EntityRepoDB {
	self.setPropertiesIn(paramsRepoDB).setPropertiesMore()
	return self.Derived
}

func (self *EntityRepoDBBase) setPropertiesIn(paramsRepoDB typeMap.AttrT1) *EntityRepoDBBase {
	self.ParamsRepoDB = paramsRepoDB
	return self
}

func (self *EntityRepoDBBase) setPropertiesMore() *EntityRepoDBBase {
	self.setArrSQLName().setAttrArgsForQuery()
	return self
}

func (self *EntityRepoDBBase) setArrSQLName() *EntityRepoDBBase {
	arrSQLName := self.Derived.GetArrSQLName()
	self.ArrSQLName = arrSQLName
	return self
}

func (self *EntityRepoDBBase) GetArrSQLName() (arrSQLName []string) {
	arrSQLName = nil
	return arrSQLName
}

func (self *EntityRepoDBBase) setAttrArgsForQuery() *EntityRepoDBBase {
	attrArgsForQuery := self.Derived.GetAttrArgsForQuery()
	self.AttrArgsForQuery = attrArgsForQuery
	return self
}

func (self *EntityRepoDBBase) GetAttrArgsForQuery() (attrArgsForQuery typeMap.AttrS1) {
	attrArgsForQuery = nil
	return attrArgsForQuery
}
