package serviceGetMobileNoAndOrgIdByShortUserId

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/serviceUser/serviceGetMobileNoAndOrgIdByShortUserId/service"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

func Exec(paramsIn base.AttrT1) (paramsOut base.AttrT1) {
	entityService := service.MakeEntityService(paramsIn)
	paramsOut = entityService.Exec()
	return paramsOut
}
