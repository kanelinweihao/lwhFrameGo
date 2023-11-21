package main

import (
	"embed"
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/api/frontEnd"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/ip"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/os"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/pack"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/time"
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
	ip.ShowIP()
	projectTips := conf.GetProjectTips()
	time.ShowTimeAndMsg(projectTips)
	return
}

func showWeb() {
	frontEnd.ExecFrontEnd()
	return
}
