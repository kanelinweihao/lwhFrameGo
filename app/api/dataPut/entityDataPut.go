package dataPut

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/excel"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/time"
)

var sheetNameDefault string = "Sheet1"

type EntityDataPut struct {
	UserId          string
	PathDirThisTime string
	ParamsPathFile  base.AttrS1
	BoxExcelData    base.AttrS3
	BoxExcelTitle   base.AttrS3
}

/*
GET
*/

func (self *EntityDataPut) getUserId() (userId string) {
	userId = self.UserId
	if userId == "0" {
		msgError := "The |userId| is required"
		err.ErrPanic(msgError)
	}
	return userId
}

func errPanicExcelParamsRequired(shortName string, part string) {
	msgError := fmt.Sprintf(
		"The |%s| of |%s| is required",
		part,
		shortName)
	err.ErrPanic(msgError)
}

/*
EXEC
*/

func (self *EntityDataPut) BatchPutExcel() {
	paramsPathFile := self.ParamsPathFile
	boxExcelTitle := self.BoxExcelTitle
	boxExcelData := self.BoxExcelData
	for shortName, pathFile := range paramsPathFile {
		attrS2ExcelTitle, ok := boxExcelTitle[shortName]
		if !ok {
			errPanicExcelParamsRequired(shortName, "excelTitle")
		}
		attrS2ExcelData, ok := boxExcelData[shortName]
		if !ok {
			errPanicExcelParamsRequired(shortName, "excelData")
		}
		go putOne(
			shortName,
			pathFile,
			attrS2ExcelTitle,
			attrS2ExcelData)
	}
	return
}

func putOne(shortName string, pathFile string, attrS2ExcelTitle base.AttrS2, attrS2ExcelData base.AttrS2) {
	writeOne(pathFile, attrS2ExcelTitle, attrS2ExcelData)
	time.Sleep(100, "ms")
	// readOne(pathFile)
	return
}

func writeOne(pathFile string, attrS2ExcelTitle base.AttrS2, attrS2ExcelData base.AttrS2) {
	entityExcel := excel.MakeEntityOfExcelNew(pathFile)
	defer entityExcel.CloseExcel()
	entityExcel.WriteExcel(sheetNameDefault, attrS2ExcelTitle, attrS2ExcelData)
	entityExcel.SaveFile()
	return
}

func readOne(pathFile string) {
	entityExcel := excel.MakeEntityOfExcelOpen(pathFile)
	defer entityExcel.CloseExcel()
	entityExcel.ReadExcel(sheetNameDefault)
	return
}