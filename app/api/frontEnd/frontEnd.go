package frontEnd

import (
	"github.com/kanelinweihao/lwhFrameGo/app/api/htmlSet"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
)

func ExecFrontEnd() {
	htmlSet.SetHtml()
	return
}
