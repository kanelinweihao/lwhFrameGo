package serviceUser

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/serviceUser/serviceGetMobileNoAndOrgIdByShortUserId"
	"github.com/kanelinweihao/lwhFrameGo/app/service/serviceUser/serviceGetUserIdByMobileNoAndOrgId"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

var entityService base.TEntityService

func GetMobileNoAndOrgIdByShortUserId(paramsIn base.AttrT1) (paramsOut base.AttrT1) {
	entityService = serviceGetMobileNoAndOrgIdByShortUserId.InitEntityService(paramsIn)
	paramsOut = entityService.Exec()
	return paramsOut
}

func GetUserIdByMobileNoAndOrgId(paramsIn base.AttrT1) (paramsOut base.AttrT1) {
	entityService = serviceGetUserIdByMobileNoAndOrgId.InitEntityService(paramsIn)
	paramsOut = entityService.Exec()
	return paramsOut
}
