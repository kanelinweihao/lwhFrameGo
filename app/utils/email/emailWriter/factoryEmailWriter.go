package emailWriter

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/email/emailConnector"
)

func MakeEntityEmailWriter(entityEmailConnector *emailConnector.EntityEmailConnector, boxForEmail base.BoxData) (entityEmailWriter *EntityEmailWriter) {
	entityEmailWriter = new(EntityEmailWriter)
	entityEmailWriter.Init(entityEmailConnector, boxForEmail)
	return entityEmailWriter
}

func (self *EntityEmailWriter) Init(entityEmailConnector *emailConnector.EntityEmailConnector, boxForEmail base.BoxData) *EntityEmailWriter {
	self.setParamsIn(entityEmailConnector, boxForEmail).setParamsMore()
	return self
}

func (self *EntityEmailWriter) setParamsIn(entityEmailConnector *emailConnector.EntityEmailConnector, boxForEmail base.BoxData) *EntityEmailWriter {
	self.EntityEmailConnector = entityEmailConnector
	self.BoxForEmail = boxForEmail
	return self
}

func (self *EntityEmailWriter) setParamsMore() *EntityEmailWriter {
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
	boxForEmail := self.BoxForEmail
	emailFrom := self.EmailFrom
	attrEntityEmailData := make(map[string]*EntityEmailData)
	for emailSubject, attrT1ForEmailToAssign := range boxForEmail {
		attrT1ForEmail := attrT1ForEmailToAssign.(base.AttrT1)
		entityEmailData := MakeEntityEmailData(emailSender, attrT1ForEmail, emailFrom)
		attrEntityEmailData[emailSubject] = entityEmailData
	}
	self.AttrEntityEmailData = attrEntityEmailData
	return self
}
