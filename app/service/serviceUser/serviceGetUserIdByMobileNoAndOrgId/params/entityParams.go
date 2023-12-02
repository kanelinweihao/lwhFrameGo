package params

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
)

type EntityParams struct {
	MobileNo string `validate:"required,min=11,max=11"`
	OrgId    int    `validate:"required,min=31"`
	Sign     string `validate:"-"`
	UserId   int    `validate:"omitempty,min=1000001"`
	MsgShow  string `validate:"-"`
}

func (self *EntityParams) SetPropertiesAppend(paramsAppend base.AttrT1) base.TEntityParams {
	conv.ToEntityFromAttr(paramsAppend, self)
	return self
}

func (self *EntityParams) ToAttr() (paramsOut base.AttrT1) {
	paramsOut = conv.ToAttrFromEntity(self)
	return paramsOut
}
