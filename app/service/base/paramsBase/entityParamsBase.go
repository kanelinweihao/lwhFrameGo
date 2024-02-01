package paramsBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

type EntityParamsBase struct {
	Derived typeInterface.EntityParams
}

func (self *EntityParamsBase) ToAttr() (paramsOut typeMap.AttrT1) {
	paramsOut = conv.ToAttrFromEntity(self.Derived)
	return paramsOut
}

func (self *EntityParamsBase) SetPropertiesAppend(paramsAppend typeMap.AttrT1) typeInterface.EntityParams {
	conv.ToEntityFromAttr(paramsAppend, self.Derived)
	return self
}
