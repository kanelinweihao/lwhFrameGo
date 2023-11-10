package conf

import (
	"fmt"
)

func GetModuleTips() (moduleTips string) {
	moduleDesc := GetModuleDesc()
	webTips := getWebTips()
	moduleTips = fmt.Sprintf(
		"%s%s",
		moduleDesc,
		webTips)
	return moduleTips
}

func GetModuleDesc() (moduleDesc string) {
	moduleNameCN := GetModuleNameCN()
	msgModuleName := fmt.Sprintf(
		"项目名称 = %s\n",
		moduleNameCN)
	moduleVersion := GetModuleVersion()
	msgModuleVersion := fmt.Sprintf(
		"项目版本 = %s\n",
		moduleVersion)
	versionDescription := GetVersionDescriptionNow()
	msgVersionDescription := fmt.Sprintf(
		"版本描述 = %s\n",
		versionDescription)
	moduleDesc = fmt.Sprintf(
		"\n%s%s%s\n",
		msgModuleName,
		msgModuleVersion,
		msgVersionDescription)
	return moduleDesc
}

func getWebTips() (webTips string) {
	// strFormatEN := "Ready!Please open |%s| in your browser to continue"
	strFormatCN := "准备就绪!请在浏览器中打开|%s|继续操作"
	address := GetDomain()
	webTips = fmt.Sprintf(
		strFormatCN,
		address)
	return webTips
}
