package serviceCustomerBehavior

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/customerBehavior/serviceGetCustomerBehaviorByUserId"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

var entityService typeInterface.EntityService1

func GetCustomerBehaviorByUserId(paramsIn typeMap.AttrT1) (paramsOut typeMap.AttrT1) {
	entityService = serviceGetCustomerBehaviorByUserId.InitEntityService()
	paramsOut = entityService.Exec(paramsIn)
	return paramsOut
}
