package serviceGetRateExactlyHalf

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/base/serviceBase"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var serviceName = "GetRateExactlyHalf"

func InitEntityService() (entityService typeStruct.EntityService2) {
	entityService = new(EntityService)
	entityServiceBase := new(serviceBase.EntityServiceBase2)
	entityService.Load(entityServiceBase).Init(serviceName)
	return entityService
}

func (self *EntityService) Load(entityServiceBase typeStruct.EntityService2) typeStruct.EntityService2 {
	self.EntityService2 = entityServiceBase
	entityServiceBase.Load(self)
	return self
}
