package excelWriter

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

type typeChanData = string

type EntityExcelWriter struct {
	ArrPathFile              []string
	BoxToExcel               base.BoxData
	AttrEntityExcelDataWrite map[string]*EntityExcelDataWrite
	AttrChan                 map[string]chan typeChanData
}

func (self *EntityExcelWriter) BatchWriteExcel() (arrPathFile []string) {
	self.writeAttrChan().readAttrChan()
	arrPathFile = self.ArrPathFile
	return arrPathFile
}

func (self *EntityExcelWriter) writeAttrChan() *EntityExcelWriter {
	attrEntityExcelDataWrite := self.AttrEntityExcelDataWrite
	attrChan := make(map[string]chan typeChanData, len(attrEntityExcelDataWrite))
	for pathFile, entityExcelDataWrite := range attrEntityExcelDataWrite {
		chanBase := make(chan typeChanData, 1)
		attrChan[pathFile] = chanBase
		go entityExcelDataWrite.WriteToChannelOfWriteExcelData(chanBase)
	}
	self.AttrChan = attrChan
	return self
}

func (self *EntityExcelWriter) readAttrChan() *EntityExcelWriter {
	attrChan := self.AttrChan
	arrPathFile := make([]string, 0, len(attrChan))
	for _, chanBase := range attrChan {
		pathFile := readFromChannelOfWriteExcelData(chanBase)
		arrPathFile = append(arrPathFile, pathFile)
	}
	self.ArrPathFile = arrPathFile
	return self
}
