package serviceGetMobileNoAndOrgIdByShortUserId

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/user/serviceGetMobileNoAndOrgIdByShortUserId/notify/notifyEmail"
	"github.com/kanelinweihao/lwhFrameGo/app/service/user/serviceGetMobileNoAndOrgIdByShortUserId/notify/notifyExcel"
	"github.com/kanelinweihao/lwhFrameGo/app/service/user/serviceGetMobileNoAndOrgIdByShortUserId/params"
	"github.com/kanelinweihao/lwhFrameGo/app/service/user/serviceGetMobileNoAndOrgIdByShortUserId/repo/repoCache"
	"github.com/kanelinweihao/lwhFrameGo/app/service/user/serviceGetMobileNoAndOrgIdByShortUserId/repo/repoDB"
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
