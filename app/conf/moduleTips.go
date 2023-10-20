package conf

import (
	"fmt"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/codeLine"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
)

func GetModuleTips() (moduleTips string) {
	moduleNameCN := GetModuleNameCN()
	msgModuleName := fmt.Sprintf(
		"项目名称 = %s\n",
		moduleNameCN)
	moduleVersion := GetVersion()
	msgModuleVersion := fmt.Sprintf(
		"项目版本 = %s\n",
		moduleVersion)
	versionDescription := GetVersionDescriptionNow()
	msgVersionDescription := fmt.Sprintf(
		"版本描述 = %s\n",
		versionDescription)
	pathDirRel := "./app/"
	numCodeLine := codeLine.GetCodeLine(pathDirRel)
	msgNumCodeLine := fmt.Sprintf(
		"项目代码行数 = %d\n",
		numCodeLine)
	moduleTips = fmt.Sprintf(
		"\n%s\n%s\n%s\n%s\n",
		msgModuleName,
		msgModuleVersion,
		msgVersionDescription,
		msgNumCodeLine)
	// dd.DD(moduleTips)
	return moduleTips
}
