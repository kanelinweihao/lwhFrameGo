package conf

func GetModuleNameEN() (moduleNameEN string) {
	moduleNameEN = getEnvValue("ModuleNameEN")
	return moduleNameEN
}

func GetModuleNameCN() (moduleNameCN string) {
	moduleNameCN = getEnvValue("ModuleNameCN")
	return moduleNameCN
}
