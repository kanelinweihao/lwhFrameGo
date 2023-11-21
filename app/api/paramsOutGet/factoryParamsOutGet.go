package paramsOutGet

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/calc"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
)

func MakeEntityParamsOutGet(paramsIn base.AttrT1) (entityParamsOutGet *EntityParamsOutGet) {
	entityParamsOutGet = new(EntityParamsOutGet)
	entityParamsOutGet.Init(paramsIn)
	return entityParamsOutGet
}

func (self *EntityParamsOutGet) Init(paramsIn base.AttrT1) *EntityParamsOutGet {
	self.setParamsIn(paramsIn).setParamsMore()
	return self
}

func (self *EntityParamsOutGet) setParamsIn(paramsIn base.AttrT1) *EntityParamsOutGet {
	conv.ToEntityFromAttr(paramsIn, self)
	self.ParamsIn = paramsIn
	return self
}

func (self *EntityParamsOutGet) setParamsMore() *EntityParamsOutGet {
	self.setUserId().setMsgOut().setParamsOut()
	return self
}

func (self *EntityParamsOutGet) setUserId() *EntityParamsOutGet {
	self.UserId = calc.Add(self.ShortUserId, "1000000")
	return self
}

func (self *EntityParamsOutGet) setMsgOut() *EntityParamsOutGet {
	msgOut := fmt.Sprintf(
		"\n%s\n%s\n",
		self.UserId,
		"Success")
	self.MsgOut = msgOut
	return self
}

func (self *EntityParamsOutGet) setParamsOut() *EntityParamsOutGet {
	paramsOut := make(base.AttrT1)
	paramsOut = base.AttrT1{
		"Sign":        self.Sign,
		"ShortUserId": self.ShortUserId,
		"UserId":      self.UserId,
		"MsgOut":      self.MsgOut,
	}
	self.ParamsOut = paramsOut
	return self
}
