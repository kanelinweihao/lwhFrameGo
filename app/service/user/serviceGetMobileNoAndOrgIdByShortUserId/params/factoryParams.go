package params

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/base/paramsBase"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/calc"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

func InitEntityParams(paramsIn typeMap.AttrT1) (entityParams typeInterface.EntityParams) {
	entityParams = new(EntityParams)
	entityParamsBase := new(paramsBase.EntityParamsBase)
	entityParams.Load(entityParamsBase).Init(paramsIn)
	return entityParams
}

func (self *EntityParams) Load(entityParamsBase typeInterface.EntityParams) typeInterface.EntityParams {
	self.EntityParams = entityParamsBase
	entityParamsBase.Load(self)
	return self
}

func (self *EntityParams) SetParamsExec() typeInterface.EntityParams {
	strShortUserId := conv.ToStrFromInt(self.ShortUserId)
	strUserId := calc.Add(strShortUserId, "1000000")
	userId := conv.ToIntFromStr(strUserId)
	self.UserId = userId
	return self
}
