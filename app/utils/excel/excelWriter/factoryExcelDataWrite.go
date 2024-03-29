package excelWriter

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/file"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/xuri/excelize/v2"
)

func InitEntityExcelDataWrite(attrT1ForExcel typeMap.AttrT1) (entityExcelDataWrite *EntityExcelDataWrite) {
	entityExcelDataWrite = new(EntityExcelDataWrite)
	entityExcelDataWrite.Init(attrT1ForExcel)
	return entityExcelDataWrite
}

func (self *EntityExcelDataWrite) Init(attrT1ForExcel typeMap.AttrT1) *EntityExcelDataWrite {
	self.setPropertiesIn(attrT1ForExcel).setPropertiesMore()
	return self
}

func (self *EntityExcelDataWrite) setPropertiesIn(attrT1ForExcel typeMap.AttrT1) *EntityExcelDataWrite {
	conv.ToEntityFromAttr(attrT1ForExcel, self)
	self.validate()
	return self
}

func (self *EntityExcelDataWrite) validate() *EntityExcelDataWrite {
	self.validatePathFile().validateTitle().validateData()
	return self
}

func (self *EntityExcelDataWrite) validatePathFile() *EntityExcelDataWrite {
	pathFile := self.PathFile
	ext := file.GetExt(pathFile)
	extExcel := conf.ExtExcel
	if ext == extExcel {
		return self
	}
	msgError := fmt.Sprintf(
		"The ext |%s| of pathFile |%s| is invalid, it shoule be |%s|",
		ext,
		pathFile,
		extExcel)
	err.ErrPanic(msgError)
	return self
}

func (self *EntityExcelDataWrite) validateTitle() *EntityExcelDataWrite {
	if len(self.AttrS2ExcelTitle) > 0 {
		return self
	}
	pathFile := self.PathFile
	msgError := fmt.Sprintf(
		"The title of |%s| is required",
		pathFile)
	err.ErrPanic(msgError)
	return self
}

func (self *EntityExcelDataWrite) validateData() *EntityExcelDataWrite {
	if len(self.AttrS2ExcelData) > 0 {
		return self
	}
	pathFile := self.PathFile
	msgError := fmt.Sprintf(
		"The data of |%s| is required",
		pathFile)
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
