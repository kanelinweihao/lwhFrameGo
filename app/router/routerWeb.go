package router

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/auth/controllerWebAuth"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/common/controllerWebCommon"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/customerBehavior/controllerWebCustomerBehavior"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/rate/controllerWebRate"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/user/controllerWebUser"
	"github.com/kanelinweihao/lwhFrameGo/app/middleware"
)

func SetRouterWeb() {
	arrRouterWeb := getArrRouterWeb()
	setRouter(arrRouterWeb)
}

func getArrRouterWeb() (arrRouterWeb []Router) {
	arrRouterWeb = []Router{
		// common
		// 默认首页
		Router{
			RouteName:      "/",
			FuncController: controllerWebCommon.Root,
			FuncMiddleware: middleware.WithMiddlewareWeb,
		},
		// 列表主页
		Router{
			RouteName:      "/home",
			FuncController: controllerWebCommon.Home,
			FuncMiddleware: middleware.WithMiddlewareWeb,
		},
		// 项目名称
		Router{
			RouteName:      "/projectTitle",
			FuncController: controllerWebCommon.ProjectTitle,
			FuncMiddleware: middleware.WithMiddlewareWeb,
		},
		// 项目版本
		Router{
			RouteName:      "/projectVersion",
			FuncController: controllerWebCommon.ProjectVersion,
			FuncMiddleware: middleware.WithMiddlewareWeb,
		},
		// 客户端IP
		Router{
			RouteName:      "/ip",
			FuncController: controllerWebCommon.ClientIP,
			FuncMiddleware: middleware.WithMiddlewareWeb,
		},
		// auth
		// 获取用户手机号和所属机构
		Router{
			RouteName:      "/web/auth/login",
			FuncController: controllerWebAuth.Login,
			FuncMiddleware: middleware.WithMiddlewareWeb,
		},
		// user
		// 获取用户手机号和所属机构
		Router{
			RouteName:      "/web/user/getMobileNoAndOrgIdByShortUserId",
			FuncController: controllerWebUser.GetMobileNoAndOrgIdByShortUserId,
			FuncMiddleware: middleware.WithMiddlewareWeb,
		},
		// 获取UID
		Router{
			RouteName:      "/web/user/getUserIdByMobileNoAndOrgId",
			FuncController: controllerWebUser.GetUserIdByMobileNoAndOrgId,
			FuncMiddleware: middleware.WithMiddlewareWeb,
		},
		// customerBehavior
		// 获取客诉数据
		Router{
			RouteName:      "/web/customerBehavior/getCustomerBehaviorByUserId",
			FuncController: controllerWebCustomerBehavior.GetCustomerBehaviorByUserId,
			FuncMiddleware: middleware.WithMiddlewareWeb,
		},
		// rate
		// 计算内部收益率IRR
		Router{
			RouteName:      "/web/rate/getIRRByArrAmount",
			FuncController: controllerWebRate.GetIRRByArrAmount,
			FuncMiddleware: middleware.WithMiddlewareWeb,
		},
		// 计算抽中次数正好是总次数一半的概率
		Router{
			RouteName:      "/web/rate/getRateExactlyHalf",
			FuncController: controllerWebRate.GetRateExactlyHalf,
			FuncMiddleware: middleware.WithMiddlewareWeb,
		},
	}
	return arrRouterWeb
}
