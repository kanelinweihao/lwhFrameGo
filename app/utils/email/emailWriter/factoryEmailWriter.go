package emailWriter

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/email/emailConnector"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
)

func InitEntityEmailWriter(entityEmailConnector *emailConnector.EntityEmailConnector, boxToEmail typeMap.BoxData) (entityEmailWriter *EntityEmailWriter) {
	entityEmailWriter = new(EntityEmailWriter)
	entityEmailWriter.Init(entityEmailConnector, boxToEmail)
	return entityEmailWriter
}

func (self *EntityEmailWriter) Init(entityEmailConnector *emailConnector.EntityEmailConnector, boxToEmail typeMap.BoxData) *EntityEmailWriter {
	self.setPropertiesIn(entityEmailConnector, boxToEmail).setPropertiesMore()
	return self
}

func (self *EntityEmailWriter) setPropertiesIn(entityEmailConnector *emailConnector.EntityEmailConnector, boxToEmail typeMap.BoxData) *EntityEmailWriter {
	self.EntityEmailConnector = entityEmailConnector
	self.BoxToEmail = boxToEmail
	return self
}

func (self *EntityEmailWriter) setPropertiesMore() *EntityEmailWriter {
	self.setEmailFrom().setAttrEntityEmailData()
	return self
}

func (self *EntityEmailWriter) setEmailFrom() *EntityEmailWriter {
	emailFrom := conf.GetEmailFrom()
	self.EmailFrom = emailFrom
	return self
}

func (self *EntityEmailWriter) setAttrEntityEmailData() *EntityEmailWriter {
	emailSender := self.EntityEmailConnector.EmailSender
	boxToEmail := self.BoxToEmail
	emailFrom := self.EmailFrom
	attrEntityEmailData := make(map[string]*EntityEmailData)
	for emailSubject, attrT1ForEmailToAssign := range boxToEmail {
		attrT1ForEmail := attrT1ForEmailToAssign.(typeMap.AttrT1)
		entityEmailData := InitEntityEmailData(emailSender, attrT1ForEmail, emailFrom)
		attrEntityEmailData[emailSubject] = entityEmailData
	}
	self.AttrEntityEmailData = attrEntityEmailData
	return self
}
