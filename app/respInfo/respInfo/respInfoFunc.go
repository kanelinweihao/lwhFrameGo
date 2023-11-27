package respInfo

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
)

func GetRespInfoByRouteName(routeName string) (routeType int, paramsInDefault base.AttrT1, paramsOutDefault base.AttrT1, funcService conf.FuncService, pathTmpl string) {
	switch routeName {
	case "/":
		routeType = conf.RouteTypeWeb
		paramsInDefault, paramsOutDefault, funcService, pathTmpl = getRespInfoOfGetMobileNoAndOrgIdByShortUserId()
	case "/web/home":
		routeType = conf.RouteTypeWeb
		paramsInDefault, paramsOutDefault, funcService, pathTmpl = getRespInfoOfGetMobileNoAndOrgIdByShortUserId()
	case "/api/projectName":
		routeType = conf.RouteTypeApi
		paramsInDefault, paramsOutDefault, funcService, pathTmpl = getRespInfoOfProjectInfo()
	case "/web/projectName":
		routeType = conf.RouteTypeWeb
		paramsInDefault, paramsOutDefault, funcService, pathTmpl = getRespInfoOfProjectInfo()
	case "/api/projectVersion":
		routeType = conf.RouteTypeApi
		paramsInDefault, paramsOutDefault, funcService, pathTmpl = getRespInfoOfProjectInfo()
	case "/web/projectVersion":
		routeType = conf.RouteTypeWeb
		paramsInDefault, paramsOutDefault, funcService, pathTmpl = getRespInfoOfProjectInfo()
	case "/api/user/getMobileNoAndOrgIdByShortUserId":
		routeType = conf.RouteTypeApi
		paramsInDefault, paramsOutDefault, funcService, pathTmpl = getRespInfoOfGetMobileNoAndOrgIdByShortUserId()
	case "/web/user/getMobileNoAndOrgIdByShortUserId":
		routeType = conf.RouteTypeWeb
		paramsInDefault, paramsOutDefault, funcService, pathTmpl = getRespInfoOfGetMobileNoAndOrgIdByShortUserId()
	case "/api/user/getUserIdByMobileNoAndOrgId":
		routeType = conf.RouteTypeApi
		paramsInDefault, paramsOutDefault, funcService, pathTmpl = getRespInfoOfGetUserIdByMobileNoAndOrgId()
	case "/web/user/getUserIdByMobileNoAndOrgId":
		routeType = conf.RouteTypeWeb
		paramsInDefault, paramsOutDefault, funcService, pathTmpl = getRespInfoOfGetUserIdByMobileNoAndOrgId()
	default:
		msgError := fmt.Sprintf(
			"The respInfo of routeName |%s| not found",
			routeName)
		err.ErrPanic(msgError)
	}
	return routeType, paramsInDefault, paramsOutDefault, funcService, pathTmpl
}

func getParamsOutDefaultForWeb(paramsOutDefault base.AttrT1) (paramsOutDefaultForWeb base.AttrT1) {
	paramsOutFromENV := conf.GetParamsTmpl()
	paramsOutDefaultForWeb = base.MergeAttr(paramsOutDefault, paramsOutFromENV)
	return paramsOutDefaultForWeb
}
