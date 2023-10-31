package codeLine

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/file"
	"os"
	"path/filepath"
)

var attrValidExt = base.AttrS1{
	"go": "",
}
var attrValidFilePath = make(base.AttrS1)

func GetCodeLine(pathDirRel string) (countCodeLine int) {
	pathDirEmbed := getPathDirEmbed(pathDirRel)
	countCodeLine = getCodeLineValid(pathDirEmbed)
	return countCodeLine
}

func getPathDirEmbed(pathDirRel string) (pathDirEmbed string) {
	if pathDirRel == "" {
		pathDirRel = "./"
	}
	pathDirEmbed = file.GetFilePathEmbed(pathDirRel)
	pathDirEmbed = pathDirRel
	return pathDirEmbed
}

func getCodeLineValid(pathDirEmbed string) (countCodeLine int) {
	countCodeLine = 0
	filepath.Walk(pathDirEmbed, walkFunc)
	countFile := len(attrValidFilePath)
	if countFile == 0 {
		return 0
	}
	for pathFileValid, _ := range attrValidFilePath {
		countCodeLineEach := getCodeLineValidEach(pathFileValid)
		countCodeLine += countCodeLineEach
	}
	return countCodeLine
}

func walkFunc(path string, info os.FileInfo, err error) error {
	isValidFile := isValidFile(path)
	if !isValidFile {
		return nil
	}
	attrValidFilePath[path] = ""
	return nil
}

func isValidFile(path string) (isValidFile bool) {
	isValidFile = false
	isDir := file.IsDir(path)
	if isDir {
		return false
	}
	ext := file.GetExt(path)
	isValidExt := base.IsKeyOfAttrStr(ext, attrValidExt)
	if !isValidExt {
		return false
	}
	return true
}

func getCodeLineValidEach(pathFileValid string) (countCodeLineEach int) {
	countCodeLineEach = 0
	arrStrCodeLine := file.ReadFileAsArrayString(pathFileValid)
	for _, strCodeLine := range arrStrCodeLine {
		if strCodeLine == "" {
			continue
		}
		countCodeLineEach++
	}
	return countCodeLineEach
}
