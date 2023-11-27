package service

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

func MakeEntityService(paramsIn base.AttrT1) (entityService *EntityService) {
	entityService = new(EntityService)
	entityService.Init(paramsIn)
	return entityService
}

func (self *EntityService) Init(paramsIn base.AttrT1) *EntityService {
	self.ParamsIn = paramsIn
	return self
}
