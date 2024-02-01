package serviceBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

func (self *EntityServiceBase1) Load(entityServiceDerived typeInterface.EntityService1) typeInterface.EntityService1 {
	self.Derived = entityServiceDerived
	return self.Derived
}

func (self *EntityServiceBase1) Init(serviceName string) typeInterface.EntityService1 {
	self.ServiceName = serviceName
	return self.Derived
}
