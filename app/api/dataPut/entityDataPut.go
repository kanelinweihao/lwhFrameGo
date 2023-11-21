package dataPut

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/excel"
)

type EntityDataPut struct {
	BoxForExcel        base.BoxData
	ParamsOut          base.AttrT1
	BoxExcelData       base.AttrS3
	UserId             string
	ArrSQLName         []string
	BoxExcelTitle      base.AttrS3
	PathDirThisTime    string
	ParamsPathFile     base.AttrS1
	AttrEntityForExcel map[string]*EntityForExcel
	ArrPathFile        []string
	BoxFromExcel       base.AttrS3
}

type EntityForExcel struct {
	SQLName          string
	PathFile         string
	SheetName        string
	AttrS2ExcelTitle base.AttrS2
	AttrS2ExcelData  base.AttrS2
}

func (self *EntityDataPut) PutData() (boxFromExcel base.AttrS3) {
	boxForExcel := self.BoxForExcel
	entityExcel := excel.MakeEntityExcel()
	arrPathFile := entityExcel.BatchSetDataToExcel(boxForExcel)
	self.ArrPathFile = arrPathFile
	boxFromExcel = entityExcel.BatchGetDataFromExcel(arrPathFile)
	self.BoxFromExcel = boxFromExcel
	return boxFromExcel
}
