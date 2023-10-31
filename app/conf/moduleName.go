package conf

import (
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
)

func GetModuleNameEN() (moduleNameEN string) {
	moduleNameEN = getEnvValue("ModuleNameEN")
	return moduleNameEN
}

func GetModuleNameCN() (moduleNameCN string) {
	moduleNameCN = getEnvValue("ModuleNameCN")
	return moduleNameCN
}
