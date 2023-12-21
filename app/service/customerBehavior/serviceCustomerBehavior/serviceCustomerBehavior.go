package serviceCustomerBehavior

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/customerBehavior/serviceGetCustomerBehaviorByUserId"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var entityService typeStruct.EntityService1

func GetCustomerBehaviorByUserId(paramsIn typeMap.AttrT1) (paramsOut typeMap.AttrT1) {
	entityService = serviceGetCustomerBehaviorByUserId.InitEntityService()
	paramsOut = entityService.Exec(paramsIn)
	return paramsOut
}
