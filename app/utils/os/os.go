package os

import (
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"runtime"
)

func GetOSName() (osName string) {
	osName = runtime.GOOS
	return osName
}

func GetOSArch() (osArch string) {
	osArch = runtime.GOARCH
	return osArch
}

func GetPathRoot() (pathRoot string) {
	pathRoot = runtime.GOROOT()
	return pathRoot
}

func getNumGoroutine() (numGoroutine int) {
	numGoroutine = runtime.NumGoroutine()
	return numGoroutine
}

func getNumCPU() (numCPU int) {
	numCPU = runtime.NumCPU()
	return numCPU
}

func EnbaleCPU() (numCPUEnabled int) {
	numCPU := getNumCPU()
	numCPUEnabled = runtime.GOMAXPROCS(numCPU)
	return numCPUEnabled
}
