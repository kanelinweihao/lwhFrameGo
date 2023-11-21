package conf

func GetProjectNameEN() (projectNameEN string) {
	projectNameEN = getEnvValue("ProjectNameEN")
	return projectNameEN
}

func GetProjectNameCN() (projectNameCN string) {
	projectNameCN = getEnvValue("ProjectNameCN")
	return projectNameCN
}
