package serviceRate

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/rate/serviceGetIRRByArrAmount"
	"github.com/kanelinweihao/lwhFrameGo/app/service/rate/serviceGetRateExactlyHalf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var entityService typeStruct.EntityService2

func GetIRRByArrAmount(paramsIn typeMap.AttrT1) (paramsOut typeMap.AttrT1) {
	entityService = serviceGetIRRByArrAmount.InitEntityService()
	paramsOut = entityService.Exec(paramsIn)
	return paramsOut
}

func GetRateExactlyHalf(paramsIn typeMap.AttrT1) (paramsOut typeMap.AttrT1) {
	entityService = serviceGetRateExactlyHalf.InitEntityService()
	paramsOut = entityService.Exec(paramsIn)
	return paramsOut
}
