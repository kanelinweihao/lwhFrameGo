package serviceGetUserIdByMobileNoAndOrgId

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/user/serviceGetUserIdByMobileNoAndOrgId/notify/notifyEmail"
	"github.com/kanelinweihao/lwhFrameGo/app/service/user/serviceGetUserIdByMobileNoAndOrgId/notify/notifyExcel"
	"github.com/kanelinweihao/lwhFrameGo/app/service/user/serviceGetUserIdByMobileNoAndOrgId/params"
	"github.com/kanelinweihao/lwhFrameGo/app/service/user/serviceGetUserIdByMobileNoAndOrgId/repo/repoCache"
	"github.com/kanelinweihao/lwhFrameGo/app/service/user/serviceGetUserIdByMobileNoAndOrgId/repo/repoDB"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

type EntityService struct {
	typeStruct.EntityService1
}

func (self *EntityService) GetFuncInitEntityParams() (funcInitEntityParams typeStruct.FuncInitEntityParams) {
	funcInitEntityParams = params.InitEntityParams
	return funcInitEntityParams
}

func (self *EntityService) GetFuncInitEntityRepoDB() (funcInitEntityRepoDB typeStruct.FuncInitEntityRepoDB) {
	funcInitEntityRepoDB = repoDB.InitEntityRepoDB
	return funcInitEntityRepoDB
}

func (self *EntityService) GetFuncInitEntityNotifyExcel() (funcInitEntityNotifyExcel typeStruct.FuncInitEntityNotifyExcel) {
	funcInitEntityNotifyExcel = notifyExcel.InitEntityNotifyExcel
	return funcInitEntityNotifyExcel
}

func (self *EntityService) GetFuncInitEntityNotifyEmail() (funcInitEntityNotifyEmail typeStruct.FuncInitEntityNotifyEmail) {
	funcInitEntityNotifyEmail = notifyEmail.InitEntityNotifyEmail
	return funcInitEntityNotifyEmail
}

func (self *EntityService) GetFuncInitEntityRepoCache() (funcInitEntityRepoCache typeStruct.FuncInitEntityRepoCache) {
	funcInitEntityRepoCache = repoCache.InitEntityRepoCache
	return funcInitEntityRepoCache
}
