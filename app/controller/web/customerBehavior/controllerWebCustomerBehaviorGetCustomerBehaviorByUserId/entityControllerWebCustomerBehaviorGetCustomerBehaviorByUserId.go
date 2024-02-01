package controllerWebCustomerBehaviorGetCustomerBehaviorByUserId

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/customerBehavior/serviceCustomerBehavior"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/consts"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var EntityDataController *typeStruct.EntityDataController = &typeStruct.EntityDataController{
	Basic: typeStruct.EntityDataControllerBasic{
		RouteName:   "/web/customerBehavior/getCustomerBehaviorByUserId",
		RouteType:   consts.RouteTypeWeb,
		FuncService: serviceCustomerBehavior.GetCustomerBehaviorByUserId,
		ParamsInDefault: typeMap.AttrT1{
			"UserId": 0,
		},
		ParamsOutDefault: typeMap.AttrT1{
			"UserId":                   1005704,
			"CountProductOrderPoolNFT": 0,
			"CountProductOrderNFTBuy":  0,
			"CountProductOrderNFTSell": 0,
			"SumAmountGot":             "0",
		},
	},
	Web: typeStruct.EntityDataControllerWeb{
		PathTmpl:       "./res/view/auto/auto1.tmpl",
		IsReqToApi:     true,
		RouteNameCN:    "获取客诉数据",
		TextSectionMsg: "\n",
		ArrEntitySectionIn: []typeStruct.EntitySection{
			typeStruct.EntitySection{
				FieldName:         "UserId",
				FieldNameCN:       "用户编号",
				FieldRemark:       "",
				InputType:         "text",
				IsDivDisplay:      true,
				IsInputDisabled:   false,
				IsInputOnlyNumber: true,
				Value:             "",
			},
		},
		ArrEntitySectionOut: []typeStruct.EntitySection{
			typeStruct.EntitySection{
				FieldName:         "CountProductOrderPoolNFT",
				FieldNameCN:       "首发买入条数",
				FieldRemark:       "",
				InputType:         "text",
				IsDivDisplay:      true,
				IsInputDisabled:   true,
				IsInputOnlyNumber: true,
				Value:             "",
			},
			typeStruct.EntitySection{
				FieldName:         "CountProductOrderNFTBuy",
				FieldNameCN:       "寄售买入条数",
				FieldRemark:       "",
				InputType:         "text",
				IsDivDisplay:      true,
				IsInputDisabled:   true,
				IsInputOnlyNumber: true,
				Value:             "",
			},
			typeStruct.EntitySection{
				FieldName:         "CountProductOrderNFTSell",
				FieldNameCN:       "寄售卖出条数",
				FieldRemark:       "",
				InputType:         "text",
				IsDivDisplay:      true,
				IsInputDisabled:   true,
				IsInputOnlyNumber: true,
				Value:             "",
			},
			typeStruct.EntitySection{
				FieldName:         "SumAmountGot",
				FieldNameCN:       "用户水晶已领取总数",
				FieldRemark:       "",
				InputType:         "text",
				IsDivDisplay:      true,
				IsInputDisabled:   true,
				IsInputOnlyNumber: true,
				Value:             "",
			},
		},
	},
}

type EntityController struct {
	typeInterface.EntityController
}

func (self *EntityController) GetEntityDataController() (entityDataController *typeStruct.EntityDataController) {
	entityDataController = EntityDataController
	return entityDataController
}
