package params

import (
	"github.com/go-playground/validator/v10"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
)

func InitEntityParams(paramsIn base.AttrT1) (entityParams base.TEntityParams) {
	entityParams = new(EntityParams)
	entityParams.Init(paramsIn)
	entityParams.Validate()
	return entityParams
}

func (self *EntityParams) Init(paramsIn base.AttrT1) base.TEntityParams {
	self.setPropertiesIn(paramsIn).setPropertiesMore()
	return self
}

func (self *EntityParams) setPropertiesIn(paramsIn base.AttrT1) *EntityParams {
	conv.ToEntityFromAttr(paramsIn, self)
	return self
}

func (self *EntityParams) setPropertiesMore() *EntityParams {
	self.setMsgShow()
	return self
}

func (self *EntityParams) setMsgShow() *EntityParams {
	msgShow := "Success"
	self.MsgShow = msgShow
	return self
}

func (self *EntityParams) Validate() base.TEntityParams {
	validate := validator.New()
	errValidate := validate.Struct(self)
	err.ErrValidate(errValidate)
	return self
}
