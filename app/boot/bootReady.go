package boot

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/ip"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/times"
)

func ready() {
	showTips()
	return
}

func showTips() {
	ip.ShowIP()
	projectTips := conf.GetProjectTips()
	times.ShowTimeAndMsg(projectTips)
	return
}
