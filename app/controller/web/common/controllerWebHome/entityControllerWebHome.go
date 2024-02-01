package controllerWebHome

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/consts"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var EntityDataController *typeStruct.EntityDataController = &typeStruct.EntityDataController{
	Basic: typeStruct.EntityDataControllerBasic{
		RouteName:       "/projectVersion",
		RouteType:       consts.RouteTypeWeb,
		FuncService:     nil,
		ParamsInDefault: typeMap.AttrT1{},
		ParamsOutDefault: typeMap.AttrT1{
			"ArrAttrS1ForSectionListRouter": []typeMap.AttrS1{
				typeMap.AttrS1{
					"RouteNameCN": "0.1 客户端IP",
					"RouteName":   "/web/ip",
				},
				typeMap.AttrS1{
					"RouteNameCN": "1.1 获取客诉数据",
					"RouteName":   "/web/customerBehavior/getCustomerBehaviorByUserId",
				},
				typeMap.AttrS1{
					"RouteNameCN": "2.1 获取用户手机号和所属机构",
					"RouteName":   "/web/user/getMobileNoAndOrgIdByShortUserId",
				},
				typeMap.AttrS1{
					"RouteNameCN": "2.2 获取UID",
					"RouteName":   "/web/user/getUserIdByMobileNoAndOrgId",
				},
				typeMap.AttrS1{
					"RouteNameCN": "3.1 计算内部收益率IRR",
					"RouteName":   "/web/rate/getIRRByArrAmount",
				},
				typeMap.AttrS1{
					"RouteNameCN": "3.2 计算抽中次数正好是总次数一半的概率",
					"RouteName":   "/web/rate/getRateExactlyHalf",
				},
			},
		},
	},
	Web: typeStruct.EntityDataControllerWeb{
		PathTmpl:            "./res/view/common/home.tmpl",
		IsReqToApi:          false,
		RouteNameCN:         "列表主页",
		TextSectionMsg:      "\n",
		ArrEntitySectionIn:  nil,
		ArrEntitySectionOut: nil,
	},
}

type EntityController struct {
	typeInterface.EntityController
}

func (self *EntityController) GetEntityDataController() (entityDataController *typeStruct.EntityDataController) {
	entityDataController = EntityDataController
	return entityDataController
}
