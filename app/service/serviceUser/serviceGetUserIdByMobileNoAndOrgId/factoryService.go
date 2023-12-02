package serviceGetUserIdByMobileNoAndOrgId

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

func InitEntityService(paramsIn base.AttrT1) (entityService base.TEntityService) {
	serviceName := "GetUserIdByMobileNoAndOrgId"
	entityService = new(EntityService)
	entityService.Init(serviceName, paramsIn)
	return entityService
}

func (self *EntityService) Init(serviceName string, paramsIn base.AttrT1) base.TEntityService {
	self.ServiceName = serviceName
	self.ParamsIn = paramsIn
	return self
}
