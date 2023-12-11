package res

import "github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"

func GetResSuccess(paramsOut typeMap.AttrT1) (jsonRes string) {
	entityRes := initEntityResSuccess(paramsOut)
	jsonRes = entityRes.getResSuccess()
	return jsonRes
}
