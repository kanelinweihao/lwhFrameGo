package serviceBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

func (self *EntityServiceBase2) Load(entityServiceDerived typeStruct.EntityService2) typeStruct.EntityService2 {
	self.Derived = entityServiceDerived
	return self.Derived
}

func (self *EntityServiceBase2) Init(serviceName string) typeStruct.EntityService2 {
	self.ServiceName = serviceName
	return self.Derived
}
