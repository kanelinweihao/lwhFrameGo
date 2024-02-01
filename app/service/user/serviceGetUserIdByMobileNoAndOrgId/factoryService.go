package serviceGetUserIdByMobileNoAndOrgId

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/base/serviceBase"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

var serviceName = "GetUserIdByMobileNoAndOrgId"

func InitEntityService() (entityService typeInterface.EntityService1) {
	entityService = new(EntityService)
	entityServiceBase := new(serviceBase.EntityServiceBase1)
	entityService.Load(entityServiceBase).Init(serviceName)
	return entityService
}

func (self *EntityService) Load(entityServiceBase typeInterface.EntityService1) typeInterface.EntityService1 {
	self.EntityService1 = entityServiceBase
	entityServiceBase.Load(self)
	return self
}
