package router

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/api/common/controllerApiCommon"
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
			FuncMiddleware: middleware.WithMiddlewareApi,
		},
		// 项目版本
		Router{
			RouteName:      "/api/projectVersion",
			FuncController: controllerApiCommon.ProjectVersion,
			FuncMiddleware: middleware.WithMiddlewareApi,
		},
		// 客户端IP
		Router{
			RouteName:      "/api/ip",
			FuncController: controllerApiCommon.ClientIP,
			FuncMiddleware: middleware.WithMiddlewareApi,
		},
		// user
		// 获取用户手机号和所属机构
		Router{
			RouteName:      "/api/user/getMobileNoAndOrgIdByShortUserId",
			FuncController: controllerApiUser.GetMobileNoAndOrgIdByShortUserId,
			FuncMiddleware: middleware.WithMiddlewareApi,
		},
		// 获取UID
		Router{
			RouteName:      "/api/user/getUserIdByMobileNoAndOrgId",
			FuncController: controllerApiUser.GetUserIdByMobileNoAndOrgId,
			FuncMiddleware: middleware.WithMiddlewareApi,
		},
	}
	return arrRouterApi
}
