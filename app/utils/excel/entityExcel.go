package excel

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/file"
	"github.com/xuri/excelize/v2"
)

var extExcel string = "xlsx"
var sheetName string
var cellColumn string
var cellRow int
var cellPosition string
var cellValue string
var attrCellColumn = base.AttrS1{
	"1":  "A",
	"2":  "B",
	"3":  "C",
	"4":  "D",
	"5":  "E",
	"6":  "F",
	"7":  "G",
	"8":  "H",
	"9":  "I",
	"10": "J",
	"11": "K",
	"12": "L",
	"13": "M",
	"14": "N",
	"15": "O",
	"16": "P",
	"17": "Q",
	"18": "R",
	"19": "S",
	"20": "T",
	"21": "U",
	"22": "V",
	"23": "W",
	"24": "X",
	"25": "Y",
	"26": "Z",
}

type EntityExcel struct {
	ExcelFile *excelize.File
	IsNew     bool
	PathFile  string
}

/*
Check
*/

func checkExt(pathFile string) {
	ext := file.GetExt(pathFile)
	if ext == extExcel {
		return
	}
	msgError := fmt.Sprintf(
		"文件[%s]格式错误,正确后缀应是[%s],当前实际是[%s]",
		pathFile,
		extExcel,
		ext)
	err.ErrPanic(msgError)
}

/*
Save
*/

func (self *EntityExcel) SaveFile() *EntityExcel {
	pathFile := self.PathFile
	f := self.ExcelFile
	errExcelSave := f.SaveAs(pathFile)
	err.ErrCheck(errExcelSave)
	return self
}

/*
Write
*/

func (self *EntityExcel) WriteExcel(sheetName string, attrS2ExcelTitle base.AttrS2, attrS2ExcelData base.AttrS2) *EntityExcel {
	// self.setSheetEmpty(sheetName)
	self.batchSetCellOfTitle(sheetName, attrS2ExcelTitle)
	self.batchSetCellOfData(sheetName, attrS2ExcelTitle, attrS2ExcelData)
	return self
}

func (self *EntityExcel) setSheetEmpty(sheetName string) *EntityExcel {
	f := self.ExcelFile
	indexSheet, errSheet := f.NewSheet(sheetName)
	err.ErrCheck(errSheet)
	f.SetActiveSheet(indexSheet)
	return self
}

func (self *EntityExcel) batchSetCellOfTitle(sheetName string, attrS2ExcelTitle base.AttrS2) *EntityExcel {
	cellRow = 1
	for _, attrS1ExcelTitle := range attrS2ExcelTitle {
		sort, ok := attrS1ExcelTitle["sort"]
		if !ok {
			continue
		}
		cellColumn = getCellColumnBySort(sort)
		cellPosition = cellColumn + conv.ToStrFromInt(cellRow)
		title, ok := attrS1ExcelTitle["title"]
		if !ok {
			continue
		}
		cellValue = title
		self.setCellValue(sheetName, cellPosition, cellValue)
	}
	return self
}

func (self *EntityExcel) batchSetCellOfData(sheetName string, attrS2ExcelTitle base.AttrS2, attrS2ExcelData base.AttrS2) *EntityExcel {
	for sortRow, attrS1ExcelData := range attrS2ExcelData {
		cellRow = getCellRowBySortRow(sortRow)
		for field, value := range attrS1ExcelData {
			cellColumn = getCellColumnByField(attrS2ExcelTitle, field)
			cellPosition = cellColumn + conv.ToStrFromInt(cellRow)
			cellValue = value
			self.setCellValue(sheetName, cellPosition, cellValue)
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
	cellColumn, ok := attrCellColumn[sort]
	if !ok {
		msgError := fmt.Sprintf(
			"The cellColumn of sort |%s| not found",
			sort)
		err.ErrPanic(msgError)
	}
	return cellColumn
}

func (self *EntityExcel) setCellValue(sheetName string, cellPosition string, value string) *EntityExcel {
	f := self.ExcelFile
	f.SetCellValue(sheetName, cellPosition, value)
	return self
}

/*
Read
*/

func (self *EntityExcel) ReadExcel(sheetName string) (attrS2ExcelData base.AttrS2) {
	attrS2ExcelData = self.batchGetCell(sheetName)
	return attrS2ExcelData
}

func (self *EntityExcel) batchGetCell(sheetName string) (attrS2ExcelData base.AttrS2) {
	f := self.ExcelFile
	arrRows, errExcelGetRows := f.GetRows(sheetName)
	err.ErrCheck(errExcelGetRows)
	lenExcel := len(arrRows)
	if lenExcel < 2 {
		pathFile := self.PathFile
		msgError := fmt.Sprintf(
			"The excel |%s| is empty",
			pathFile)
		err.ErrPanic(msgError)
	}
	arrRowTitle := arrRows[0]
	countColumn := len(arrRowTitle)
	numColumnMax := countColumn - 1
	attrS2ExcelData = make(base.AttrS2)
	for numRow, arrRow := range arrRows {
		if numRow == 0 {
			continue
		}
		attrS1ExcelData := make(base.AttrS1)
		for numColumn, value := range arrRow {
			if numColumn > numColumnMax {
				msgError := fmt.Sprintf(
					"The excel title of column |%d| not found in |%s|",
					numColumn,
					self.PathFile)
				err.ErrPanic(msgError)
			}
			title := arrRowTitle[numColumn]
			attrS1ExcelData[title] = value
		}
		sortRow := conv.ToStrFromInt(numRow)
		attrS2ExcelData[sortRow] = attrS1ExcelData
	}
	return attrS2ExcelData
}

func (self *EntityExcel) getCellValue(sheetName string, cellPosition string) (cellValue string) {
	f := self.ExcelFile
	cellValue, errExcelGetCell := f.GetCellValue(sheetName, cellPosition)
	err.ErrCheck(errExcelGetCell)
	return cellValue
}
