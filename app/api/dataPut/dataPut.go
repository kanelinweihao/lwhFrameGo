package dataPut

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
)

func PutData(paramsOut base.AttrT1, boxExcelData base.AttrS3) (pathDirExcel string, arrPathFileExcel []string) {
	entityDataPut := MakeEntityDataPut(paramsOut, boxExcelData)
	entityDataPut.PutData()
	pathDirExcel = entityDataPut.PathDirThisTime
	arrPathFileExcel = conv.ToArrStrFromAttrStr(entityDataPut.ParamsPathFile)
	return pathDirExcel, arrPathFileExcel
}
