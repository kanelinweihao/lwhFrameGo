package excelWriter

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

func MakeEntityExcelWriter(boxForExcel base.BoxData) (entityExcelWriter *EntityExcelWriter) {
	entityExcelWriter = new(EntityExcelWriter)
	entityExcelWriter.Init(boxForExcel)
	return entityExcelWriter
}

func (self *EntityExcelWriter) Init(boxForExcel base.BoxData) *EntityExcelWriter {
	self.setParamsIn(boxForExcel).setParamsMore()
	return self
}

func (self *EntityExcelWriter) setParamsIn(boxForExcel base.BoxData) *EntityExcelWriter {
	self.BoxForExcel = boxForExcel
	return self
}

func (self *EntityExcelWriter) setParamsMore() *EntityExcelWriter {
	self.setAttrEntityExcelDataWrite()
	return self
}

func (self *EntityExcelWriter) setAttrEntityExcelDataWrite() *EntityExcelWriter {
	boxForExcel := self.BoxForExcel
	attrEntityExcelDataWrite := make(map[string]*EntityExcelDataWrite)
	for pathFile, attrT1ForExcelToAssign := range boxForExcel {
		attrT1ForExcel := attrT1ForExcelToAssign.(base.AttrT1)
		entityExcelDataWrite := MakeEntityExcelDataWrite(attrT1ForExcel)
		attrEntityExcelDataWrite[pathFile] = entityExcelDataWrite
	}
	self.AttrEntityExcelDataWrite = attrEntityExcelDataWrite
	return self
}
