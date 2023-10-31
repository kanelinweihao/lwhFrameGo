package file

import (
	_ "fmt"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	_ "io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

/*
Path
*/

func isFilePathAbs(filePath string) (isAbs bool) {
	isAbs = filepath.IsAbs(filePath)
	return isAbs
}

// 相对路径转绝对路径(实现对windows\分隔符的兼容)
func GetFilePathAbs(filePath string) (filePathAbs string) {
	isAbs := isFilePathAbs(filePath)
	if isAbs {
		filePathAbs = filePath
		return filePathAbs
	}
	filePathAbs, errAbs := filepath.Abs(filePath)
	err.ErrCheck(errAbs)
	return filePathAbs
}

// 绝对路径转相对路径(实现对embed静态文件打包的兼容)
func GetFilePathRel(filePath string) (filePathRel string) {
	isAbs := isFilePathAbs(filePath)
	if !isAbs {
		filePath = GetFilePathAbs(filePath)
	}
	filePathBase, errBase := filepath.Abs("./")
	err.ErrCheck(errBase)
	filePathRel, errRel := filepath.Rel(filePathBase, filePath)
	err.ErrCheck(errRel)
	return filePathRel
}

// 文件路径强制用/作分隔符(实现对embed静态文件打包的兼容)
func GetFilePathEmbed(filePath string) (filePathEmbed string) {
	filePathRel := GetFilePathRel(filePath)
	filePathEmbed = filepath.ToSlash(filePathRel)
	return filePathEmbed
}

/*
FileInfo
*/

func getFileInfo(filePath string) (fileInfo os.FileInfo, errFileInfo error) {
	fileInfo, errFileInfo = os.Stat(filePath)
	return fileInfo, errFileInfo
}

func GetExt(filePath string) (ext string) {
	// kDot := strings.LastIndex(fileName, ".")
	// kBegin := kDot + 1
	// ext = fileName[kBegin:]
	extWithDot := filepath.Ext(filePath)
	if len(extWithDot) == 0 {
		return ""
	}
	ext = extWithDot[1:]
	return ext
}

/*
Dir
*/

func IsDir(filePath string) (isDir bool) {
	fileInfo, errFileInfo := getFileInfo(filePath)
	err.ErrCheck(errFileInfo)
	isDir = fileInfo.IsDir()
	return isDir
}

func HasDir(filePath string) (isExisted bool, err error) {
	_, err = getFileInfo(filePath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func MakeDir(filePathDir string) {
	isExisted, errDirExisted := HasDir(filePathDir)
	err.ErrCheck(errDirExisted)
	if isExisted {
		// msgError := fmt.Sprintf(
		//     "The dir of |%s| is existed",
		//     filePathDir)
		// err.ErrPanic(msgError)
		return
	}
	errDirMk := os.MkdirAll(filePathDir, os.ModePerm)
	err.ErrCheck(errDirMk)
	return
}

/*
FileRead
*/

// --> pack.ReadFileEmbedAsArrayByte(path)
func readFileAsArrayByte(filePath string) (arrByte []byte) {
	fs, errFileOpen := os.OpenFile(
		filePath,
		os.O_RDONLY,
		0666)
	err.ErrCheck(errFileOpen)
	defer fs.Close()
	arrByte, errFileRead := ioutil.ReadAll(fs)
	err.ErrCheck(errFileRead)
	return arrByte
}

// --> pack.ReadFileEmbedAsString(path)
func readFileAsString(filePath string) (str string) {
	arrByte := readFileAsArrayByte(filePath)
	str = string(arrByte)
	return str
}

// --> pack.ReadFileEmbedAsArrayString(path)
func ReadFileAsArrayString(filePath string) (arrStr []string) {
	str := readFileAsString(filePath)
	arrStr = strings.Split(str, "\n")
	return arrStr
}
