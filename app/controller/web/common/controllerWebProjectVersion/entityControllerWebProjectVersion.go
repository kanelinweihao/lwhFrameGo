package controllerWebProjectVersion

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/consts"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var EntityDataController *typeStruct.EntityDataController = &typeStruct.EntityDataController{
	Basic: typeStruct.EntityDataControllerBasic{
		RouteName:       "/web/projectVersion",
		RouteType:       consts.RouteTypeWeb,
		FuncService:     nil,
		ParamsInDefault: typeMap.AttrT1{},
		ParamsOutDefault: typeMap.AttrT1{
			"ProjectVersion": "v1.0.0",
		},
	},
	Web: typeStruct.EntityDataControllerWeb{
		PathTmpl:            "./res/view/common/projectInfo.tmpl",
		IsReqToApi:          false,
		RouteNameCN:         "项目版本",
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
