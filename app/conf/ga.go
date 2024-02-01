package conf

func GetGaSecret() (gaSecret string) {
	gaSecret = getEnvValue("GaSecret")
	return gaSecret
}
