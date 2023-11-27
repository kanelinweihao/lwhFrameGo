package service

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/serviceUser/serviceGetUserIdByMobileNoAndOrgId/notify/notifyEmail"
	"github.com/kanelinweihao/lwhFrameGo/app/service/serviceUser/serviceGetUserIdByMobileNoAndOrgId/notify/notifyExcel"
	"github.com/kanelinweihao/lwhFrameGo/app/service/serviceUser/serviceGetUserIdByMobileNoAndOrgId/params"
	"github.com/kanelinweihao/lwhFrameGo/app/service/serviceUser/serviceGetUserIdByMobileNoAndOrgId/repo/repoCache"
	"github.com/kanelinweihao/lwhFrameGo/app/service/serviceUser/serviceGetUserIdByMobileNoAndOrgId/repo/repoDB"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

type EntityService struct {
	ParamsOut         base.AttrT1
	ParamsIn          base.AttrT1
	EntityParams      *params.EntityParams
	ParamsToRepo      base.AttrT1
	EntityRepoDB      *repoDB.EntityRepoDB
	AttrT3DBData      base.AttrT3
	ParamsToNotify    base.AttrT1
	EntityNotifyExcel *notifyExcel.EntityNotifyExcel
	PathDirExcel      string
	ArrPathFileExcel  []string
	EntityNotifyEmail *notifyEmail.EntityNotifyEmail
	EntityRepoCache   *repoCache.EntityRepoCache
}

func (self *EntityService) Exec() (paramsOut base.AttrT1) {
	// dd.DD(666)
	self.initParamsByEntityParams()
	self.getDBDataByEntityRepoDB()
	self.putExcelByEntityNotifyExcel()
	self.sendEmailByEntityNotifyEmail()
	self.setCacheByEntityRepoCache()
	paramsOut = self.getParamsOut()
	return paramsOut
}

func (self *EntityService) initParamsByEntityParams() *EntityService {
	paramsIn := self.ParamsIn
	entityParams := params.MakeEntityParams(paramsIn)
	self.EntityParams = entityParams
	paramsToRepo := entityParams.GetParams()
	self.ParamsToRepo = paramsToRepo
	// time.ShowTimeAndMsg("Params init success")
	return self
}

func (self *EntityService) getDBDataByEntityRepoDB() *EntityService {
	paramsToRepo := self.ParamsToRepo
	entityRepoDB := repoDB.MakeEntityRepoDB(paramsToRepo)
	self.EntityRepoDB = entityRepoDB
	attrT3DBData, userId := entityRepoDB.GetDBData()
	self.AttrT3DBData = attrT3DBData
	entityParams := self.EntityParams
	entityParams.UserId = userId
	self.EntityParams = entityParams
	paramsToNotify := entityParams.GetParams()
	self.ParamsToNotify = paramsToNotify
	// time.ShowTimeAndMsg("DBData get success")
	return self
}

func (self *EntityService) putExcelByEntityNotifyExcel() *EntityService {
	paramsToNotify := self.ParamsToNotify
	attrT3DBData := self.AttrT3DBData
	entityNotifyExcel := notifyExcel.MakeEntityNotifyExcel(paramsToNotify, attrT3DBData)
	self.EntityNotifyExcel = entityNotifyExcel
	_, pathDirExcel, arrPathFileExcel := entityNotifyExcel.PutExcel()
	self.PathDirExcel = pathDirExcel
	self.ArrPathFileExcel = arrPathFileExcel
	// time.ShowTimeAndMsg("Excel put success")
	return self
}

func (self *EntityService) sendEmailByEntityNotifyEmail() *EntityService {
	pathDirExcel := self.PathDirExcel
	arrPathFileExcel := self.ArrPathFileExcel
	entityNotifyEmail := notifyEmail.MakeEntityNotifyEmail(pathDirExcel, arrPathFileExcel)
	self.EntityNotifyEmail = entityNotifyEmail
	entityNotifyEmail.SendEmail()
	// time.ShowTimeAndMsg("Email send success")
	return self
}

func (self *EntityService) setCacheByEntityRepoCache() *EntityService {
	arrPathFileExcel := self.ArrPathFileExcel
	entityRepoCache := repoCache.MakeEntityRepoCache(arrPathFileExcel)
	self.EntityRepoCache = entityRepoCache
	entityRepoCache.SetCache()
	// time.ShowTimeAndMsg("Cache set success")
	return self
}

func (self *EntityService) getParamsOut() (paramsOut base.AttrT1) {
	entityParams := self.EntityParams
	paramsOut = entityParams.SetMsgShow().GetParams()
	self.ParamsOut = paramsOut
	// time.ShowTimeAndMsg("ParamsOut get success")
	return paramsOut
}
