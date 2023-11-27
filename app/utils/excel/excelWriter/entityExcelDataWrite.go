package excelWriter

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/goroutine"
	"github.com/xuri/excelize/v2"
)

var (
	sheetName    string
	cellColumn   string
	cellRow      int
	cellPosition string
	cellValue    string
)

type EntityExcelDataWrite struct {
	ExcelFile        *excelize.File
	PathFile         string
	SheetName        string
	AttrS2ExcelTitle base.AttrS2
	AttrS2ExcelData  base.AttrS2
}

func (self *EntityExcelDataWrite) WriteToChannelOfWriteExcelData(entityChannel *goroutine.EntityChannel) {
	defer self.CloseExcel()
	self.WriteExcel().SaveFile()
	pathFile := self.PathFile
	entityChannel.WriteOnce(pathFile)
	return
}

func (self *EntityExcelDataWrite) WriteExcel() *EntityExcelDataWrite {
	self.setSheetNew().batchSetCellOfTitle().batchSetCellOfData()
	return self
}

func (self *EntityExcelDataWrite) setSheetNew() *EntityExcelDataWrite {
	sheetName = self.SheetName
	f := self.ExcelFile
	indexSheet, errSheet := f.NewSheet(sheetName)
	err.ErrCheck(errSheet)
	f.SetActiveSheet(indexSheet)
	return self
}

func (self *EntityExcelDataWrite) batchSetCellOfTitle() *EntityExcelDataWrite {
	sheetName = self.SheetName
	attrS2ExcelTitle := self.AttrS2ExcelTitle
	cellRow = 1
	strCellRow := conv.ToStrFromInt(cellRow)
	for _, attrS1ExcelTitle := range attrS2ExcelTitle {
		sort, ok := attrS1ExcelTitle["sort"]
		if !ok {
			continue
		}
		cellColumn = getCellColumnBySort(sort)
		cellPosition = cellColumn + strCellRow
		title, ok := attrS1ExcelTitle["title"]
		if !ok {
			continue
		}
		cellValue = title
		self.setCellValue(cellPosition, cellValue)
	}
	return self
}

func (self *EntityExcelDataWrite) batchSetCellOfData() *EntityExcelDataWrite {
	sheetName = self.SheetName
	attrS2ExcelTitle := self.AttrS2ExcelTitle
	attrS2ExcelData := self.AttrS2ExcelData
	for sortRow, attrS1ExcelData := range attrS2ExcelData {
		cellRow = getCellRowBySortRow(sortRow)
		strCellRow := conv.ToStrFromInt(cellRow)
		for field, value := range attrS1ExcelData {
			cellColumn = getCellColumnByField(attrS2ExcelTitle, field)
			cellPosition = cellColumn + strCellRow
			cellValue = value
			self.setCellValue(cellPosition, cellValue)
		}
	}
	return self
}

func getCellRowBySortRow(sortRow string) (cellRow int) {
	numRow := conv.ToIntFromStr(sortRow)
	cellRow = numRow + 2
	return cellRow
}

func getCellColumnByField(attrS2ExcelTitle base.AttrS2, field string) (cellColumn string) {
	sort := getSortByField(attrS2ExcelTitle, field)
	cellColumn = getCellColumnBySort(sort)
	return cellColumn
}

func getSortByField(attrS2ExcelTitle base.AttrS2, field string) (sort string) {
	sort, ok := attrS2ExcelTitle[field]["sort"]
	if !ok {
		msgError := fmt.Sprintf(
			"The sort of field |%s| not found",
			field)
		err.ErrPanic(msgError)
	}
	return sort
}

func getCellColumnBySort(sort string) (cellColumn string) {
	attrCellColumn := conf.AttrCellColumn
	cellColumn, ok := attrCellColumn[sort]
	if !ok {
		msgError := fmt.Sprintf(
			"The cellColumn of sort |%s| not found",
			sort)
		err.ErrPanic(msgError)
	}
	return cellColumn
}

func (self *EntityExcelDataWrite) setCellValue(cellPosition string, value string) *EntityExcelDataWrite {
	sheetName = self.SheetName
	f := self.ExcelFile
	errSetCellValue := f.SetCellValue(sheetName, cellPosition, value)
	err.ErrCheck(errSetCellValue)
	return self
}