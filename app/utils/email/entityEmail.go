package email

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/email/emailConnector"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/email/emailWriter"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
)

type EntityEmail struct {
	EntityEmailConnector *emailConnector.EntityEmailConnector
	EntityEmailWriter    *emailWriter.EntityEmailWriter
	BoxToEmail           typeMap.BoxData
}

func (self *EntityEmail) BatchSendEmail(boxToEmail typeMap.BoxData) (arrEmailSubject []string) {
	self.BoxToEmail = boxToEmail
	entityEmailConnector := self.EntityEmailConnector
	entityEmailWriter := emailWriter.InitEntityEmailWriter(entityEmailConnector, boxToEmail)
	self.EntityEmailWriter = entityEmailWriter
	arrEmailSubject = entityEmailWriter.BatchSendEmail()
	return arrEmailSubject
}
