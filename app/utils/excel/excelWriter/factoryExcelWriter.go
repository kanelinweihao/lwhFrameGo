package excelWriter

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
)

func InitEntityExcelWriter(boxToExcel typeMap.BoxData) (entityExcelWriter *EntityExcelWriter) {
	entityExcelWriter = new(EntityExcelWriter)
	entityExcelWriter.Init(boxToExcel)
	return entityExcelWriter
}

func (self *EntityExcelWriter) Init(boxToExcel typeMap.BoxData) *EntityExcelWriter {
	self.setPropertiesIn(boxToExcel).setPropertiesMore()
	return self
}

func (self *EntityExcelWriter) setPropertiesIn(boxToExcel typeMap.BoxData) *EntityExcelWriter {
	self.BoxToExcel = boxToExcel
	return self
}

func (self *EntityExcelWriter) setPropertiesMore() *EntityExcelWriter {
	self.setAttrEntityExcelDataWrite()
	return self
}

func (self *EntityExcelWriter) setAttrEntityExcelDataWrite() *EntityExcelWriter {
	boxToExcel := self.BoxToExcel
	attrEntityExcelDataWrite := make(map[string]*EntityExcelDataWrite)
	for pathFile, attrT1ForExcelToAssign := range boxToExcel {
		attrT1ForExcel := attrT1ForExcelToAssign.(typeMap.AttrT1)
		entityExcelDataWrite := InitEntityExcelDataWrite(attrT1ForExcel)
		attrEntityExcelDataWrite[pathFile] = entityExcelDataWrite
	}
	self.AttrEntityExcelDataWrite = attrEntityExcelDataWrite
	return self
}
