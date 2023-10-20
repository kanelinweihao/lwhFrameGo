package codeLine

import (
	_ "fmt"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/base"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/file"
	"os"
	"path/filepath"
)

var attrValidExt = base.AttrS1{
	"go": "",
}
var attrValidFilePath = base.AttrS1{}

func GetCodeLine(pathDirRel string) (countCodeLine int) {
	pathDirAbs := getPathDirAbs(pathDirRel)
	countCodeLine = getCodeLineValid(pathDirAbs)
	return countCodeLine
}

func getPathDirAbs(pathDirRel string) (pathDirAbs string) {
	if pathDirRel == "" {
		pathDirRel = "./"
	}
	pathDirAbs = file.GetFilePathAbs(pathDirRel)
	return pathDirAbs
}

func getCodeLineValid(pathDirAbs string) (countCodeLine int) {
	countCodeLine = 0
	filepath.Walk(pathDirAbs, walkFunc)
	// dd.DD(attrValidFilePath)
	countFile := len(attrValidFilePath)
	// dd.DD(countFile)
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
	// fmt.Println(path)
	isValidFile := isValidFile(path)
	if !isValidFile {
		return nil
	}
	attrValidFilePath[path] = ""
	// fmt.Println(path)
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
	// fmt.Println(pathFileValid)
	countCodeLineEach = 0
	arrStrCodeLine := file.ReadFileAsArrayString(pathFileValid)
	// dd.DD(arrStrCodeLine)
	for _, strCodeLine := range arrStrCodeLine {
		// fmt.Println(strCodeLine)
		if strCodeLine == "" {
			continue
		}
		countCodeLineEach++
		// fmt.Println(strCodeLine)
	}
	// fmt.Println(countCodeLineEach)
	return countCodeLineEach
}
