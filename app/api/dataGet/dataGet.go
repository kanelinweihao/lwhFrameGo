package dataGet

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

func GetData(paramsOut base.AttrT1) (boxData base.AttrS3) {
	entityDataGet := MakeEntityDataGet(paramsOut)
	boxData = entityDataGet.GetData()
	return boxData
}
