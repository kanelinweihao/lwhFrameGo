package serviceGetUserIdByMobileNoAndOrgId

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/base/serviceBase"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var serviceName = "GetUserIdByMobileNoAndOrgId"

func InitEntityService() (entityService typeStruct.EntityService) {
	entityService = new(EntityService)
	entityServiceBase := new(serviceBase.EntityServiceBase)
	entityService.Load(entityServiceBase).Init(serviceName)
	return entityService
}

func (self *EntityService) Load(entityServiceBase typeStruct.EntityService) typeStruct.EntityService {
	self.EntityService = entityServiceBase
	entityServiceBase.Load(self)
	return self
}
