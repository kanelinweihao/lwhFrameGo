package conf

import "github.com/kanelinweihao/lwhFrameGo/app/utils/base"

const RouteTypeApi int = 1
const RouteTypeWeb int = 2

type FuncService func(paramsIn base.AttrT1) (paramsOut base.AttrT1)

var AttrRouteIgnoreValidate = base.AttrS1{
	"/favicon.ico": "",
	"/502.html":    "",
}
