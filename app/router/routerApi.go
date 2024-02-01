package router

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/api/auth/controllerApiAuth"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/api/common/controllerApiCommon"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/api/customerBehavior/controllerApiCustomerBehavior"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/api/rate/controllerApiRate"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/api/user/controllerApiUser"
	"github.com/kanelinweihao/lwhFrameGo/app/middleware"
)

func SetRouterApi() {
	arrRouterApi := getArrRouterApi()
	setRouter(arrRouterApi)
}

func getArrRouterApi() (arrRouterApi []Router) {
	arrRouterApi = []Router{
		// common
		// 项目名称
		Router{
			RouteName:      "/api/projectTitle",
			FuncController: controllerApiCommon.ProjectTitle,
			FuncMiddleware: middleware.Api,
		},
		// 项目版本
		Router{
			RouteName:      "/api/projectVersion",
			FuncController: controllerApiCommon.ProjectVersion,
			FuncMiddleware: middleware.Api,
		},
		// 客户端IP
		Router{
			RouteName:      "/api/ip",
			FuncController: controllerApiCommon.ClientIP,
			FuncMiddleware: middleware.Api,
		},
		// auth
		// 登录
		Router{
			RouteName:      "/api/auth/login",
			FuncController: controllerApiAuth.Login,
			FuncMiddleware: middleware.Api,
		},
		// customerBehavior
		// 获取客诉数据
		Router{
			RouteName:      "/api/customerBehavior/getCustomerBehaviorByUserId",
			FuncController: controllerApiCustomerBehavior.GetCustomerBehaviorByUserId,
			FuncMiddleware: middleware.Api,
		},
		// user
		// 获取用户手机号和所属机构
		Router{
			RouteName:      "/api/user/getMobileNoAndOrgIdByShortUserId",
			FuncController: controllerApiUser.GetMobileNoAndOrgIdByShortUserId,
			FuncMiddleware: middleware.Api,
		},
		// 获取UID
		Router{
			RouteName:      "/api/user/getUserIdByMobileNoAndOrgId",
			FuncController: controllerApiUser.GetUserIdByMobileNoAndOrgId,
			FuncMiddleware: middleware.Api,
		},
		// rate
		// 计算内部收益率IRR
		Router{
			RouteName:      "/api/rate/getIRRByArrAmount",
			FuncController: controllerApiRate.GetIRRByArrAmount,
			FuncMiddleware: middleware.Api,
		},
		// 计算抽中次数正好是总次数一半的概率
		Router{
			RouteName:      "/api/rate/getRateExactlyHalf",
			FuncController: controllerApiRate.GetRateExactlyHalf,
			FuncMiddleware: middleware.Api,
		},
	}
	return arrRouterApi
}
