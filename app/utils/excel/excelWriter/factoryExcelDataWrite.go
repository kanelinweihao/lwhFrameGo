package excelWriter

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/file"
	"github.com/xuri/excelize/v2"
)

func InitEntityExcelDataWrite(attrT1ForExcel base.AttrT1) (entityExcelDataWrite *EntityExcelDataWrite) {
	entityExcelDataWrite = new(EntityExcelDataWrite)
	entityExcelDataWrite.Init(attrT1ForExcel)
	return entityExcelDataWrite
}

func (self *EntityExcelDataWrite) Init(attrT1ForExcel base.AttrT1) *EntityExcelDataWrite {
	self.setPropertiesIn(attrT1ForExcel).setPropertiesMore()
	return self
}

func (self *EntityExcelDataWrite) setPropertiesIn(attrT1ForExcel base.AttrT1) *EntityExcelDataWrite {
	conv.ToEntityFromAttr(attrT1ForExcel, self)
	self.validateExt()
	return self
}

func (self *EntityExcelDataWrite) validateExt() *EntityExcelDataWrite {
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

func (self *EntityExcelDataWrite) setPropertiesMore() *EntityExcelDataWrite {
	self.setSheetName().setExcelFile()
	return self
}

func (self *EntityExcelDataWrite) setSheetName() *EntityExcelDataWrite {
	if self.SheetName != "" {
		return self
	}
	sheetName := conf.SheetNameDefault
	self.SheetName = sheetName
	return self
}

func (self *EntityExcelDataWrite) setExcelFile() *EntityExcelDataWrite {
	f := excelize.NewFile()
	self.ExcelFile = f
	return self
}

func (self *EntityExcelDataWrite) SaveFile() *EntityExcelDataWrite {
	pathFile := self.PathFile
	f := self.ExcelFile
	errExcelSave := f.SaveAs(pathFile)
	err.ErrCheck(errExcelSave)
	return self
}

func (self *EntityExcelDataWrite) CloseExcel() {
	errExcelClose := self.ExcelFile.Close()
	err.ErrCheck(errExcelClose)
	self.ExcelFile = nil
	return
}
