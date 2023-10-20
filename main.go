package main

import (
	"embed"
	"fmt"
	"go.lwh.com/linweihao/lwhFrameGo/app/api/frontEnd"
	"go.lwh.com/linweihao/lwhFrameGo/app/conf"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/base"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/err"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/os"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/pack"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/rfl"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/time"
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
	doIt()
	return
}

func doIt() {
	showTips()
	showWeb()
}

func showTips() {
	now := time.GetNow()
	moduleTips := conf.GetModuleTips()
	strShow := now + "\n" + moduleTips
	fmt.Println(strShow)
	return
}

func showWeb() {
	frontEnd.ExecFrontEnd()
	return
}
