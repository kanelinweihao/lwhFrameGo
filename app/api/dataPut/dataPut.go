package dataPut

import (
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/base"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/time"
)

func PutDataToExcel(paramsOut base.AttrT1, boxExcelData base.AttrS3) {
	// time.ShowTimeAndMsg("Data put begin")
	entityDataPut := &EntityDataPut{}
	// dd.DD(entityDataPut)
	entityDataPut.SetUserId(paramsOut)
	entityDataPut.SetBoxExcelData(boxExcelData)
	entityDataPut.SetBoxExcelTitle()
	entityDataPut.SetPathDirThisTime()
	entityDataPut.SetParamsPathFile()
	// dd.DD(entityDataPut)
	entityDataPut.BatchPutExcel()
	time.ShowTimeAndMsg("Data put success")
	return
}
