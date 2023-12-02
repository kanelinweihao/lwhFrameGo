package dictRoute

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/dict/dictRoute/dictRouteSwitch"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

func GetDictRouteByRouteName(routeName string) (routeType int, paramsInDefault base.AttrT1, paramsOutDefault base.AttrT1, funcService conf.FuncService, pathTmpl string) {
	routeType, paramsInDefault, paramsOutDefault, funcService, pathTmpl = dictRouteSwitch.GetDictRoute(routeName)
	return routeType, paramsInDefault, paramsOutDefault, funcService, pathTmpl
}
