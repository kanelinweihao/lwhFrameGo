package emailSend

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/file"
)

func MakeEntityEmailSend(pathDirExcel string, arrPathFileExcel []string) (entityEmailSend *EntityEmailSend) {
	entityEmailSend = new(EntityEmailSend)
	entityEmailSend.Init(pathDirExcel, arrPathFileExcel)
	return entityEmailSend
}

func (self *EntityEmailSend) Init(pathDirExcel string, arrPathFileExcel []string) *EntityEmailSend {
	self.setParamsIn(pathDirExcel, arrPathFileExcel).setParamsMore()
	return self
}

func (self *EntityEmailSend) setParamsIn(pathDirExcel string, arrPathFileExcel []string) *EntityEmailSend {
	self.PathDirExcel = pathDirExcel
	self.EmailSubject = file.GetFilename(pathDirExcel)
	self.ArrPathFileExcel = arrPathFileExcel
	self.ArrEmailAttach = arrPathFileExcel
	return self
}

func (self *EntityEmailSend) setParamsMore() *EntityEmailSend {
	self.setEmailHeader().setEmailBody().setBoxForEmail()
	return self
}

func (self *EntityEmailSend) setEmailHeader() *EntityEmailSend {
	self.ArrEmailTo = conf.ArrEmailToDefault
	self.ArrEmailCc = conf.ArrEmailCcDefault
	return self
}

func (self *EntityEmailSend) setEmailBody() *EntityEmailSend {
	self.EmailBodyType = conf.EmailBodyTypeHtml
	self.EmailBodyText = conf.EmailBodyTextDefault
	return self
}

func (self *EntityEmailSend) setBoxForEmail() *EntityEmailSend {
	attrT1ForEmail := conv.ToAttrFromEntity(self)
	boxForEmail := make(base.BoxData)
	boxForEmail[self.EmailSubject] = attrT1ForEmail
	self.BoxForEmail = boxForEmail
	return self
}
