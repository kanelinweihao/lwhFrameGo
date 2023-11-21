package emailWriter

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/goroutine"
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

func (self *EntityEmailData) WriteToChannelOfWriteEmailData(entityChannel *goroutine.EntityChannel) {
	self.sendEmail()
	m := self.EmailMessage
	entityChannel.WriteOnce(m)
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
