package dataGet

import (
	_ "fmt"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/base"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/conv"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/err"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/time"
)

func GetData(paramsIn base.AttrT1) (paramsOut base.AttrT1, boxData base.AttrS3) {
	defer err.ThrowError()
	// dd.DD(paramsIn)
	entityDataGet := &EntityDataGet{}
	// dd.DD(entityDataGet)
	conv.ToEntityFromAttr(paramsIn, entityDataGet)
	// fmt.Printf("%#v\n\n", entityDataGet)
	entityDataGet.Exec()
	// fmt.Printf("%#v\n\n", entityDataGet)
	boxData = entityDataGet.GetBoxData()
	paramsOut = entityDataGet.GetParamsOut()
	// dd.DD(paramsOut)
	time.ShowTimeAndMsg("Data get success")
	return paramsOut, boxData
}
