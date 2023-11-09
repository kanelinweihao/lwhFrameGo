package main

import (
	"embed"
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/api/frontEnd"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/codeLine"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/ip"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/os"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/pack"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/rfl"
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
	finishIt()
}

func showTips() {
	ip.ShowIP()
	moduleTips := conf.GetModuleTips()
	time.ShowTimeAndMsg(moduleTips)
	return
}

func showWeb() {
	frontEnd.ExecFrontEnd()
	return
}

func finishIt() {
	fmt.Println("FINISH")
	time.Sleep(10, "s")
	dd.IGNORE("")
	dd.DIE("")
}
