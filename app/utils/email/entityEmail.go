package email

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/email/emailConnector"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/email/emailWriter"
)

type EntityEmail struct {
	EntityEmailConnector *emailConnector.EntityEmailConnector
	EntityEmailWriter    *emailWriter.EntityEmailWriter
	BoxToEmail           base.BoxData
}

func (self *EntityEmail) BatchSendEmail(boxToEmail base.BoxData) (arrEmailSubject []string) {
	self.BoxToEmail = boxToEmail
	entityEmailConnector := self.EntityEmailConnector
	entityEmailWriter := emailWriter.MakeEntityEmailWriter(entityEmailConnector, boxToEmail)
	self.EntityEmailWriter = entityEmailWriter
	arrEmailSubject = entityEmailWriter.BatchSendEmail()
	return arrEmailSubject
}
