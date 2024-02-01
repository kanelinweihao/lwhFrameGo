package controllerWebAuthLogin

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/auth/serviceAuth"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/consts"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var EntityDataController *typeStruct.EntityDataController = &typeStruct.EntityDataController{
	Basic: typeStruct.EntityDataControllerBasic{
		RouteName:   "/web/auth/login",
		RouteType:   consts.RouteTypeWeb,
		FuncService: serviceAuth.Login,
		ParamsInDefault: typeMap.AttrT1{
			"AdminUserName":     "",
			"AdminUserPassword": "",
		},
		ParamsOutDefault: typeMap.AttrT1{
			"AdminUserName":     "test",
			"AdminUserPassword": "123abc",
			"GoogleCode":        "123456",
			"AdminUserId":       0,
			"AdminUserRole":     0,
			"JwtToken":          "",
		},
	},
	Web: typeStruct.EntityDataControllerWeb{
		PathTmpl:       "./res/view/auto/auto1.tmpl",
		IsReqToApi:     true,
		RouteNameCN:    "登录",
		TextSectionMsg: "\n",
		ArrEntitySectionIn: []typeStruct.EntitySection{
			typeStruct.EntitySection{
				FieldName:         "AdminUserName",
				FieldNameCN:       "用户名",
				FieldRemark:       "",
				InputType:         "text",
				IsDivDisplay:      true,
				IsInputDisabled:   false,
				IsInputOnlyNumber: false,
				Value:             "",
			},
			typeStruct.EntitySection{
				FieldName:         "AdminUserPassword",
				FieldNameCN:       "密码",
				FieldRemark:       "",
				InputType:         "password",
				IsDivDisplay:      true,
				IsInputDisabled:   false,
				IsInputOnlyNumber: false,
				Value:             "",
			},
			typeStruct.EntitySection{
				FieldName:         "GoogleCode",
				FieldNameCN:       "谷歌验证码",
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
				FieldName:         "AdminUserId",
				FieldNameCN:       "管理员编号",
				FieldRemark:       "",
				InputType:         "text",
				IsDivDisplay:      true,
				IsInputDisabled:   true,
				IsInputOnlyNumber: true,
				Value:             "",
			},
			typeStruct.EntitySection{
				FieldName:         "AdminUserRole",
				FieldNameCN:       "管理员角色",
				FieldRemark:       "",
				InputType:         "text",
				IsDivDisplay:      true,
				IsInputDisabled:   true,
				IsInputOnlyNumber: true,
				Value:             "",
			},
			typeStruct.EntitySection{
				FieldName:         "JwtToken",
				FieldNameCN:       "登录凭证",
				FieldRemark:       "",
				InputType:         "text",
				IsDivDisplay:      false,
				IsInputDisabled:   true,
				IsInputOnlyNumber: false,
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
