package conf

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
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

func GetParamsTmpl() (paramsTmpl base.AttrT1) {
	paramsKeyTmpl := base.AttrS1{
		"ProjectTitle": "ProjectNameCN",
	}
	paramsTmpl = getParamsEnvNeed(paramsKeyTmpl)
	ProjectVersion := GetProjectVersion()
	paramsTmpl["ProjectVersion"] = ProjectVersion
	return paramsTmpl
}
