package emailWriter

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"gopkg.in/gomail.v2"
)

type EntityEmailData struct {
	EmailSender    gomail.SendCloser
	EmailMessage   *gomail.Message
	EmailFrom      string
	ArrEmailTo     []string
	ArrEmailCc     []string
	EmailSubject   string
	EmailBodyType  string
	EmailBodyText  string
	ArrEmailAttach []string
}

func (self *EntityEmailData) WriteToChannelOfWriteEmailData(chanWrite chan<- typeChanData) {
	self.sendEmail().writeChan(chanWrite)
	return
}

func (self *EntityEmailData) sendEmail() *EntityEmailData {
	s := self.EmailSender
	m := self.EmailMessage
	isSendEmailReal := conf.IsSendEmailReal()
	if !isSendEmailReal {
		return self
	}
	errSend := gomail.Send(s, m)
	err.ErrCheck(errSend)
	return self
}

func (self *EntityEmailData) writeChan(chanWrite chan<- typeChanData) *EntityEmailData {
	m := self.EmailMessage
	chanWrite <- m
	close(chanWrite)
	return self
}

func readFromChannelOfWriteExcelData(chanRead <-chan typeChanData) (emailSubject string) {
	m := <-chanRead
	arrEmailSubjectFromEmailMessage := m.GetHeader("Subject")
	emailSubjectEncoded := arrEmailSubjectFromEmailMessage[0]
	emailSubject = emailSubjectEncoded
	m.Reset()
	return emailSubject
}
