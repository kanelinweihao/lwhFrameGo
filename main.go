package main

import (
	"embed"
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/boot"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/ga"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/os"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/pack"
)

//go:embed res/*
var FilesResource embed.FS

func init() {
	// error
	fmt.Println()
	defer err.ThrowError()
	// os
	os.EnbaleCPU()
	// embed
	pack.SetFilesResource(FilesResource)
	return
}

func main() {
	// test()
	boot.Boot()
	return
}

func test() {
	// TODO
	// secret := ga.GetSecret()
	secret := "THPVAWMJ7J6RADNDPTLI7XVT3LGCPLX5"
	code, errGetCode := ga.GetCode(secret)
	err.ErrCheck(errGetCode)
	ok, errVerifyCode := ga.VerifyCode(secret, code)
	err.ErrCheck(errVerifyCode)
	dd.DDD(secret, code, ok)
}
