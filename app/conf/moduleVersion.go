package conf

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/pack"
	"strings"
)

var versionMajor = "1"
var versionMinor = "0"
var versionPatch = "2"

func GetModuleVersion() (modVersion string) {
	version := fmt.Sprintf(
		"%s.%s.%s",
		versionMajor,
		versionMinor,
		versionPatch)
	modVersion = "v" + version
	return modVersion
}

func GetParamsVersion() (paramsVersion base.AttrS1) {
	paramsVersion = make(base.AttrS1)
	pathVersion := getPathVersion()
	strFromVersion := pack.ReadFileEmbedAsString(pathVersion)
	sliceVersionLine := strings.Split(strFromVersion, "\n")
	strAnnotation := "#"
	strSeparator := "|"
	for _, strVersionLine := range sliceVersionLine {
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
	version := GetModuleVersion()
	versionDescription = GetVersionDescription(version)
	return versionDescription
}
