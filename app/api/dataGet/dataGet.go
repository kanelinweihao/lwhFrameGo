package dataGet

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
)

func GetData(paramsIn base.AttrT1) (paramsOut base.AttrT1, boxData base.AttrS3) {
	entityDataGet := MakeEntityOfDataGet(paramsIn)
	boxData, paramsOut = entityDataGet.GetData()
	return paramsOut, boxData
}
