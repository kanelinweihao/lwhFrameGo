package paramsOutGet

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

func GetParamsOut(paramsIn base.AttrT1) (paramsOut base.AttrT1) {
	entityParamsOutGet := MakeEntityParamsOutGet(paramsIn)
	paramsOut = entityParamsOutGet.ParamsOut
	return paramsOut
}
