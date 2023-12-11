package conf

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
)

func GetDomain() (domain string) {
	host := getEnvValue("HtmlHost")
	port := getEnvValue("HtmlPort")
	domain = fmt.Sprintf(
		"%s:%s",
		host,
		port)
	return domain
}

func GetParamsTmpl(isReqValid bool) (paramsTmpl typeMap.AttrT1) {
	paramsKeyTmpl := typeMap.AttrS1{
		"ProjectTitle": "ProjectNameCN",
	}
	paramsTmpl = getParamsEnvNeed(paramsKeyTmpl)
	ProjectVersion := GetProjectVersion()
	paramsTmpl["ProjectVersion"] = ProjectVersion
	msgShow := "Ready"
	if isReqValid {
		msgShow = "Success"
	}
	paramsTmpl["MsgShow"] = msgShow
	return paramsTmpl
}
