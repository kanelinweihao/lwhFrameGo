package serviceBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

func (self *EntityServiceBase1) Load(entityServiceDerived typeStruct.EntityService1) typeStruct.EntityService1 {
	self.Derived = entityServiceDerived
	return self.Derived
}

func (self *EntityServiceBase1) Init(serviceName string) typeStruct.EntityService1 {
	self.ServiceName = serviceName
	return self.Derived
}
