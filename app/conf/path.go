package conf

import (
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/file"
)

var pathEnv string = "./res/env/env.env"
var pathVersion string = "./res/version/version.env"
var pathSQL string = "./res/sql/sql.env"

/*
PathSelf
*/

func getPathEnv() (pathEnvEmbed string) {
	pathEnvEmbed = file.GetFilePathEmbed(pathEnv)
	return pathEnvEmbed
}

func getPathVersion() (pathVersionEmbed string) {
	pathVersionEmbed = file.GetFilePathEmbed(pathVersion)
	return pathVersionEmbed
}

func getPathSQL() (pathSQLEmbed string) {
	pathSQLEmbed = file.GetFilePathEmbed(pathSQL)
	return pathSQLEmbed
}

/*
PathFromEnv
*/

func getPathFromEnv(envKey string) (pathEmbed string) {
	pathRel := getEnvValue(envKey)
	pathEmbed = file.GetFilePathEmbed(pathRel)
	return pathEmbed
}

func GetPathPrivateKey() (pathPrivateKey string) {
	pathPrivateKey = getPathFromEnv("PathPrivateKey")
	return pathPrivateKey
}

func GetPathDirPutExcel() (pathDirPutExcel string) {
	pathDirPutExcel = getPathFromEnv("PathDirPutExcel")
	return pathDirPutExcel
}
