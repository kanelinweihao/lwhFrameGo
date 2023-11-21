package emailSend

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/email"
)

type EntityEmailSend struct {
	BoxForEmail      base.BoxData
	PathDirExcel     string
	ArrPathFileExcel []string
	ArrEmailTo       []string
	ArrEmailCc       []string
	EmailSubject     string
	EmailBodyType    string
	EmailBodyText    string
	ArrEmailAttach   []string
}

func (self *EntityEmailSend) SendEmail() (arrEmailSubject []string) {
	boxForEmail := self.BoxForEmail
	entityEmail := email.MakeEntityEmail()
	defer entityEmail.CloseEmail()
	arrEmailSubject = entityEmail.BatchSendEmail(boxForEmail)
	return arrEmailSubject
}
