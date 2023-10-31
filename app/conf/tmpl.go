package conf

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
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
		"ModTitle": "ModuleNameCN",
	}
	paramsTmpl = getParamsEnvNeed(paramsKeyTmpl)
	ModVersion := GetModuleVersion()
	paramsTmpl["ModVersion"] = ModVersion
	return paramsTmpl
}
