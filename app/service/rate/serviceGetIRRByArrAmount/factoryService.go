package serviceGetIRRByArrAmount

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/base/serviceBase"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

var serviceName = "GetIRRByArrAmount"

func InitEntityService() (entityService typeInterface.EntityService2) {
	entityService = new(EntityService)
	entityServiceBase := new(serviceBase.EntityServiceBase2)
	entityService.Load(entityServiceBase).Init(serviceName)
	return entityService
}

func (self *EntityService) Load(entityServiceBase typeInterface.EntityService2) typeInterface.EntityService2 {
	self.EntityService2 = entityServiceBase
	entityServiceBase.Load(self)
	return self
}
