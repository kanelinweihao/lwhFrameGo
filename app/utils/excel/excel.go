package excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/base"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/conv"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/err"
)

var sheetName string = "Sheet1"
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
	PathFile string
}

////
// Init
////

func InitEntityExcelNew(pathFile string) (entityExcel *EntityExcel) {
	f := excelize.NewFile()
	entityExcel = &EntityExcel{
		ExcelFile : f,
		PathFile : pathFile,
	}
	return entityExcel
}

func InitEntityExcelOpen(pathFile string) (entityExcel *EntityExcel) {
	f, errExcelOpen := excelize.OpenFile(pathFile)
	err.ErrCheck(errExcelOpen)
	entityExcel = &EntityExcel{
		ExcelFile : f,
		PathFile : pathFile,
	}
	return entityExcel
}

////
// Close
////

func (self *EntityExcel) CloseExcel() {
	f := self.ExcelFile
	errExcelClose := f.Close()
	err.ErrCheck(errExcelClose)
	return
}

////
// Save
////

func (self *EntityExcel) SaveFile() {
	f := self.ExcelFile
	pathFile := self.PathFile
	errExcelSave := f.SaveAs(pathFile)
	err.ErrCheck(errExcelSave)
	return
}

////
// Write
////

func (self *EntityExcel) WriteExcel(attrS2ExcelTitle base.AttrS2, attrS2ExcelData base.AttrS2) {
	// dd.DDD(attrS2ExcelTitle, attrS2ExcelData)
	f := self.ExcelFile
	batchSetCell(
		f,
		sheetName,
		attrS2ExcelTitle,
		attrS2ExcelData)
	self.SaveFile()
	return
}

func setSheetNew(f *excelize.File, numSheet int) {
	numString := conv.ToStrFromInt(numSheet)
	sheetNameNew := "Sheet" + numString
	numSheet, errSheet := f.NewSheet(sheetNameNew)
	err.ErrCheck(errSheet)
	f.SetActiveSheet(numSheet)
	return
}

func batchSetCell(f *excelize.File, sheetName string, attrS2ExcelTitle base.AttrS2, attrS2ExcelData base.AttrS2) {
	var cellColumn string
	var cellRow int
	var cellPosition string
	var cellValue string
	for sortRow, attrS1ExcelData := range attrS2ExcelData {
		numRow := conv.ToIntFromStr(sortRow)
		// dd.DD(numRow)
		// dd.DD(attrS1ExcelData)
		cellRow = numRow + 2
		// dd.DD(cellRow)
		for field, value := range attrS1ExcelData {
			// dd.DD(field)
			// dd.DD(value)
			cellColumn = getCellColumnByField(attrS2ExcelTitle, field)
			if cellColumn == "未知表格列序号" {
				continue
			}
			cellPosition = fmt.Sprintf(
				"%s%d",
				cellColumn,
				cellRow)
			// dd.DD(cellPosition)
			cellValue = value
			setCellValue(
				f,
				sheetName,
				cellPosition,
				cellValue)
			if sortRow == "0" {
				batchSetCellOfTitle(
					f,
					sheetName,
					cellColumn,
					attrS2ExcelTitle,
					field)
			}
		}
	}
}

func getCellColumnByField(attrS2ExcelTitle base.AttrS2, field string) (cellColumn string) {
	// dd.DD(field)
	// dd.DD(attrS2ExcelTitle)
	sort := getSortByField(attrS2ExcelTitle, field)
	// dd.DD(sort)
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
	// dd.DD(cellColumn)
	return cellColumn
}

func setCellValue(f *excelize.File, sheetName string, cellPosition string, value string) {
	f.SetCellValue(
		sheetName,
		cellPosition,
		value)
	return
}

func batchSetCellOfTitle(f *excelize.File, sheetName string, cellColumn string, attrS2ExcelTitle base.AttrS2, field string) {
	cellRowTitle := 1
	cellPositionTitle := fmt.Sprintf(
		"%s%d",
		cellColumn,
		cellRowTitle)
	title := getTitle(attrS2ExcelTitle, field)
	f.SetCellValue(
		sheetName,
		cellPositionTitle,
		title)
	return
}

func getTitle(attrS2ExcelTitle base.AttrS2, field string) (title string) {
	title, ok := attrS2ExcelTitle[field]["title"]
	if !ok {
		title = "未知字段"
	}
	return title
}

////
// Read
////

func (self *EntityExcel) ReadExcel() (attrS2ExcelData base.AttrS2) {
	f := self.ExcelFile
	attrS2ExcelData = batchGetCell(f, sheetName)
	// dd.DD(attrS2ExcelData)
	return attrS2ExcelData
}

func batchGetCell(f *excelize.File, sheetName string) (attrS2ExcelData base.AttrS2) {
	arrRows, errExcelGetRows := f.GetRows(sheetName)
	err.ErrCheck(errExcelGetRows)
	// dd.DD(arrRows)
	lenExcel := len(arrRows)
	if lenExcel < 2 {
		pathFile := f.Path
		msgError := fmt.Sprintf(
			"The excel |%s| is empty",
			pathFile)
		err.ErrPanic(msgError)
	}
	arrRowTitle := arrRows[0]
	// dd.DD(arrRowTitle)
	attrS2ExcelData = base.AttrS2{}
	for numRow, arrRow := range arrRows {
		// dd.DDD(numRow, arrRow)
		if numRow == 0 {
			continue
		}
		attrS1ExcelData := base.AttrS1{}
		for numCellColumn, value := range arrRow {
			// dd.DD(numCellColumn)
			// dd.DD(value)
			title := arrRowTitle[numCellColumn]
			attrS1ExcelData[title] = value
		}
		sortRow := conv.ToStrFromInt(numRow)
		attrS2ExcelData[sortRow] = attrS1ExcelData
	}
	// dd.DD(attrS2ExcelData)
	return attrS2ExcelData
}

func getCellValue(f *excelize.File, sheetName string, cellPosition string) (cellValue string) {
	cellValue, errExcelGetCell := f.GetCellValue(sheetName, cellPosition)
	err.ErrCheck(errExcelGetCell)
	return cellValue
}
