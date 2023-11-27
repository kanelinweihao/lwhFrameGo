package emailWriter

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/email/emailConnector"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/goroutine"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/rfl"
	"gopkg.in/gomail.v2"
)

type EntityEmailWriter struct {
	ArrEmailSubject      []string
	EmailFrom            string
	BoxToEmail           base.BoxData
	EntityEmailConnector *emailConnector.EntityEmailConnector
	AttrEntityEmailData  map[string]*EntityEmailData
	AttrEntityChannel    base.AttrT1
}

func (self *EntityEmailWriter) BatchSendEmail() (arrEmailSubject []string) {
	self.writeAttrEntityChannel().readAttrEntityChannel()
	arrEmailSubject = self.ArrEmailSubject
	return arrEmailSubject
}

func (self *EntityEmailWriter) writeAttrEntityChannel() *EntityEmailWriter {
	attrEntityEmailData := self.AttrEntityEmailData
	attrEntityChannel := make(base.AttrT1)
	for emailSubject, entityEmailData := range attrEntityEmailData {
		entityChannel := goroutine.MakeEntityChannel()
		attrEntityChannel[emailSubject] = entityChannel
		go entityEmailData.WriteToChannelOfWriteEmailData(entityChannel)
	}
	self.AttrEntityChannel = attrEntityChannel
	return self
}

func (self *EntityEmailWriter) readAttrEntityChannel() *EntityEmailWriter {
	attrEntityChannel := self.AttrEntityChannel
	arrEmailSubject := make([]string, 0, len(attrEntityChannel))
	for _, entityChannelToAssign := range attrEntityChannel {
		entityChannel := entityChannelToAssign.(*goroutine.EntityChannel)
		emailSubject := readFromChannelOfWriteExcelData(entityChannel)
		arrEmailSubject = append(arrEmailSubject, emailSubject)
	}
	self.ArrEmailSubject = arrEmailSubject
	return self
}

func readFromChannelOfWriteExcelData(entityChannel *goroutine.EntityChannel) (emailSubject string) {
	dataOnce := entityChannel.ReadOnce()
	m, ok := dataOnce.(*gomail.Message)
	if !ok {
		rfl.ErrPanicFormat(m, "emailMessage", "*gomail.Message")
	}
	arrEmailSubjectFromEmailMessage := m.GetHeader("Subject")
	emailSubjectEncoded := arrEmailSubjectFromEmailMessage[0]
	emailSubject = emailSubjectEncoded
	m.Reset()
	return emailSubject
}
