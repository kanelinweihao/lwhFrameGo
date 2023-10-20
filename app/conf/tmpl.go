package conf

import (
	"fmt"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/base"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
)

func GetDomain() (domain string) {
	host := getEnvValue("HtmlHost")
	port := getEnvValue("HtmlPort")
	domain = fmt.Sprintf(
		"%s:%s",
		host,
		port)
	// dd.DD(domain)
	return domain
}

func GetParamsTmpl() (paramsTmpl base.AttrT1) {
	// ModTitle := getEnvValue("ModuleNameCN")
	// paramsTmpl = base.AttrT1{
	// 	"ModTitle": ModTitle,
	// }
	paramsKeyTmpl := base.AttrS1{
		"ModTitle": "ModuleNameCN",
	}
	paramsTmpl = getParamsEnvNeed(paramsKeyTmpl)
	ModVersion := GetVersion()
	paramsTmpl["ModVersion"] = ModVersion
	// dd.DD(paramsTmpl)
	return paramsTmpl
}
