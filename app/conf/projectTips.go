package conf

import (
	"fmt"
)

func GetProjectTips() (projectTips string) {
	projectDesc := GetProjectDesc()
	webTips := getWebTips()
	projectTips = fmt.Sprintf(
		"%s%s",
		projectDesc,
		webTips)
	return projectTips
}

func GetProjectDesc() (projectDesc string) {
	projectNameCN := GetProjectNameCN()
	msgProjectName := fmt.Sprintf(
		"项目名称 = %s\n",
		projectNameCN)
	projectVersion := GetProjectVersion()
	msgProjectVersion := fmt.Sprintf(
		"项目版本 = %s\n",
		projectVersion)
	versionDescription := GetVersionDescriptionNow()
	msgVersionDescription := fmt.Sprintf(
		"版本描述 = %s\n",
		versionDescription)
	projectDesc = fmt.Sprintf(
		"\n%s%s%s\n",
		msgProjectName,
		msgProjectVersion,
		msgVersionDescription)
	return projectDesc
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
