package params

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
)

type EntityParams struct {
	ShortUserId int    `validate:"required,min=1"`
	Sign        string `validate:"-"`
	UserId      int    `validate:"required,min=1000001"`
	MobileNo    string `validate:"omitempty,min=11,max=11"`
	OrgId       int    `validate:"omitempty,min=31"`
	MsgShow     string `validate:"-"`
}

func (self *EntityParams) SetPropertiesAppend(paramsAppend base.AttrT1) base.TEntityParams {
	conv.ToEntityFromAttr(paramsAppend, self)
	return self
}

func (self *EntityParams) ToAttr() (paramsOut base.AttrT1) {
	paramsOut = conv.ToAttrFromEntity(self)
	return paramsOut
}
