package conf

import (
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/file"
)

func GetPathPrivateKey() (pathPrivateKey string) {
	PathPrivateKey := getEnvValue("PathPrivateKey")
	pathPrivateKey = file.GetFilePathAbs(PathPrivateKey)
	// dd.DD(pathPrivateKey)
	return pathPrivateKey
}

func GetPathDirPutExcel() (pathDirPutExcel string) {
	PathDirPutExcel := getEnvValue("PathDirPutExcel")
	pathDirPutExcel = file.GetFilePathAbs(PathDirPutExcel)
	// dd.DD(pathDirPutExcel)
	return pathDirPutExcel
}
