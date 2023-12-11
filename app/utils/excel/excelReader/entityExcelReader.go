package excelReader

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
)

type typeChanData = typeMap.AttrS2

type EntityExcelReader struct {
	AttrS3ExcelData         typeMap.AttrS3
	ArrPathFile             []string
	AttrEntityExcelDataRead map[string]*EntityExcelDataRead
	AttrChan                map[string]chan typeChanData
}

func (self *EntityExcelReader) BatchReadExcel() (attrS3ExcelData typeMap.AttrS3) {
	self.writeAttrChan().readAttrChan()
	attrS3ExcelData = self.AttrS3ExcelData
	return attrS3ExcelData
}

func (self *EntityExcelReader) writeAttrChan() *EntityExcelReader {
	attrEntityExcelDataRead := self.AttrEntityExcelDataRead
	attrChan := make(map[string]chan typeChanData, len(attrEntityExcelDataRead))
	for pathFile, entityExcelDataRead := range attrEntityExcelDataRead {
		chanBase := make(chan typeChanData, 1)
		attrChan[pathFile] = chanBase
		go entityExcelDataRead.WriteToChannelOfReadExcelData(chanBase)
	}
	self.AttrChan = attrChan
	return self
}

func (self *EntityExcelReader) readAttrChan() *EntityExcelReader {
	attrChan := self.AttrChan
	attrS3ExcelData := make(typeMap.AttrS3)
	for pathFile, chanBase := range attrChan {
		attrS2ExcelData := readFromChannelOfReadExcelData(chanBase)
		attrS3ExcelData[pathFile] = attrS2ExcelData
	}
	self.AttrS3ExcelData = attrS3ExcelData
	return self
}
