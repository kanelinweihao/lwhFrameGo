package router

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/auth/controllerWebAuth"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/common/controllerWebCommon"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/customerBehavior/controllerWebCustomerBehavior"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/rate/controllerWebRate"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/user/controllerWebUser"
	"github.com/kanelinweihao/lwhFrameGo/app/middleware"
	"net/http"
)

func SetRouterWeb() {
	// common
	// 默认首页
	http.HandleFunc(
		"/",
		controllerWebCommon.Root)
	// 列表主页
	http.HandleFunc(
		"/home",
		controllerWebCommon.Home)
	// 客户端IP
	http.HandleFunc(
		"/ip",
		controllerWebCommon.ClientIP)
	// 项目名称
	http.HandleFunc(
		"/projectName",
		controllerWebCommon.ProjectName)
	// 项目版本
	http.HandleFunc(
		"/projectVersion",
		controllerWebCommon.ProjectVersion)
	// auth
	// 获取用户手机号和所属机构
	http.HandleFunc(
		"/web/auth/login",
		middleware.WithMiddlewareWeb(controllerWebAuth.Login))
	// user
	// 获取用户手机号和所属机构
	http.HandleFunc(
		"/web/user/getMobileNoAndOrgIdByShortUserId",
		middleware.WithMiddlewareWeb(controllerWebUser.GetMobileNoAndOrgIdByShortUserId))
	// 获取UID
	http.HandleFunc(
		"/web/user/getUserIdByMobileNoAndOrgId",
		middleware.WithMiddlewareWeb(controllerWebUser.GetUserIdByMobileNoAndOrgId))
	// customerBehavior
	// 获取客诉数据
	http.HandleFunc(
		"/web/customerBehavior/getCustomerBehaviorByUserId",
		middleware.WithMiddlewareWeb(controllerWebCustomerBehavior.GetCustomerBehaviorByUserId))
	// rate
	// 计算内部收益率IRR
	http.HandleFunc(
		"/web/rate/getIRRByArrAmount",
		middleware.WithMiddlewareWeb(controllerWebRate.GetIRRByArrAmount))
	// 计算抽中次数正好是总次数一半的概率
	http.HandleFunc(
		"/web/rate/getRateExactlyHalf",
		middleware.WithMiddlewareWeb(controllerWebRate.GetRateExactlyHalf))
	return
}
