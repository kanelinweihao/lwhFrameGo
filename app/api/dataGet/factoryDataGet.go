package dataGet

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

func MakeEntityDataGet(paramsOut base.AttrT1) (entityDataGet *EntityDataGet) {
	entityDataGet = initDataGet(paramsOut)
	return entityDataGet
}

func initDataGet(paramsOut base.AttrT1) (entityDataGet *EntityDataGet) {
	entityDataGet = new(EntityDataGet)
	entityDataGet.Init(paramsOut)
	return entityDataGet
}

func (self *EntityDataGet) Init(paramsOut base.AttrT1) *EntityDataGet {
	self.setParamsIn(paramsOut).setParamsMore()
	return self
}

func (self *EntityDataGet) setParamsIn(paramsOut base.AttrT1) *EntityDataGet {
	self.ParamsOut = paramsOut
	return self
}

func (self *EntityDataGet) setParamsMore() *EntityDataGet {
	self.setArrSQLName().setAttrArgsForQuery().setBoxForDB()
	return self
}

func (self *EntityDataGet) setArrSQLName() *EntityDataGet {
	self.ArrSQLName = arrSQLName
	return self
}

func (self *EntityDataGet) setAttrArgsForQuery() *EntityDataGet {
	paramsOut := self.ParamsOut
	userId := paramsOut["UserId"].(string)
	attrArgsForQuery := base.AttrS1{
		"UserId": userId,
	}
	self.AttrArgsForQuery = attrArgsForQuery
	return self
}

func (self *EntityDataGet) setBoxForDB() *EntityDataGet {
	boxForDB := make(base.BoxData)
	boxForDB["ArrSQLName"] = self.ArrSQLName
	boxForDB["AttrArgsForQuery"] = self.AttrArgsForQuery
	self.BoxForDB = boxForDB
	return self
}
