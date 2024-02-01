package controllerWebRateGetIRRByArrAmount

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/rate/serviceRate"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/consts"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var EntityDataController *typeStruct.EntityDataController = &typeStruct.EntityDataController{
	Basic: typeStruct.EntityDataControllerBasic{
		RouteName:   "/web/rate/getIRRByArrAmount",
		RouteType:   consts.RouteTypeWeb,
		FuncService: serviceRate.GetIRRByArrAmount,
		ParamsInDefault: typeMap.AttrT1{
			"StrArrAmount":   "0,0,-1,-1.5,0,0,5,0,0",
			"ErrorPrecision": 4,
		},
		ParamsOutDefault: typeMap.AttrT1{
			"StrArrAmount":   "0,0,-1,-1.5,0,0,5,0,0",
			"ErrorPrecision": 4,
			"IRR":            "0",
		},
	},
	Web: typeStruct.EntityDataControllerWeb{
		PathTmpl:       "./res/view/auto/auto1.tmpl",
		IsReqToApi:     true,
		RouteNameCN:    "计算内部收益率IRR",
		TextSectionMsg: "\n",
		ArrEntitySectionIn: []typeStruct.EntitySection{
			typeStruct.EntitySection{
				FieldName:         "StrArrAmount",
				FieldNameCN:       "每期变化值",
				FieldRemark:       "用英文逗号隔开,负数为支出,正数为收入",
				InputType:         "text",
				IsDivDisplay:      true,
				IsInputDisabled:   false,
				IsInputOnlyNumber: false,
				Value:             "",
			},
			typeStruct.EntitySection{
				FieldName:         "ErrorPrecision",
				FieldNameCN:       "误差精度",
				FieldRemark:       "IRR小数表示时,小数点后保留位数",
				InputType:         "text",
				IsDivDisplay:      true,
				IsInputDisabled:   false,
				IsInputOnlyNumber: true,
				Value:             "",
			},
		},
		ArrEntitySectionOut: []typeStruct.EntitySection{
			typeStruct.EntitySection{
				FieldName:         "IRR",
				FieldNameCN:       "内部收益率",
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
