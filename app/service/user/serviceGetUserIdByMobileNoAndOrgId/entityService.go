package serviceGetUserIdByMobileNoAndOrgId

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/user/serviceGetUserIdByMobileNoAndOrgId/notify/notifyEmail"
	"github.com/kanelinweihao/lwhFrameGo/app/service/user/serviceGetUserIdByMobileNoAndOrgId/notify/notifyExcel"
	"github.com/kanelinweihao/lwhFrameGo/app/service/user/serviceGetUserIdByMobileNoAndOrgId/params"
	"github.com/kanelinweihao/lwhFrameGo/app/service/user/serviceGetUserIdByMobileNoAndOrgId/repo/repoCache"
	"github.com/kanelinweihao/lwhFrameGo/app/service/user/serviceGetUserIdByMobileNoAndOrgId/repo/repoDB"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

type EntityService struct {
	typeInterface.EntityService1
}

func (self *EntityService) GetFuncInitEntityParams() (funcInitEntityParams typeInterface.FuncInitEntityParams) {
	funcInitEntityParams = params.InitEntityParams
	return funcInitEntityParams
}

func (self *EntityService) GetFuncInitEntityRepoDB() (funcInitEntityRepoDB typeInterface.FuncInitEntityRepoDB) {
	funcInitEntityRepoDB = repoDB.InitEntityRepoDB
	return funcInitEntityRepoDB
}

func (self *EntityService) GetFuncInitEntityNotifyExcel() (funcInitEntityNotifyExcel typeInterface.FuncInitEntityNotifyExcel) {
	funcInitEntityNotifyExcel = notifyExcel.InitEntityNotifyExcel
	return funcInitEntityNotifyExcel
}

func (self *EntityService) GetFuncInitEntityNotifyEmail() (funcInitEntityNotifyEmail typeInterface.FuncInitEntityNotifyEmail) {
	funcInitEntityNotifyEmail = notifyEmail.InitEntityNotifyEmail
	return funcInitEntityNotifyEmail
}

func (self *EntityService) GetFuncInitEntityRepoCache() (funcInitEntityRepoCache typeInterface.FuncInitEntityRepoCache) {
	funcInitEntityRepoCache = repoCache.InitEntityRepoCache
	return funcInitEntityRepoCache
}
