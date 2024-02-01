package serviceBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

func (self *EntityServiceBase2) Load(entityServiceDerived typeInterface.EntityService2) typeInterface.EntityService2 {
	self.Derived = entityServiceDerived
	return self.Derived
}

func (self *EntityServiceBase2) Init(serviceName string) typeInterface.EntityService2 {
	self.ServiceName = serviceName
	return self.Derived
}
