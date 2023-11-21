package conf

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/pack"
	"strings"
)

func GetProjectVersion() (modVersion string) {
	paramsKeyVersion := base.AttrS1{
		"VersionMajor": "VersionMajor",
		"VersionMinor": "VersionMinor",
		"VersionPatch": "VersionPatch",
	}
	paramsVersion := getParamsEnvNeed(paramsKeyVersion)
	versionMajor := paramsVersion["VersionMajor"].(string)
	versionMinor := paramsVersion["VersionMinor"].(string)
	versionPatch := paramsVersion["VersionPatch"].(string)
	modVersion = fmt.Sprintf(
		"v%s.%s.%s",
		versionMajor,
		versionMinor,
		versionPatch)
	return modVersion
}

func GetParamsVersion() (paramsVersion base.AttrS1) {
	paramsVersion = make(base.AttrS1)
	pathVersion := getPathVersion()
	strFromEnv := pack.ReadFileEmbedAsString(pathVersion)
	arrLine := strings.Split(strFromEnv, "\n")
	strAnnotation := "#"
	strSeparator := "|"
	for _, strLine := range arrLine {
		if len(strLine) == 0 {
			continue
		}
		firstLetter := string(strLine[0])
		if firstLetter == strAnnotation {
			continue
		}
		arrPart := strings.Split(strLine, strSeparator)
		if len(arrPart) != 2 {
			continue
		}
		version := arrPart[0]
		versionDescription := arrPart[1]
		paramsVersion[version] = versionDescription
	}
	return paramsVersion
}

func GetVersionDescription(version string) (versionDescription string) {
	paramsVersion := GetParamsVersion()
	versionDescription, ok := paramsVersion[version]
	if !ok {
		msgError := fmt.Sprintf(
			"The version |%s| not found in version.env",
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
	version := GetProjectVersion()
	versionDescription = GetVersionDescription(version)
	return versionDescription
}
