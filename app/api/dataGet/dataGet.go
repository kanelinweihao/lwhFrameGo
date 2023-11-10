package dataGet

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

func GetData(paramsIn base.AttrT1, arrSQLName []string) (paramsOut base.AttrT1, boxData base.AttrS3) {
	entityDataGet := MakeEntityOfDataGet(paramsIn, arrSQLName)
	boxData, paramsOut = entityDataGet.GetData()
	return paramsOut, boxData
}
