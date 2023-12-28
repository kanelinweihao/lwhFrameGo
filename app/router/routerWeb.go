package router

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/common/controllerWebCommon"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/customerBehavior/controllerWebCustomerBehavior"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/rate/controllerWebRate"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/user/controllerWebUser"
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
		"/web/home",
		controllerWebCommon.Home)
	// 项目名称
	http.HandleFunc(
		"/web/projectName",
		controllerWebCommon.ProjectName)
	// 项目版本
	http.HandleFunc(
		"/web/projectVersion",
		controllerWebCommon.ProjectVersion)
	// user
	// 获取用户手机号和所属机构
	http.HandleFunc(
		"/web/user/getMobileNoAndOrgIdByShortUserId",
		controllerWebUser.GetMobileNoAndOrgIdByShortUserId)
	// 获取UID
	http.HandleFunc(
		"/web/user/getUserIdByMobileNoAndOrgId",
		controllerWebUser.GetUserIdByMobileNoAndOrgId)
	// customerBehavior
	// 获取客诉数据
	http.HandleFunc(
		"/web/customerBehavior/getCustomerBehaviorByUserId",
		controllerWebCustomerBehavior.GetCustomerBehaviorByUserId)
	// rate
	// 计算内部收益率IRR
	http.HandleFunc(
		"/web/rate/getIRRByArrAmount",
		controllerWebRate.GetIRRByArrAmount)
	// 计算抽中次数正好是总次数一半的概率
	http.HandleFunc(
		"/web/rate/getRateExactlyHalf",
		controllerWebRate.GetRateExactlyHalf)
	return
}
