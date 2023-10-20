package browser

import (
	"fmt"
	"go.lwh.com/linweihao/lwhFrameGo/app/conf"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/base"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/err"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/os"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/time"
	"os/exec"
)

var commands = base.AttrS1{
	"windows": "cmd /c start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

////
// Browser
////

func OpenBrowser() {
	osName := os.GetOSName()
	// dd.DD(osName)
	if osName != "windows" {
		return
	}
	command, _ := commands[osName]
	domain := conf.GetDomain()
	url := fmt.Sprintf(
		"%s://%s/%s",
		"http",
		domain,
		"index")
	run := fmt.Sprintf(
		"%s %s",
		command,
		url)
	// dd.DD(run)
	cmd := exec.Command(run)
	byteOutput, errCmd := cmd.Output()
	fmt.Println(byteOutput)
	err.ErrCheck(errCmd)
	msg := "Open Browser Success"
	time.ShowTimeAndMsg(msg)
	return
}
