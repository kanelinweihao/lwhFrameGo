package pack

import (
	"embed"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/file"
	"strings"
)

var FilesResource embed.FS

func SetFilesResource(FilesResourceFromEmbed embed.FS) {
	FilesResource = FilesResourceFromEmbed
	return
}

func GetFSAndPath(path string) (fsResource embed.FS, pathEmbed string) {
	fsResource = FilesResource
	pathEmbed = file.GetFilePathEmbed(path)
	return fsResource, pathEmbed
}

func ReadFileEmbedAsArrayByte(path string) (arrByte []byte) {
	fsResource, pathEmbed := GetFSAndPath(path)
	arrByte, errRead := fsResource.ReadFile(pathEmbed)
	err.ErrCheck(errRead)
	return arrByte
}

func ReadFileEmbedAsString(path string) (str string) {
	arrByte := ReadFileEmbedAsArrayByte(path)
	str = string(arrByte)
	return str
}

func ReadFileEmbedAsArrayString(path string) (arrStr []string) {
	str := ReadFileEmbedAsString(path)
	arrStr = strings.Split(str, "\n")
	return arrStr
}
