package dataPut

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

func PutDataToExcel(paramsOut base.AttrT1, boxExcelData base.AttrS3) {
	entityDataPut := MakeEntityOfDataPut(paramsOut, boxExcelData)
	entityDataPut.BatchPutExcel()
	return
}
