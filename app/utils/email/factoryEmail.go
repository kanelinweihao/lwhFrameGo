package email

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/email/emailConnector"
)

func InitEntityEmail() (entityEmail *EntityEmail) {
	entityEmail = new(EntityEmail)
	entityEmail.Init()
	return entityEmail
}

func (self *EntityEmail) Init() *EntityEmail {
	self.setEntityEmailConnector()
	return self
}

func (self *EntityEmail) setEntityEmailConnector() *EntityEmail {
	entityEmailConnector := emailConnector.InitEntityEmailConnector()
	self.EntityEmailConnector = entityEmailConnector
	return self
}

func (self *EntityEmail) CloseEmail() {
	entityEmailConnector := self.EntityEmailConnector
	entityEmailConnector.CloseEmailConnector()
	self.EntityEmailConnector = nil
	return
}
