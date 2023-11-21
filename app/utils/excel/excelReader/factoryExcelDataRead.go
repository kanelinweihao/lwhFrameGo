package excelReader

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/file"
	"github.com/xuri/excelize/v2"
)

/*
Init
*/

func MakeEntityExcelDataRead(pathFile string) (entityExcelDataRead *EntityExcelDataRead) {
	entityExcelDataRead = new(EntityExcelDataRead)
	entityExcelDataRead.Init(pathFile)
	return entityExcelDataRead
}

func (self *EntityExcelDataRead) Init(pathFile string) *EntityExcelDataRead {
	self.setParamsIn(pathFile).setParamsMore()
	return self
}

func (self *EntityExcelDataRead) setParamsIn(pathFile string) *EntityExcelDataRead {
	self.PathFile = pathFile
	self.checkExt()
	return self
}

func (self *EntityExcelDataRead) checkExt() *EntityExcelDataRead {
	pathFile := self.PathFile
	ext := file.GetExt(pathFile)
	extExcel := conf.ExtExcel
	if ext == extExcel {
		return self
	}
	msgError := fmt.Sprintf(
		"文件[%s]格式错误,正确后缀应是[%s],当前实际是[%s]",
		pathFile,
		extExcel,
		ext)
	err.ErrPanic(msgError)
	return self
}

func (self *EntityExcelDataRead) setParamsMore() *EntityExcelDataRead {
	self.setSheetName().setExcelFile()
	return self
}

func (self *EntityExcelDataRead) setSheetName() *EntityExcelDataRead {
	if self.SheetName != "" {
		return self
	}
	sheetName := conf.SheetNameDefault
	self.SheetName = sheetName
	return self
}

func (self *EntityExcelDataRead) setExcelFile() *EntityExcelDataRead {
	pathFile := self.PathFile
	f, errExcelOpen := excelize.OpenFile(pathFile)
	err.ErrCheck(errExcelOpen)
	self.ExcelFile = f
	return self
}

/*
Close
*/

func (self *EntityExcelDataRead) CloseExcel() {
	errExcelClose := self.ExcelFile.Close()
	err.ErrCheck(errExcelClose)
	self.ExcelFile = nil
	return
}
