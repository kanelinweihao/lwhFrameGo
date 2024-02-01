package notifyEmailBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/email"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

type EntityNotifyEmailBase struct {
	Derived    typeInterface.EntityNotifyEmail
	BoxToEmail typeMap.BoxData
	PathDirExcel     string
	ArrPathFileExcel []string
	ArrEmailTo       []string
	ArrEmailCc       []string
	EmailSubject     string
	EmailBodyType    string
	EmailBodyText    string
	ArrEmailAttach   []string
}

func (self *EntityNotifyEmailBase) ToAttr() (paramsOut typeMap.AttrT1) {
	paramsOut = conv.ToAttrFromEntity(self)
	return paramsOut
}

func (self *EntityNotifyEmailBase) SendEmail() (arrEmailSubject []string) {
	boxToEmail := self.BoxToEmail
	entityEmail := email.InitEntityEmail()
	defer entityEmail.CloseEmail()
	arrEmailSubject = entityEmail.BatchSendEmail(boxToEmail)
	return arrEmailSubject
}
