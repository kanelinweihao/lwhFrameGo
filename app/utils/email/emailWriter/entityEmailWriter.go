package emailWriter

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/email/emailConnector"
	"gopkg.in/gomail.v2"
)

type typeChanData = *gomail.Message

type EntityEmailWriter struct {
	ArrEmailSubject      []string
	EmailFrom            string
	BoxToEmail           base.BoxData
	EntityEmailConnector *emailConnector.EntityEmailConnector
	AttrEntityEmailData  map[string]*EntityEmailData
	AttrChan             map[string]chan typeChanData
}

func (self *EntityEmailWriter) BatchSendEmail() (arrEmailSubject []string) {
	self.writeAttrChan().readAttrChan()
	arrEmailSubject = self.ArrEmailSubject
	return arrEmailSubject
}

func (self *EntityEmailWriter) writeAttrChan() *EntityEmailWriter {
	attrEntityEmailData := self.AttrEntityEmailData
	attrChan := make(map[string]chan typeChanData, len(attrEntityEmailData))
	for emailSubject, entityEmailData := range attrEntityEmailData {
		chanBase := make(chan typeChanData, 1)
		attrChan[emailSubject] = chanBase
		go entityEmailData.WriteToChannelOfWriteEmailData(chanBase)
	}
	self.AttrChan = attrChan
	return self
}

func (self *EntityEmailWriter) readAttrChan() *EntityEmailWriter {
	attrChan := self.AttrChan
	arrEmailSubject := make([]string, 0, len(attrChan))
	for _, chanBase := range attrChan {
		emailSubject := readFromChannelOfWriteExcelData(chanBase)
		arrEmailSubject = append(arrEmailSubject, emailSubject)
	}
	self.ArrEmailSubject = arrEmailSubject
	return self
}
