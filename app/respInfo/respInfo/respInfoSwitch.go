package respInfo

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/respInfo/respInfoProjectInfo"
	"github.com/kanelinweihao/lwhFrameGo/app/respInfo/respInfoUser"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

func getRespInfoOfProjectInfo() (paramsInDefault base.AttrT1, paramsOutDefault base.AttrT1, funcService conf.FuncService, pathTmpl string) {
	paramsInDefault = respInfoProjectInfo.ParamsInDefaultProjectInfo
	paramsOutDefault = getParamsOutDefaultForWeb(respInfoProjectInfo.ParamsOutDefaultProjectInfo)
	funcService = respInfoProjectInfo.FuncServiceProjectInfo
	pathTmpl = respInfoProjectInfo.PathTmplProjectInfo
	return paramsInDefault, paramsOutDefault, funcService, pathTmpl
}

func getRespInfoOfGetMobileNoAndOrgIdByShortUserId() (paramsInDefault base.AttrT1, paramsOutDefault base.AttrT1, funcService conf.FuncService, pathTmpl string) {
	paramsInDefault = respInfoUser.ParamsInDefaultGetMobileNoAndOrgIdByShortUserId
	paramsOutDefault = getParamsOutDefaultForWeb(respInfoUser.ParamsOutDefaultGetMobileNoAndOrgIdByShortUserId)
	funcService = respInfoUser.FuncServiceGetMobileNoAndOrgIdByShortUserId
	pathTmpl = respInfoUser.PathTmplGetMobileNoAndOrgIdByShortUserId
	return paramsInDefault, paramsOutDefault, funcService, pathTmpl
}

func getRespInfoOfGetUserIdByMobileNoAndOrgId() (paramsInDefault base.AttrT1, paramsOutDefault base.AttrT1, funcService conf.FuncService, pathTmpl string) {
	paramsInDefault = respInfoUser.ParamsInDefaultGetUserIdByMobileNoAndOrgId
	paramsOutDefault = getParamsOutDefaultForWeb(respInfoUser.ParamsOutDefaultGetUserIdByMobileNoAndOrgId)
	funcService = respInfoUser.FuncServiceGetUserIdByMobileNoAndOrgId
	pathTmpl = respInfoUser.PathTmplGetUserIdByMobileNoAndOrgId
	return paramsInDefault, paramsOutDefault, funcService, pathTmpl
}
