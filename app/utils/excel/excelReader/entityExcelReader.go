package excelReader

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/goroutine"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/rfl"
)

type EntityExcelReader struct {
	AttrS3ExcelData         base.AttrS3
	ArrPathFile             []string
	AttrEntityExcelDataRead map[string]*EntityExcelDataRead
	AttrEntityChannel       base.AttrT1
}

func (self *EntityExcelReader) BatchReadExcel() (attrS3ExcelData base.AttrS3) {
	self.writeAttrEntityChannel().readAttrEntityChannel()
	attrS3ExcelData = self.AttrS3ExcelData
	return attrS3ExcelData
}

func (self *EntityExcelReader) writeAttrEntityChannel() *EntityExcelReader {
	attrEntityExcelDataRead := self.AttrEntityExcelDataRead
	attrEntityChannel := make(base.AttrT1)
	for pathFile, entityExcelDataRead := range attrEntityExcelDataRead {
		entityChannel := goroutine.MakeEntityChannel()
		attrEntityChannel[pathFile] = entityChannel
		go entityExcelDataRead.WriteToChannelOfReadExcelData(entityChannel)
	}
	self.AttrEntityChannel = attrEntityChannel
	return self
}

func (self *EntityExcelReader) readAttrEntityChannel() *EntityExcelReader {
	attrEntityChannel := self.AttrEntityChannel
	attrS3ExcelData := make(base.AttrS3)
	for pathFile, entityChannelToAssign := range attrEntityChannel {
		entityChannel := entityChannelToAssign.(*goroutine.EntityChannel)
		attrS2ExcelData := readFromChannelOfReadExcelData(entityChannel)
		attrS3ExcelData[pathFile] = attrS2ExcelData
	}
	self.AttrS3ExcelData = attrS3ExcelData
	return self
}

func readFromChannelOfReadExcelData(entityChannel *goroutine.EntityChannel) (attrS2ExcelData base.AttrS2) {
	dataOnce := entityChannel.ReadOnce()
	attrS2ExcelData, ok := dataOnce.(base.AttrS2)
	if !ok {
		rfl.ErrPanicFormat(attrS2ExcelData, "attrS2ExcelData", "base.AttrS2")
	}
	return attrS2ExcelData
}
