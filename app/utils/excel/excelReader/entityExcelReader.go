package excelReader

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/goroutine"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/rfl"
)

type EntityExcelReader struct {
	BoxFromExcel            base.AttrS3
	ArrPathFile             []string
	AttrEntityExcelDataRead map[string]*EntityExcelDataRead
	AttrEntityChannel       base.AttrT1
}

func (self *EntityExcelReader) BatchReadExcel() (boxFromExcel base.AttrS3) {
	self.writeAttrEntityChannel().readAttrEntityChannel()
	boxFromExcel = self.BoxFromExcel
	return boxFromExcel
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
	boxFromExcel := make(base.AttrS3)
	for pathFile, entityChannelToAssign := range attrEntityChannel {
		entityChannel := entityChannelToAssign.(*goroutine.EntityChannel)
		attrS2ExcelData := readFromChannelOfReadExcelData(entityChannel)
		boxFromExcel[pathFile] = attrS2ExcelData
	}
	self.BoxFromExcel = boxFromExcel
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
