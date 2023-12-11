package notifyEmailBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/file"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

func (self *EntityNotifyEmailBase) Load(entityNotifyEmailDerived typeStruct.EntityNotifyEmail) typeStruct.EntityNotifyEmail {
	self.Derived = entityNotifyEmailDerived
	return self.Derived
}

func (self *EntityNotifyEmailBase) Init(pathDirExcel string, arrPathFileExcel []string) typeStruct.EntityNotifyEmail {
	self.setPropertiesIn(pathDirExcel, arrPathFileExcel).setPropertiesMore()
	return self.Derived
}

func (self *EntityNotifyEmailBase) setPropertiesIn(pathDirExcel string, arrPathFileExcel []string) *EntityNotifyEmailBase {
	self.PathDirExcel = pathDirExcel
	self.EmailSubject = file.GetFilename(pathDirExcel)
	self.ArrPathFileExcel = arrPathFileExcel
	self.ArrEmailAttach = arrPathFileExcel
	return self
}

func (self *EntityNotifyEmailBase) setPropertiesMore() *EntityNotifyEmailBase {
	self.setEmailHeader().setEmailBody().setBoxToEmail()
	return self
}

func (self *EntityNotifyEmailBase) setEmailHeader() *EntityNotifyEmailBase {
	self.ArrEmailTo = conf.ArrEmailToDefault
	self.ArrEmailCc = conf.ArrEmailCcDefault
	return self
}

func (self *EntityNotifyEmailBase) setEmailBody() *EntityNotifyEmailBase {
	self.EmailBodyType = conf.EmailBodyTypeHtml
	self.EmailBodyText = conf.EmailBodyTextDefault
	return self
}

func (self *EntityNotifyEmailBase) setBoxToEmail() *EntityNotifyEmailBase {
	attrT1ForEmail := conv.ToAttrFromEntity(self)
	boxToEmail := make(typeMap.BoxData)
	boxToEmail[self.EmailSubject] = attrT1ForEmail
	self.BoxToEmail = boxToEmail
	return self
}
