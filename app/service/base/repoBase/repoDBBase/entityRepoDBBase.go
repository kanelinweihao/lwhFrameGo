package repoDBBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/db"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

type EntityRepoDBBase struct {
	Derived      typeInterface.EntityRepoDB
	AttrT3DBData typeMap.AttrT3
	ParamsAppend     typeMap.AttrT1
	ParamsRepoDB     typeMap.AttrT1
	ArrSQLName       []string
	AttrArgsForQuery typeMap.AttrS1
}

func (self *EntityRepoDBBase) ToAttr() (paramsOut typeMap.AttrT1) {
	paramsOut = conv.ToAttrFromEntity(self)
	return paramsOut
}

func (self *EntityRepoDBBase) GetDBData() (attrT3DBData typeMap.AttrT3, paramsAppend typeMap.AttrT1) {
	self.setAttrT3DBData().setPropertiesNeed()
	attrT3DBData = self.AttrT3DBData
	paramsAppend = self.ParamsAppend
	return attrT3DBData, paramsAppend
}

func (self *EntityRepoDBBase) setAttrT3DBData() *EntityRepoDBBase {
	arrSqlName := self.ArrSQLName
	attrArgsForQuery := self.AttrArgsForQuery
	entityDB := db.InitEntityDB()
	defer entityDB.CloseDB()
	attrT3DBData := entityDB.BatchGetDataFromDB(arrSqlName, attrArgsForQuery)
	self.AttrT3DBData = attrT3DBData
	return self
}

func (self *EntityRepoDBBase) setPropertiesNeed() *EntityRepoDBBase {
	self.ParamsAppend = self.Derived.GetPropertiesNeed()
	return self
}

func (self *EntityRepoDBBase) GetPropertiesNeed() (paramsAppend typeMap.AttrT1) {
	paramsAppend = nil
	return paramsAppend
}
