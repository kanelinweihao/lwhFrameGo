package frontEnd

import (
	"go.lwh.com/linweihao/lwhFrameGo/app/api/htmlSet"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/err"
)

func ExecFrontEnd() {
	defer err.ThrowError()
	htmlSet.SetHtml()
	return
}
