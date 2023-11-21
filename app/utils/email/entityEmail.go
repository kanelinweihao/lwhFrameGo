package email

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/email/emailConnector"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/email/emailWriter"
)

type EntityEmail struct {
	EntityEmailConnector *emailConnector.EntityEmailConnector
	EntityEmailWriter    *emailWriter.EntityEmailWriter
	BoxForEmail          base.BoxData
}

func (self *EntityEmail) BatchSendEmail(boxForEmail base.BoxData) (arrEmailSubject []string) {
	self.BoxForEmail = boxForEmail
	entityEmailConnector := self.EntityEmailConnector
	entityEmailWriter := emailWriter.MakeEntityEmailWriter(entityEmailConnector, boxForEmail)
	self.EntityEmailWriter = entityEmailWriter
	arrEmailSubject = entityEmailWriter.BatchSendEmail()
	return arrEmailSubject
}
