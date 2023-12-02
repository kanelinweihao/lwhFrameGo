package serviceGetMobileNoAndOrgIdByShortUserId

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/serviceUser/serviceGetMobileNoAndOrgIdByShortUserId/notify/notifyEmail"
	"github.com/kanelinweihao/lwhFrameGo/app/service/serviceUser/serviceGetMobileNoAndOrgIdByShortUserId/notify/notifyExcel"
	"github.com/kanelinweihao/lwhFrameGo/app/service/serviceUser/serviceGetMobileNoAndOrgIdByShortUserId/params"
	"github.com/kanelinweihao/lwhFrameGo/app/service/serviceUser/serviceGetMobileNoAndOrgIdByShortUserId/repo/repoCache"
	"github.com/kanelinweihao/lwhFrameGo/app/service/serviceUser/serviceGetMobileNoAndOrgIdByShortUserId/repo/repoDB"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

type EntityService struct {
	ParamsOut        base.AttrT1
	ServiceName      string
	ParamsIn         base.AttrT1
	EntityParams     base.TEntityParams
	AttrT3DBData     base.AttrT3
	ParamsAppend     base.AttrT1
	PathDirExcel     string
	ArrPathFileExcel []string
}

func (self *EntityService) Exec() (paramsOut base.AttrT1) {
	self.initParams().getDBData()
	self.putExcel().sendEmail().setCache()
	paramsOut = self.ParamsOut
	return paramsOut
}

func (self *EntityService) initParams() *EntityService {
	paramsIn := self.ParamsIn
	entityParams := params.InitEntityParams(paramsIn)
	self.EntityParams = entityParams
	paramsOut := entityParams.ToAttr()
	self.ParamsOut = paramsOut
	return self
}

func (self *EntityService) getDBData() *EntityService {
	paramsToRepo := self.ParamsOut
	entityRepoDB := repoDB.InitEntityRepoDB(paramsToRepo)
	attrT3DBData, paramsAppend := entityRepoDB.GetDBData()
	self.AttrT3DBData = attrT3DBData
	self.ParamsAppend = paramsAppend
	entityParams := self.EntityParams
	entityParams.SetPropertiesAppend(paramsAppend)
	self.EntityParams = entityParams
	paramsOut := entityParams.ToAttr()
	self.ParamsOut = paramsOut
	return self
}

func (self *EntityService) putExcel() *EntityService {
	paramsToNotify := self.ParamsOut
	attrT3DBData := self.AttrT3DBData
	entityNotifyExcel := notifyExcel.InitEntityNotifyExcel(paramsToNotify, attrT3DBData)
	_, pathDirExcel, arrPathFileExcel := entityNotifyExcel.PutExcel()
	self.PathDirExcel = pathDirExcel
	self.ArrPathFileExcel = arrPathFileExcel
	return self
}

func (self *EntityService) sendEmail() *EntityService {
	pathDirExcel := self.PathDirExcel
	arrPathFileExcel := self.ArrPathFileExcel
	entityNotifyEmail := notifyEmail.InitEntityNotifyEmail(pathDirExcel, arrPathFileExcel)
	entityNotifyEmail.SendEmail()
	return self
}

func (self *EntityService) setCache() *EntityService {
	arrPathFileExcel := self.ArrPathFileExcel
	entityRepoCache := repoCache.InitEntityRepoCache(arrPathFileExcel)
	entityRepoCache.SetCache()
	return self
}
