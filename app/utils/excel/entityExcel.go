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
	BoxForExcel       base.BoxData
	BoxFromExcel      base.AttrS3
	ArrPathFile       []string
}

func (self *EntityExcel) BatchSetDataToExcel(boxForExcel base.BoxData) (arrPathFile []string) {
	self.BoxForExcel = boxForExcel
	entityExcelWriter := excelWriter.MakeEntityExcelWriter(boxForExcel)
	self.EntityExcelWriter = entityExcelWriter
	arrPathFile = entityExcelWriter.BatchWriteExcel()
	self.ArrPathFile = arrPathFile
	return arrPathFile
}

func (self *EntityExcel) BatchGetDataFromExcel(arrPathFile []string) (boxFromExcel base.AttrS3) {
	entityExcelReader := excelReader.MakeEntityExcelReader(arrPathFile)
	self.EntityExcelReader = entityExcelReader
	boxFromExcel = entityExcelReader.BatchReadExcel()
	self.BoxFromExcel = boxFromExcel
	return boxFromExcel
}
