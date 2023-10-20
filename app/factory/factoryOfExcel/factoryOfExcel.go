package factoryOfExcel

import (
	"fmt"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/excel"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/err"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/file"
)

var extExcel string = "xlsx"

func MakeEntityOfExcelNew(pathFile string) (entityExcel *excel.EntityExcel) {
	checkExt(pathFile)
	entityExcel = excel.InitEntityExcelNew(pathFile)
	entityExcel.SaveFile()
	return entityExcel
}

func MakeEntityOfExcelOpen(pathFile string) (entityExcel *excel.EntityExcel) {
	checkExt(pathFile)
	entityExcel = excel.InitEntityExcelOpen(pathFile)
	return entityExcel
}

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
