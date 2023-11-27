package notifyEmail

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/file"
)

func MakeEntityNotifyEmail(pathDirExcel string, arrPathFileExcel []string) (entityNotifyEmail *EntityNotifyEmail) {
	entityNotifyEmail = new(EntityNotifyEmail)
	entityNotifyEmail.Init(pathDirExcel, arrPathFileExcel)
	return entityNotifyEmail
}

func (self *EntityNotifyEmail) Init(pathDirExcel string, arrPathFileExcel []string) *EntityNotifyEmail {
	self.setPropertiesIn(pathDirExcel, arrPathFileExcel).setPropertiesMore()
	return self
}

func (self *EntityNotifyEmail) setPropertiesIn(pathDirExcel string, arrPathFileExcel []string) *EntityNotifyEmail {
	self.PathDirExcel = pathDirExcel
	self.EmailSubject = file.GetFilename(pathDirExcel)
	self.ArrPathFileExcel = arrPathFileExcel
	self.ArrEmailAttach = arrPathFileExcel
	return self
}

func (self *EntityNotifyEmail) setPropertiesMore() *EntityNotifyEmail {
	self.setEmailHeader().setEmailBody().setBoxToEmail()
	return self
}

func (self *EntityNotifyEmail) setEmailHeader() *EntityNotifyEmail {
	self.ArrEmailTo = conf.ArrEmailToDefault
	self.ArrEmailCc = conf.ArrEmailCcDefault
	return self
}

func (self *EntityNotifyEmail) setEmailBody() *EntityNotifyEmail {
	self.EmailBodyType = conf.EmailBodyTypeHtml
	self.EmailBodyText = conf.EmailBodyTextDefault
	return self
}

func (self *EntityNotifyEmail) setBoxToEmail() *EntityNotifyEmail {
	attrT1ForEmail := conv.ToAttrFromEntity(self)
	boxToEmail := make(base.BoxData)
	boxToEmail[self.EmailSubject] = attrT1ForEmail
	self.BoxToEmail = boxToEmail
	return self
}
