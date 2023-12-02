package notifyEmail

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/email"
)

type EntityNotifyEmail struct {
	BoxToEmail       base.BoxData
	PathDirExcel     string
	ArrPathFileExcel []string
	ArrEmailTo       []string
	ArrEmailCc       []string
	EmailSubject     string
	EmailBodyType    string
	EmailBodyText    string
	ArrEmailAttach   []string
}

func (self *EntityNotifyEmail) SendEmail() (arrEmailSubject []string) {
	boxToEmail := self.BoxToEmail
	entityEmail := email.InitEntityEmail()
	defer entityEmail.CloseEmail()
	arrEmailSubject = entityEmail.BatchSendEmail(boxToEmail)
	return arrEmailSubject
}
