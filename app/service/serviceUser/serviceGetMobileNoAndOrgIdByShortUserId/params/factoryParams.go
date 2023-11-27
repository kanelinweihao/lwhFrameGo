package params

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/calc"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
)

func MakeEntityParams(paramsIn base.AttrT1) (entityParams *EntityParams) {
	entityParams = new(EntityParams)
	entityParams.Init(paramsIn)
	return entityParams
}

func (self *EntityParams) Init(paramsIn base.AttrT1) *EntityParams {
	self.setPropertiesIn(paramsIn).setPropertiesMore()
	return self
}

func (self *EntityParams) setPropertiesIn(paramsIn base.AttrT1) *EntityParams {
	conv.ToEntityFromAttr(paramsIn, self)
	return self
}

func (self *EntityParams) setPropertiesMore() *EntityParams {
	self.setUserId()
	return self
}

func (self *EntityParams) setUserId() *EntityParams {
	strShortUserId := conv.ToStrFromInt(self.ShortUserId)
	strUserId := calc.Add(strShortUserId, "1000000")
	userId := conv.ToIntFromStr(strUserId)
	self.UserId = userId
	return self
}
