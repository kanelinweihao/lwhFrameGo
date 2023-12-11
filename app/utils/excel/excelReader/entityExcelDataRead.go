package excelReader

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/xuri/excelize/v2"
)

type EntityExcelDataRead struct {
	AttrS2ExcelData typeMap.AttrS2
	PathFile        string
	SheetName       string
	ExcelFile       *excelize.File
}

func (self *EntityExcelDataRead) WriteToChannelOfReadExcelData(chanWrite chan<- typeChanData) *EntityExcelDataRead {
	defer self.CloseExcel()
	self.readExcel().writeChan(chanWrite)
	return self
}

func (self *EntityExcelDataRead) readExcel() *EntityExcelDataRead {
	self.batchGetCell()
	return self
}

func (self *EntityExcelDataRead) batchGetCell() *EntityExcelDataRead {
	sheetName := self.SheetName
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
	attrS2ExcelData := make(typeMap.AttrS2)
	for numRow, arrRow := range arrRows {
		if numRow == 0 {
			continue
		}
		attrS1ExcelData := make(typeMap.AttrS1)
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
	self.AttrS2ExcelData = attrS2ExcelData
	return self
}

func (self *EntityExcelDataRead) getCellValue(sheetName string, cellPosition string) (cellValue string) {
	f := self.ExcelFile
	cellValue, errExcelGetCell := f.GetCellValue(sheetName, cellPosition)
	err.ErrCheck(errExcelGetCell)
	return cellValue
}

func (self *EntityExcelDataRead) writeChan(chanWrite chan<- typeChanData) *EntityExcelDataRead {
	attrS2ExcelData := self.AttrS2ExcelData
	chanWrite <- attrS2ExcelData
	close(chanWrite)
	return self
}

func readFromChannelOfReadExcelData(chanRead <-chan typeChanData) (attrS2ExcelData typeMap.AttrS2) {
	attrS2ExcelData = <-chanRead
	return attrS2ExcelData
}
