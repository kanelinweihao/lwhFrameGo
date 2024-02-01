package controllerWebRateGetRateExactlyHalf

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/rate/serviceRate"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/consts"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var EntityDataController *typeStruct.EntityDataController = &typeStruct.EntityDataController{
	Basic: typeStruct.EntityDataControllerBasic{
		RouteName:   "/web/rate/getRateExactlyHalf",
		RouteType:   consts.RouteTypeWeb,
		FuncService: serviceRate.GetRateExactlyHalf,
		ParamsInDefault: typeMap.AttrT1{
			"X": 10,
		},
		ParamsOutDefault: typeMap.AttrT1{
			"X":    10,
			"Rate": "0",
		},
	},
	Web: typeStruct.EntityDataControllerWeb{
		PathTmpl:       "./res/view/auto/auto1.tmpl",
		IsReqToApi:     true,
		RouteNameCN:    "计算抽中次数正好是总次数一半的概率",
		TextSectionMsg: "假设每次抽奖的中奖概率为[50%](即一半概率抽中一半概率抽不中),\n若总共抽奖[X]次,求抽中次数正好为总次数的一半的概率[Rate]\n",
		ArrEntitySectionIn: []typeStruct.EntitySection{
			typeStruct.EntitySection{
				FieldName:         "X",
				FieldNameCN:       "总抽奖次数",
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
				FieldName:         "Rate",
				FieldNameCN:       "目标概率",
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
