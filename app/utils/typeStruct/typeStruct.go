package typeStruct

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"net/http"
)

type EntitySection struct {
	FieldName     string `validate:"required,min=2"`
	FieldNameCN   string `validate:"required,min=2"`
	FieldRemark   string `validate:"required,min=2"`
	Value         string `validate:"omitempty"`
	InputType     string `validate:"required,oneof='text' 'password'"`
	DivDisplay    string `validate:"required,oneof='TRUE' 'FALSE'"`
	InputDisabled string `validate:"required,oneof='TRUE' 'FALSE'"`
	InputOnkeyup  string `validate:"required,oneof='' 'onlyNumber'"`
}

type EntityBase interface{}

type EntityDBData interface {
	EntityBase
}

type EntityController interface {
	Load(entityControllerBaseOrDerived EntityController) EntityController
	Init(routeName string) EntityController
	GetRouteType() (routeType int)
	GetParamsDefault() (paramsInDefault typeMap.AttrT1, paramsOutDefault typeMap.AttrT1)
	GetFuncService() (funcService FuncService)
	GetPathTmpl() (pathTmpl string)
	GetParamsOut() (paramsOut typeMap.AttrT1)
	GetArrEntitySectionIn() (arrEntitySectionIn []EntitySection)
	GetArrEntitySectionOut() (arrEntitySectionOut []EntitySection)
	SetArrEntitySectionIn(arrEntitySectionIn []EntitySection)
	SetArrEntitySectionOut(arrEntitySectionOut []EntitySection)
	GetParamsOutAppendWeb() (paramsOutAppendWeb typeMap.AttrT1)
	Close()
	Exec(resp http.ResponseWriter, req *http.Request)
}
type FuncService func(paramsIn typeMap.AttrT1) (paramsOut typeMap.AttrT1)

type EntityService1 interface {
	Load(entityServiceBaseOrDerived EntityService1) EntityService1
	Init(serviceName string) EntityService1
	Exec(paramsIn typeMap.AttrT1) (paramsOut typeMap.AttrT1)
	InitParams() EntityService1
	GetFuncInitEntityParams() (funcInitEntityParams FuncInitEntityParams)
	GetDBData() EntityService1
	GetFuncInitEntityRepoDB() (funcInitEntityRepoDB FuncInitEntityRepoDB)
	PutExcel() EntityService1
	GetFuncInitEntityNotifyExcel() (funcInitEntityNotifyExcel FuncInitEntityNotifyExcel)
	SendEmail() EntityService1
	GetFuncInitEntityNotifyEmail() (funcInitEntityNotifyEmail FuncInitEntityNotifyEmail)
	SetCache() EntityService1
	GetFuncInitEntityRepoCache() (funcInitEntityRepoCache FuncInitEntityRepoCache)
}
type EntityService2 interface {
	Load(entityServiceBaseOrDerived EntityService2) EntityService2
	Init(serviceName string) EntityService2
	Exec(paramsIn typeMap.AttrT1) (paramsOut typeMap.AttrT1)
	InitParams() EntityService2
	GetFuncInitEntityParams() (funcInitEntityParams FuncInitEntityParams)
}
type FuncInitEntityParams func(paramsIn typeMap.AttrT1) (entityParams EntityParams)
type FuncInitEntityRepoDB func(paramsRepoDB typeMap.AttrT1) (entityRepoDB EntityRepoDB)
type FuncInitEntityNotifyExcel func(paramsToNotify typeMap.AttrT1, attrT3DBData typeMap.AttrT3) (entityNotifyExcel EntityNotifyExcel)
type FuncInitEntityNotifyEmail func(pathDirExcel string, arrPathFileExcel []string) (entityNotifyEmail EntityNotifyEmail)
type FuncInitEntityRepoCache func(arrPathFileExcel []string) (entityRepoCache EntityRepoCache)

type EntityParams interface {
	Load(entityParamsBaseOrDerived EntityParams) EntityParams
	Init(paramsIn typeMap.AttrT1) EntityParams
	SetParamsExec() EntityParams
	Validate() EntityParams
	SetPropertiesAppend(paramsAppend typeMap.AttrT1) EntityParams
	ToAttr() (paramsOut typeMap.AttrT1)
}

type EntityRepoDB interface {
	Load(entityRepoDBBaseOrDerived EntityRepoDB) EntityRepoDB
	Init(paramsRepoDB typeMap.AttrT1) EntityRepoDB
	GetArrSQLName() (arrSQLName []string)
	GetAttrArgsForQuery() (attrArgsForQuery typeMap.AttrS1)
	GetDBData() (attrT3DBData typeMap.AttrT3, paramsAppend typeMap.AttrT1)
	GetPropertiesNeed() (paramsAppend typeMap.AttrT1)
	ToAttr() (paramsOut typeMap.AttrT1)
}

type EntityNotifyExcel interface {
	Load(entityNotifyExcelBaseOrDerived EntityNotifyExcel) EntityNotifyExcel
	Init(paramsToNotify typeMap.AttrT1, attrT3DBData typeMap.AttrT3) EntityNotifyExcel
	PutExcel() (attrS3ExcelData typeMap.AttrS3, pathDirExcel string, arrPathFileExcel []string)
	ToAttr() (paramsOut typeMap.AttrT1)
}

type EntityNotifyEmail interface {
	Load(entityNotifyEmailBaseOrDerived EntityNotifyEmail) EntityNotifyEmail
	Init(pathDirExcel string, arrPathFileExcel []string) EntityNotifyEmail
	SendEmail() (arrEmailSubject []string)
	ToAttr() (paramsOut typeMap.AttrT1)
}

type EntityRepoCache interface {
	Load(entityRepoCacheBaseOrDerived EntityRepoCache) EntityRepoCache
	Init(arrPathFileExcel []string) EntityRepoCache
	SetCache() (boxFromCache typeMap.BoxData)
	ToAttr() (paramsOut typeMap.AttrT1)
}
