package excel

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/excel/excelReader"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/excel/excelWriter"
)

var ExtExcel string = conf.ExtExcel
var SheetNameDefault string = conf.SheetNameDefault

type EntityExcel struct {
	EntityExcelWriter *excelWriter.EntityExcelWriter
	EntityExcelReader *excelReader.EntityExcelReader
	BoxToExcel        base.BoxData
	ArrPathFile       []string
	AttrS3ExcelData   base.AttrS3
}

func (self *EntityExcel) BatchSetDataToExcel(boxToExcel base.BoxData) (arrPathFile []string) {
	self.BoxToExcel = boxToExcel
	entityExcelWriter := excelWriter.InitEntityExcelWriter(boxToExcel)
	self.EntityExcelWriter = entityExcelWriter
	arrPathFile = entityExcelWriter.BatchWriteExcel()
	self.ArrPathFile = arrPathFile
	return arrPathFile
}

func (self *EntityExcel) BatchGetDataFromExcel(arrPathFile []string) (attrS3ExcelData base.AttrS3) {
	entityExcelReader := excelReader.InitEntityExcelReader(arrPathFile)
	self.EntityExcelReader = entityExcelReader
	attrS3ExcelData = entityExcelReader.BatchReadExcel()
	self.AttrS3ExcelData = attrS3ExcelData
	return attrS3ExcelData
}
