package serviceBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

type EntityServiceBase struct {
	Derived          typeStruct.EntityService
	ServiceName      string
	ParamsIn         typeMap.AttrT1
	EntityParams     typeStruct.EntityParams
	AttrT3DBData     typeMap.AttrT3
	ParamsAppend     typeMap.AttrT1
	PathDirExcel     string
	ArrPathFileExcel []string
	ParamsOut        typeMap.AttrT1
}

func (self *EntityServiceBase) Exec(paramsIn typeMap.AttrT1) (paramsOut typeMap.AttrT1) {
	self.ParamsIn = paramsIn
	self.Derived.InitParams().GetDBData().PutExcel().SendEmail().SetCache()
	paramsOut = self.ParamsOut
	return paramsOut
}

func (self *EntityServiceBase) InitParams() typeStruct.EntityService {
	paramsIn := self.ParamsIn
	funcInitEntityParams := self.Derived.GetFuncInitEntityParams()
	entityParams := funcInitEntityParams(paramsIn)
	self.EntityParams = entityParams
	paramsOut := entityParams.ToAttr()
	self.ParamsOut = paramsOut
	return self.Derived
}

func (self *EntityServiceBase) GetFuncInitEntityParams() (funcInitEntityParams typeStruct.FuncInitEntityParams) {
	funcInitEntityParams = nil
	return funcInitEntityParams
}

func (self *EntityServiceBase) GetDBData() typeStruct.EntityService {
	paramsToRepo := self.ParamsOut
	funcInitEntityRepoDB := self.Derived.GetFuncInitEntityRepoDB()
	entityRepoDB := funcInitEntityRepoDB(paramsToRepo)
	attrT3DBData, paramsAppend := entityRepoDB.GetDBData()
	self.AttrT3DBData = attrT3DBData
	self.ParamsAppend = paramsAppend
	entityParams := self.EntityParams
	entityParams.SetPropertiesAppend(paramsAppend)
	self.EntityParams = entityParams
	paramsOut := entityParams.ToAttr()
	self.ParamsOut = paramsOut
	return self.Derived
}

func (self *EntityServiceBase) GetFuncInitEntityRepoDB() (funcInitEntityRepoDB typeStruct.FuncInitEntityRepoDB) {
	funcInitEntityRepoDB = nil
	return funcInitEntityRepoDB
}

func (self *EntityServiceBase) PutExcel() typeStruct.EntityService {
	paramsToNotify := self.ParamsOut
	attrT3DBData := self.AttrT3DBData
	funcInitEntityNotifyExcel := self.Derived.GetFuncInitEntityNotifyExcel()
	entityNotifyExcel := funcInitEntityNotifyExcel(paramsToNotify, attrT3DBData)
	_, pathDirExcel, arrPathFileExcel := entityNotifyExcel.PutExcel()
	self.PathDirExcel = pathDirExcel
	self.ArrPathFileExcel = arrPathFileExcel
	return self.Derived
}

func (self *EntityServiceBase) GetFuncInitEntityNotifyExcel() (funcInitEntityNotifyExcel typeStruct.FuncInitEntityNotifyExcel) {
	funcInitEntityNotifyExcel = nil
	return funcInitEntityNotifyExcel
}

func (self *EntityServiceBase) SendEmail() typeStruct.EntityService {
	pathDirExcel := self.PathDirExcel
	arrPathFileExcel := self.ArrPathFileExcel
	funcInitEntityNotifyEmail := self.Derived.GetFuncInitEntityNotifyEmail()
	entityNotifyEmail := funcInitEntityNotifyEmail(pathDirExcel, arrPathFileExcel)
	entityNotifyEmail.SendEmail()
	return self.Derived
}

func (self *EntityServiceBase) GetFuncInitEntityNotifyEmail() (funcInitEntityNotifyEmail typeStruct.FuncInitEntityNotifyEmail) {
	funcInitEntityNotifyEmail = nil
	return funcInitEntityNotifyEmail
}

func (self *EntityServiceBase) SetCache() typeStruct.EntityService {
	arrPathFileExcel := self.ArrPathFileExcel
	funcInitEntityRepoCache := self.Derived.GetFuncInitEntityRepoCache()
	entityRepoCache := funcInitEntityRepoCache(arrPathFileExcel)
	entityRepoCache.SetCache()
	return self.Derived
}

func (self *EntityServiceBase) GetFuncInitEntityRepoCache() (funcInitEntityRepoCache typeStruct.FuncInitEntityRepoCache) {
	funcInitEntityRepoCache = nil
	return funcInitEntityRepoCache
}
