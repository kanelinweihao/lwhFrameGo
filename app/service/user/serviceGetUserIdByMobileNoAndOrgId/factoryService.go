package serviceGetUserIdByMobileNoAndOrgId

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/base/serviceBase"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var serviceName = "GetUserIdByMobileNoAndOrgId"

func InitEntityService() (entityService typeStruct.EntityService1) {
	entityService = new(EntityService)
	entityServiceBase := new(serviceBase.EntityServiceBase1)
	entityService.Load(entityServiceBase).Init(serviceName)
	return entityService
}

func (self *EntityService) Load(entityServiceBase typeStruct.EntityService1) typeStruct.EntityService1 {
	self.EntityService1 = entityServiceBase
	entityServiceBase.Load(self)
	return self
}
