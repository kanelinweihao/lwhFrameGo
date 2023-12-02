package dictRouteSwitch

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/dict/dictRoute/dictRouteData/dictRouteProjectInfo/dictRouteProjectInfo"
	"github.com/kanelinweihao/lwhFrameGo/app/dict/dictRoute/dictRouteData/dictRouteUser/dictRouteGetMobileNoAndOrgIdByShortId"
	"github.com/kanelinweihao/lwhFrameGo/app/dict/dictRoute/dictRouteData/dictRouteUser/dictRouteGetUserIdByMobileNoAndOrgId"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/funcAttr"
)

func GetDictRoute(routeName string) (routeType int, paramsInDefault base.AttrT1, paramsOutDefault base.AttrT1, funcService conf.FuncService, pathTmpl string) {
	switch routeName {
	case "/":
		routeType = conf.RouteTypeWeb
		paramsInDefault, paramsOutDefault, funcService, pathTmpl = dictRouteGetMobileNoAndOrgIdByShortId.GetDictRoute()
		paramsOutDefault = getParamsOutDefaultForWeb(paramsOutDefault)
	case "/web/home":
		routeType = conf.RouteTypeWeb
		paramsInDefault, paramsOutDefault, funcService, pathTmpl = dictRouteGetMobileNoAndOrgIdByShortId.GetDictRoute()
		paramsOutDefault = getParamsOutDefaultForWeb(paramsOutDefault)
	case "/api/projectName":
		routeType = conf.RouteTypeApi
		paramsInDefault, paramsOutDefault, funcService, pathTmpl = dictRouteProjectInfo.GetDictRoute()
	case "/web/projectName":
		routeType = conf.RouteTypeWeb
		paramsInDefault, paramsOutDefault, funcService, pathTmpl = dictRouteProjectInfo.GetDictRoute()
		paramsOutDefault = getParamsOutDefaultForWeb(paramsOutDefault)
	case "/api/projectVersion":
		routeType = conf.RouteTypeApi
		paramsInDefault, paramsOutDefault, funcService, pathTmpl = dictRouteProjectInfo.GetDictRoute()
	case "/web/projectVersion":
		routeType = conf.RouteTypeWeb
		paramsInDefault, paramsOutDefault, funcService, pathTmpl = dictRouteProjectInfo.GetDictRoute()
		paramsOutDefault = getParamsOutDefaultForWeb(paramsOutDefault)
	case "/api/user/getMobileNoAndOrgIdByShortUserId":
		routeType = conf.RouteTypeApi
		paramsInDefault, paramsOutDefault, funcService, pathTmpl = dictRouteGetMobileNoAndOrgIdByShortId.GetDictRoute()
	case "/web/user/getMobileNoAndOrgIdByShortUserId":
		routeType = conf.RouteTypeWeb
		paramsInDefault, paramsOutDefault, funcService, pathTmpl = dictRouteGetMobileNoAndOrgIdByShortId.GetDictRoute()
		paramsOutDefault = getParamsOutDefaultForWeb(paramsOutDefault)
	case "/api/user/getUserIdByMobileNoAndOrgId":
		routeType = conf.RouteTypeApi
		paramsInDefault, paramsOutDefault, funcService, pathTmpl = dictRouteGetUserIdByMobileNoAndOrgId.GetDictRoute()
	case "/web/user/getUserIdByMobileNoAndOrgId":
		routeType = conf.RouteTypeWeb
		paramsInDefault, paramsOutDefault, funcService, pathTmpl = dictRouteGetUserIdByMobileNoAndOrgId.GetDictRoute()
		paramsOutDefault = getParamsOutDefaultForWeb(paramsOutDefault)
	default:
		msgError := fmt.Sprintf(
			"The dictRoute of routeName |%s| not found",
			routeName)
		err.ErrPanic(msgError)
	}
	return routeType, paramsInDefault, paramsOutDefault, funcService, pathTmpl
}

func getParamsOutDefaultForWeb(paramsOutDefault base.AttrT1) (paramsOutDefaultForWeb base.AttrT1) {
	paramsOutFromENV := conf.GetParamsTmpl()
	paramsOutDefaultForWeb = funcAttr.MergeAttr(paramsOutDefault, paramsOutFromENV)
	return paramsOutDefaultForWeb
}
