package params

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/service/base/paramsBase"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/jwtt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
)

func InitEntityParams(paramsIn typeMap.AttrT1) (entityParams typeInterface.EntityParams) {
	entityParams = new(EntityParams)
	entityParamsBase := new(paramsBase.EntityParamsBase)
	entityParams.Load(entityParamsBase).Init(paramsIn)
	return entityParams
}

func (self *EntityParams) Load(entityParamsBase typeInterface.EntityParams) typeInterface.EntityParams {
	self.EntityParams = entityParamsBase
	entityParamsBase.Load(self)
	return self
}

func (self *EntityParams) SetParamsExec() typeInterface.EntityParams {
	self.checkAdminUserName().checkAdminUserPassword().setParamsFromEntityAdminUser()
	return self
}

func (self *EntityParams) checkAdminUserName() *EntityParams {
	adminUserName := self.AdminUserName
	entityAdminUser := conf.GetEntityAdminUser(adminUserName)
	self.EntityAdminUser = entityAdminUser
	return self
}

func (self *EntityParams) checkAdminUserPassword() *EntityParams {
	entityAdminUser := self.EntityAdminUser
	adminUserPasswordTrue := entityAdminUser.AdminUserPassword
	adminUserPasswordToCheck := self.AdminUserPassword
	if adminUserPasswordToCheck != adminUserPasswordTrue {
		adminUserName := self.AdminUserName
		msgError := fmt.Sprintf(
			"The adminUserPassword |%s| of adminUserName |%s| is invalid",
			adminUserPasswordToCheck,
			adminUserName)
		err.ErrPanic(msgError)
	}
	return self
}

func (self *EntityParams) setParamsFromEntityAdminUser() *EntityParams {
	self.AdminUserId = self.EntityAdminUser.AdminUserId
	self.AdminUserRole = self.EntityAdminUser.AdminUserRole
	self.JwtToken = self.getJwtToken()
	jsonAuth := jwtt.GetJsonAuth(self.JwtToken)
	dd.IGNORE(jsonAuth)
	return self
}

func (self *EntityParams) getJwtToken() (jwtToken string) {
	entityAdminUser := self.EntityAdminUser
	jsonAuth := entityAdminUser.ToJson()
	jwtToken = jwtt.GetJwtToken(jsonAuth)
	return jwtToken
}
