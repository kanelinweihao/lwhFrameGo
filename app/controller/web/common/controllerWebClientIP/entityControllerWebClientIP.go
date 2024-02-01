package controllerWebClientIP

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/consts"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var EntityDataController *typeStruct.EntityDataController = &typeStruct.EntityDataController{
	Basic: typeStruct.EntityDataControllerBasic{
		RouteName:       "/web/ip",
		RouteType:       consts.RouteTypeWeb,
		FuncService:     nil,
		ParamsInDefault: typeMap.AttrT1{},
		ParamsOutDefault: typeMap.AttrT1{
			"ClientIP": "",
		},
	},
	Web: typeStruct.EntityDataControllerWeb{
		PathTmpl:            "./res/view/common/projectInfo.tmpl",
		IsReqToApi:          false,
		RouteNameCN:         "客户端IP",
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
