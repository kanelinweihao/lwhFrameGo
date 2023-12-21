package params

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/base/paramsBase"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

func InitEntityParams(paramsIn typeMap.AttrT1) (entityParams typeStruct.EntityParams) {
	entityParams = new(EntityParams)
	entityParamsBase := new(paramsBase.EntityParamsBase)
	entityParams.Load(entityParamsBase).Init(paramsIn)
	return entityParams
}

func (self *EntityParams) Load(entityParamsBase typeStruct.EntityParams) typeStruct.EntityParams {
	self.EntityParams = entityParamsBase
	entityParamsBase.Load(self)
	return self
}
