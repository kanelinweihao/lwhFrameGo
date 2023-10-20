package conf

import (
	"fmt"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/base"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/err"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/file"
	"strings"
)

var Version string = "v1.0.1"
var PathVersion string = "./res/env/version.env"

func GetVersion() (version string) {
	version = Version
	return version
}

func getPathVersion() (pathVersion string) {
	pathVersion = file.GetFilePathAbs(PathVersion)
	// dd.DD(pathVersion)
	return pathVersion
}

func GetParamsVersion() (paramsVersion base.AttrS1) {
	paramsVersion = base.AttrS1{}
	pathVersion := getPathVersion()
	strFromVersion := file.ReadFileAsString(pathVersion)
	// dd.DD(strFromVersion)
	sliceVersionLine := strings.Split(strFromVersion, "\n")
	// dd.DD(sliceVersionLine)
	strAnnotation := "#"
	strSeparator := "|"
	for _, strVersionLine := range sliceVersionLine {
		// dd.DD(strVersionLine)
		if len(strVersionLine) == 0 {
			continue
		}
		firstLetter := string(strVersionLine[0])
		if firstLetter == strAnnotation {
			continue
		}
		sliceVersionKV := strings.Split(strVersionLine, strSeparator)
		if len(sliceVersionKV) != 2 {
			continue
		}
		paramsVersion[sliceVersionKV[0]] = sliceVersionKV[1]
	}
	return paramsVersion
}

func GetVersionDescription(version string) (versionDescription string) {
	paramsVersion := GetParamsVersion()
	// dd.DD(paramsVersion)
	versionDescription, ok := paramsVersion[version]
	if !ok {
		msgError := fmt.Sprintf(
			"The |%s| not found in version.env",
			version)
		err.ErrPanic(msgError)
	}
	if len(versionDescription) == 0 {
		msgError := fmt.Sprintf(
			"The value of |%s| is empty in version.env",
			version)
		err.ErrPanic(msgError)
	}
	return versionDescription
}

func GetVersionDescriptionNow() (versionDescription string) {
	version := GetVersion()
	versionDescription = GetVersionDescription(version)
	return versionDescription
}
