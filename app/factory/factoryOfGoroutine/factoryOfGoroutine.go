package factoryOfGoroutine

import (
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/goroutine"
)

func MakeEntityOfGoroutine() (entityChannel *goroutine.EntityChannel) {
	entityChannel = goroutine.InitEntityChannel()
	return entityChannel
}
