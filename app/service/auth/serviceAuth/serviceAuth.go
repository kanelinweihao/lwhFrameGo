package serviceAuth

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/auth/serviceLogin"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

var entityService typeInterface.EntityService2

func Login(paramsIn typeMap.AttrT1) (paramsOut typeMap.AttrT1) {
	entityService = serviceLogin.InitEntityService()
	paramsOut = entityService.Exec(paramsIn)
	return paramsOut
}
