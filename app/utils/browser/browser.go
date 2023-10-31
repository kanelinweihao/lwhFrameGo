package browser

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/os"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/time"
	"os/exec"
)

var commands = base.AttrS1{
	"windows": "cmd /c start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func OpenBrowser() {
	osName := os.GetOSName()
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
	cmd := exec.Command(run)
	byteOutput, errCmd := cmd.Output()
	err.ErrCheck(errCmd)
	msg := "Open Browser Success"
	time.ShowTimeAndMsg(msg)
	return
}
