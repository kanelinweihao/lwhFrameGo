package os

import (
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
	"runtime"
)

func GetOSName() (osName string) {
	osName = runtime.GOOS
	// dd.DD(osName)
	// osArch := runtime.GOARCH
	// dd.DD(osArch)
	// pathRoot := runtime.GOROOT()
	// dd.DD(pathRoot)
	return osName
}

func EnbaleCPU() {
	// numGoroutine := runtime.NumGoroutine()
	// dd.DD(numGoroutine)
	numCPU := runtime.NumCPU()
	// dd.DD(numCPU)
	runtime.GOMAXPROCS(numCPU)
	// numCPUEnabled := runtime.GOMAXPROCS(numCPU)
	// dd.DD(numCPUEnabled)
	return
}
