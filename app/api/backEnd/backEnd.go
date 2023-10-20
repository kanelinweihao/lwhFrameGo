package backEnd

import (
	"go.lwh.com/linweihao/lwhFrameGo/app/api/cacheSet"
	"go.lwh.com/linweihao/lwhFrameGo/app/api/dataGet"
	"go.lwh.com/linweihao/lwhFrameGo/app/api/dataPut"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/base"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/time"
)

func ExecBackEnd(paramsIn base.AttrT1) (paramsOut base.AttrT1) {
	paramsOut, boxData := dataGet.GetData(paramsIn)
	dataPut.PutDataToExcel(paramsOut, boxData)
	cacheSet.SetCache(paramsOut)
	time.ShowTimeAndMsg("OK")
	return paramsOut
}
