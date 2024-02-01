package controllerWebProjectTitle

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/consts"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var EntityDataController *typeStruct.EntityDataController = &typeStruct.EntityDataController{
	Basic: typeStruct.EntityDataControllerBasic{
		RouteName:       "/web/projectTitle",
		RouteType:       consts.RouteTypeWeb,
		FuncService:     nil,
		ParamsInDefault: typeMap.AttrT1{},
		ParamsOutDefault: typeMap.AttrT1{
			"ProjectTitle": "未知项目名称",
		},
	},
	Web: typeStruct.EntityDataControllerWeb{
		PathTmpl:            "./res/view/common/projectInfo.tmpl",
		IsReqToApi:          false,
		RouteNameCN:         "项目名称",
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
