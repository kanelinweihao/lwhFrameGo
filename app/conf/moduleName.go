package conf

import (
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
)

func GetModuleNameEN() (moduleNameEN string) {
	moduleNameEN = getEnvValue("ModuleNameEN")
	return moduleNameEN
}

func GetModuleNameCN() (moduleNameCN string) {
	moduleNameCN = getEnvValue("ModuleNameCN")
	return moduleNameCN
}
