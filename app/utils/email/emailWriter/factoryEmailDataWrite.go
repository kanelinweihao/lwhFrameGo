package emailWriter

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"gopkg.in/gomail.v2"
)

func InitEntityEmailData(emailSender gomail.SendCloser, attrT1ForEmail typeMap.AttrT1, emailFrom string) (entityEmailData *EntityEmailData) {
	entityEmailData = new(EntityEmailData)
	entityEmailData.Init(emailSender, attrT1ForEmail, emailFrom)
	return entityEmailData
}

func (self *EntityEmailData) Init(emailSender gomail.SendCloser, attrT1ForEmail typeMap.AttrT1, emailFrom string) *EntityEmailData {
	self.setPropertiesIn(emailSender, attrT1ForEmail, emailFrom).setPropertiesMore()
	return self
}

func (self *EntityEmailData) setPropertiesIn(emailSender gomail.SendCloser, attrT1ForEmail typeMap.AttrT1, emailFrom string) *EntityEmailData {
	conv.ToEntityFromAttr(attrT1ForEmail, self)
	self.EmailSender = emailSender
	self.EmailFrom = emailFrom
	return self
}

func (self *EntityEmailData) setPropertiesMore() *EntityEmailData {
	self.setEmailMessage()
	return self
}

func (self *EntityEmailData) setEmailMessage() *EntityEmailData {
	m := gomail.NewMessage()
	m.SetHeader("From", self.EmailFrom)
	m.SetHeader("To", self.ArrEmailTo...)
	if len(self.ArrEmailCc) > 0 {
		m.SetHeader("Cc", self.ArrEmailCc...)
	}
	m.SetHeader("Subject", self.EmailSubject)
	m.SetBody(self.EmailBodyType, self.EmailBodyText)
	if len(self.ArrEmailAttach) > 0 {
		for _, filename := range self.ArrEmailAttach {
			m.Attach(filename)
		}
	}
	self.EmailMessage = m
	return self
}
