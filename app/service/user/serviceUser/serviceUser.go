package serviceUser

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/user/serviceGetMobileNoAndOrgIdByShortUserId"
	"github.com/kanelinweihao/lwhFrameGo/app/service/user/serviceGetUserIdByMobileNoAndOrgId"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

var entityService typeInterface.EntityService1

func GetMobileNoAndOrgIdByShortUserId(paramsIn typeMap.AttrT1) (paramsOut typeMap.AttrT1) {
	entityService = serviceGetMobileNoAndOrgIdByShortUserId.InitEntityService()
	paramsOut = entityService.Exec(paramsIn)
	return paramsOut
}

func GetUserIdByMobileNoAndOrgId(paramsIn typeMap.AttrT1) (paramsOut typeMap.AttrT1) {
	entityService = serviceGetUserIdByMobileNoAndOrgId.InitEntityService()
	paramsOut = entityService.Exec(paramsIn)
	return paramsOut
}
