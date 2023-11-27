package params

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
)

type EntityParams struct {
	MobileNo string
	OrgId    int
	Sign     string
	UserId   int
	MsgShow  string
}

func (self *EntityParams) SetMsgShow() *EntityParams {
	msgShow := "Success"
	self.MsgShow = msgShow
	return self
}

func (self *EntityParams) GetParams() (paramsOut base.AttrT1) {
	paramsOut = conv.ToAttrFromEntity(self)
	return paramsOut
}
