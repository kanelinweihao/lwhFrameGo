package email

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/email/emailConnector"
)

/*
Init
*/

func MakeEntityEmail() (entityEmail *EntityEmail) {
	entityEmail = new(EntityEmail)
	entityEmail.Init()
	return entityEmail
}

func (self *EntityEmail) Init() *EntityEmail {
	self.setEntityEmailConnector()
	return self
}

func (self *EntityEmail) setEntityEmailConnector() *EntityEmail {
	entityEmailConnector := emailConnector.MakeEntityEmailConnector()
	self.EntityEmailConnector = entityEmailConnector
	return self
}

/*
Close
*/

func (self *EntityEmail) CloseEmail() {
	entityEmailConnector := self.EntityEmailConnector
	entityEmailConnector.CloseEmailConnector()
	self.EntityEmailConnector = nil
	return
}
