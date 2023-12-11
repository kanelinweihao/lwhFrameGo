package serviceBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

func (self *EntityServiceBase) Load(entityServiceDerived typeStruct.EntityService) typeStruct.EntityService {
	self.Derived = entityServiceDerived
	return self.Derived
}

func (self *EntityServiceBase) Init(serviceName string) typeStruct.EntityService {
	self.ServiceName = serviceName
	return self.Derived
}
