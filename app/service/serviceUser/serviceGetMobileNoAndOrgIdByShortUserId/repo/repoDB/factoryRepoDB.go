package repoDB

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
)

var arrSQLName = []string{
	"GetMobileNoByUserId",
	"GetOrgIdByUserId",
}

func InitEntityRepoDB(paramsRepoDB base.AttrT1) (entityRepoDB *EntityRepoDB) {
	entityRepoDB = initRepoDB(paramsRepoDB)
	return entityRepoDB
}

func initRepoDB(paramsRepoDB base.AttrT1) (entityRepoDB *EntityRepoDB) {
	entityRepoDB = new(EntityRepoDB)
	entityRepoDB.Init(paramsRepoDB)
	return entityRepoDB
}

func (self *EntityRepoDB) Init(paramsRepoDB base.AttrT1) *EntityRepoDB {
	self.setPropertiesIn(paramsRepoDB).setPropertiesMore()
	return self
}

func (self *EntityRepoDB) setPropertiesIn(paramsRepoDB base.AttrT1) *EntityRepoDB {
	self.ParamsRepoDB = paramsRepoDB
	return self
}

func (self *EntityRepoDB) setPropertiesMore() *EntityRepoDB {
	self.setArrSQLName().setAttrArgsForQuery()
	return self
}

func (self *EntityRepoDB) setArrSQLName() *EntityRepoDB {
	self.ArrSQLName = arrSQLName
	return self
}

func (self *EntityRepoDB) setAttrArgsForQuery() *EntityRepoDB {
	paramsRepoDB := self.ParamsRepoDB
	userId := paramsRepoDB["UserId"].(int)
	strUserId := conv.ToStrFromInt(userId)
	attrArgsForQuery := base.AttrS1{
		"UserId": strUserId,
	}
	self.AttrArgsForQuery = attrArgsForQuery
	return self
}
