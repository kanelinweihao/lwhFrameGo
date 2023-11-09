package dataGet

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/calc"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
)

func MakeEntityOfDataGet(paramsIn base.AttrT1, arrSQLName []string) (entityDataGet *EntityDataGet) {
	entityDataGet = initDataGet(paramsIn, arrSQLName)
	return entityDataGet
}

func initDataGet(paramsIn base.AttrT1, arrSQLName []string) (entityDataGet *EntityDataGet) {
	entityDataGet = new(EntityDataGet)
	entityDataGet.Init(paramsIn, arrSQLName)
	return entityDataGet
}

func (self *EntityDataGet) Init(paramsIn base.AttrT1, arrSQLName []string) *EntityDataGet {
	conv.ToEntityFromAttr(paramsIn, self)
	self.UserId = calc.Add(self.ShortUserId, "1000000")
	self.ArrSQLName = arrSQLName
	self.AttrArgsForQuery = base.AttrS1{
		"UserId": self.UserId,
	}
	return self
}
