package controllerWebUserGetMobileNoAndOrgIdByShortUserId

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/user/serviceUser"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/consts"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var EntityDataController *typeStruct.EntityDataController = &typeStruct.EntityDataController{
	Basic: typeStruct.EntityDataControllerBasic{
		RouteName:   "/web/user/getMobileNoAndOrgIdByShortUserId",
		RouteType:   consts.RouteTypeWeb,
		FuncService: serviceUser.GetMobileNoAndOrgIdByShortUserId,
		ParamsInDefault: typeMap.AttrT1{
			"ShortUserId": 0,
		},
		ParamsOutDefault: typeMap.AttrT1{
			"ShortUserId": 1,
			"UserId":      0,
			"MobileNo":    "0",
			"OrgId":       0,
		},
	},
	Web: typeStruct.EntityDataControllerWeb{
		PathTmpl:       "./res/view/auto/auto1.tmpl",
		IsReqToApi:     true,
		RouteNameCN:    "获取用户手机号和所属机构",
		TextSectionMsg: "\n",
		ArrEntitySectionIn: []typeStruct.EntitySection{
			typeStruct.EntitySection{
				FieldName:         "ShortUserId",
				FieldNameCN:       "用户编号简写",
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
				FieldName:         "UserId",
				FieldNameCN:       "用户编号",
				FieldRemark:       "",
				InputType:         "text",
				IsDivDisplay:      true,
				IsInputDisabled:   true,
				IsInputOnlyNumber: true,
				Value:             "",
			},
			typeStruct.EntitySection{
				FieldName:         "MobileNo",
				FieldNameCN:       "手机号",
				FieldRemark:       "",
				InputType:         "text",
				IsDivDisplay:      true,
				IsInputDisabled:   true,
				IsInputOnlyNumber: true,
				Value:             "",
			},
			typeStruct.EntitySection{
				FieldName:         "OrgId",
				FieldNameCN:       "机构编号",
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
