package paramsBase

import (
	"github.com/go-playground/validator/v10"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

func (self *EntityParamsBase) Load(entityParamsDerived typeInterface.EntityParams) typeInterface.EntityParams {
	self.Derived = entityParamsDerived
	return self.Derived
}

func (self *EntityParamsBase) Init(paramsIn typeMap.AttrT1) typeInterface.EntityParams {
	self.setPropertiesIn(paramsIn).setPropertiesMore()
	self.Derived.Validate()
	return self.Derived
}

func (self *EntityParamsBase) setPropertiesIn(paramsIn typeMap.AttrT1) *EntityParamsBase {
	conv.ToEntityFromAttr(paramsIn, self.Derived)
	return self
}

func (self *EntityParamsBase) setPropertiesMore() *EntityParamsBase {
	self.Derived.SetParamsExec()
	return self
}

func (self *EntityParamsBase) SetParamsExec() typeInterface.EntityParams {
	return self.Derived
}

func (self *EntityParamsBase) Validate() typeInterface.EntityParams {
	validate := validator.New()
	errValidate := validate.Struct(self.Derived)
	err.ErrValidate(errValidate)
	return self.Derived
}
