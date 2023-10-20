package pack

import (
	"embed"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/err"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/file"
)

var FilesResource embed.FS

func SetFilesResource(FilesResourceFromEmbed embed.FS) {
	FilesResource = FilesResourceFromEmbed
	return
}

func GetArrayBytePrivateKey(pathPrivateKey string) (arrBytePrivateKey []byte) {
	// dd.DD(pathPrivateKey)
	// pathPrivateKeyRel := file.GetFilePathRel(pathPrivateKey)
	// dd.DD(pathPrivateKeyRel)
	pathPrivateKeyEmbed := file.GetFilePathEmbed(pathPrivateKey)
	// dd.DD(pathPrivateKeyEmbed)
	arrBytePrivateKey, errReadPrivateKey := FilesResource.ReadFile(pathPrivateKeyEmbed)
	err.ErrCheck(errReadPrivateKey)
	return arrBytePrivateKey
}

func GetFSOfTmpl(pathTmpl string) (fsResource embed.FS, patternsTmpl string) {
	// dd.DD(pathTmpl)
	// pathTmplRel := file.GetFilePathRel(pathTmpl)
	// dd.DD(pathTmplRel)
	fsResource = FilesResource
	patternsTmpl = file.GetFilePathEmbed(pathTmpl)
	return fsResource, patternsTmpl
	// var fsTmpl fs.File
	// fsTmpl,errFsOpen := FilesResource.Open(pathTmplRel)
	// err.ErrCheck(errFsOpen)
	// // dd.DD(fsTmpl)
	// return fsTmpl
}
