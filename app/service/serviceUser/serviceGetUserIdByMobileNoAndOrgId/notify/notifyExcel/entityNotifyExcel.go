package notifyExcel

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/excel"
)

type EntityNotifyExcel struct {
	AttrS3ExcelData    base.AttrS3
	ParamsToNotify     base.AttrT1
	AttrT3DBData       base.AttrT3
	UserId             int
	StrUserId          string
	ArrSQLName         []string
	BoxExcelTitle      base.AttrS3
	PathDirThisTime    string
	ParamsPathFile     base.AttrS1
	AttrEntityForExcel map[string]*EntityForExcel
	BoxToExcel         base.BoxData
	ArrPathFileExcel   []string
}

type EntityForExcel struct {
	SQLName          string
	PathFile         string
	SheetName        string
	AttrS2ExcelTitle base.AttrS2
	AttrS2ExcelData  base.AttrS2
}

func (self *EntityNotifyExcel) PutExcel() (attrS3ExcelData base.AttrS3, pathDirExcel string, arrPathFileExcel []string) {
	boxToExcel := self.BoxToExcel
	entityExcel := excel.MakeEntityExcel()
	arrPathFileExcel = entityExcel.BatchSetDataToExcel(boxToExcel)
	self.ArrPathFileExcel = arrPathFileExcel
	attrS3ExcelData = entityExcel.BatchGetDataFromExcel(arrPathFileExcel)
	self.AttrS3ExcelData = attrS3ExcelData
	pathDirExcel = self.PathDirThisTime
	return attrS3ExcelData, pathDirExcel, arrPathFileExcel
}
