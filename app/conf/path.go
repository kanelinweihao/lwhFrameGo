package conf

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/file"
)

var pathEnv string = "./res/env/env.env"
var pathVersion string = "./res/version/version.env"

func getPathEnv() (pathEnvEmbed string) {
	pathEnvEmbed = file.GetFilePathEmbed(pathEnv)
	return pathEnvEmbed
}

func getPathVersion() (pathVersionEmbed string) {
	pathVersionEmbed = file.GetFilePathEmbed(pathVersion)
	return pathVersionEmbed
}

func getPathFromEnv(envKey string) (pathEmbed string) {
	pathRel := getEnvValue(envKey)
	pathEmbed = file.GetFilePathEmbed(pathRel)
	return pathEmbed
}

func GetPathSQL() (pathSQL string) {
	pathSQL = getPathFromEnv("PathSQL")
	return pathSQL
}

func GetPathFavicon() (pathFavicon string) {
	pathFavicon = getPathFromEnv("PathFavicon")
	return pathFavicon
}

func GetPathLog() (pathLog string) {
	pathLog = getPathFromEnv("PathLog")
	return pathLog
}

func GetPathViewHeader() (pathViewHeader string) {
	pathViewHeader = getPathFromEnv("PathViewHeader")
	pathViewHeader = file.GetFilePathEmbed(pathViewHeader)
	return pathViewHeader
}

func GetPathViewFooter() (pathViewFooter string) {
	pathViewFooter = getPathFromEnv("PathViewFooter")
	pathViewFooter = file.GetFilePathEmbed(pathViewFooter)
	return pathViewFooter
}

func GetPathPrivateKey() (pathPrivateKey string) {
	pathPrivateKey = getPathFromEnv("PathPrivateKey")
	return pathPrivateKey
}

func GetPathDirPutExcel() (pathDirPutExcel string) {
	pathDirPutExcel = getPathFromEnv("PathDirPutExcel")
	return pathDirPutExcel
}
