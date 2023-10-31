package dataGet

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/calc"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
)

func MakeEntityOfDataGet(paramsIn base.AttrT1) (entityDataGet *EntityDataGet) {
	entityDataGet = initDataGet(paramsIn)
	return entityDataGet
}

func initDataGet(paramsIn base.AttrT1) (entityDataGet *EntityDataGet) {
	entityDataGet = new(EntityDataGet)
	entityDataGet.Init(paramsIn)
	return entityDataGet
}

func (self *EntityDataGet) Init(paramsIn base.AttrT1) *EntityDataGet {
	conv.ToEntityFromAttr(paramsIn, self)
	self.UID = calc.Add(self.Field2, "1000000")
	self.Field3 = "empty"
	return self
}
