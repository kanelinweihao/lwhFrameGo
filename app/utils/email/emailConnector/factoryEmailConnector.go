package emailConnector

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"gopkg.in/gomail.v2"
)

func InitEntityEmailConnector() (entityEmailConnector *EntityEmailConnector) {
	entityEmailConnector = new(EntityEmailConnector)
	entityEmailConnector.Init()
	return entityEmailConnector
}

func (self *EntityEmailConnector) Init() *EntityEmailConnector {
	self.setEntityConfigEmail().setEmailDialer().setEmailSender()
	return self
}

func (self *EntityEmailConnector) setEntityConfigEmail() *EntityEmailConnector {
	ec := conf.GetEntityConfigEmail()
	self.EntityConfigEmail = ec
	return self
}

func (self *EntityEmailConnector) setEmailDialer() *EntityEmailConnector {
	ec := self.EntityConfigEmail
	host := ec.Host
	port := ec.Port
	username := ec.Username
	password := ec.Password
	d := gomail.NewDialer(
		host,
		port,
		username,
		password)
	self.EmailDialer = d
	return self
}

func (self *EntityEmailConnector) setEmailSender() *EntityEmailConnector {
	d := self.EmailDialer
	s, errDial := d.Dial()
	err.ErrCheck(errDial)
	self.EmailSender = s
	return self
}

func (self *EntityEmailConnector) CloseEmailConnector() {
	self.closeEmailSender()
	return
}

func (self *EntityEmailConnector) closeEmailSender() *EntityEmailConnector {
	if self.EmailSender == nil {
		return self
	}
	errClose := self.EmailSender.Close()
	err.ErrCheck(errClose)
	self.EmailSender = nil
	return self
}
