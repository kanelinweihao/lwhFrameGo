package notifyExcelBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/excel"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

type EntityNotifyExcelBase struct {
	Derived         typeInterface.EntityNotifyExcel
	AttrS3ExcelData typeMap.AttrS3
	ParamsToNotify     typeMap.AttrT1
	AttrT3DBData       typeMap.AttrT3
	UserId             int
	StrUserId          string
	ArrSQLName         []string
	BoxExcelTitle      typeMap.AttrS3
	PathDirThisTime    string
	ParamsPathFile     typeMap.AttrS1
	AttrEntityForExcel map[string]*EntityForExcel
	BoxToExcel         typeMap.BoxData
	ArrPathFileExcel   []string
}

type EntityForExcel struct {
	SQLName          string
	PathFile         string
	SheetName        string
	AttrS2ExcelTitle typeMap.AttrS2
	AttrS2ExcelData  typeMap.AttrS2
}

func (self *EntityNotifyExcelBase) ToAttr() (paramsOut typeMap.AttrT1) {
	paramsOut = conv.ToAttrFromEntity(self)
	return paramsOut
}

func (self *EntityNotifyExcelBase) PutExcel() (attrS3ExcelData typeMap.AttrS3, pathDirExcel string, arrPathFileExcel []string) {
	boxToExcel := self.BoxToExcel
	entityExcel := excel.InitEntityExcel()
	arrPathFileExcel = entityExcel.BatchSetDataToExcel(boxToExcel)
	self.ArrPathFileExcel = arrPathFileExcel
	attrS3ExcelData = entityExcel.BatchGetDataFromExcel(arrPathFileExcel)
	self.AttrS3ExcelData = attrS3ExcelData
	pathDirExcel = self.PathDirThisTime
	return attrS3ExcelData, pathDirExcel, arrPathFileExcel
}
