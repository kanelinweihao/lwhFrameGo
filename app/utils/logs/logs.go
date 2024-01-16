package logs

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/file"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/ip"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/times"
	"golang.org/x/exp/slog"
	"os"
	"strings"
)

func SetLog(msg string, logName string) {
	f := getFileOfLog(logName)
	defer closeFile(f)
	opts := getSlogHandlerOptions()
	h := slog.NewJSONHandler(f, opts)
	IPAddress := ip.GetIPOfServer()
	l := slog.New(h).With("ipServer", IPAddress)
	// slog.SetDefault(l)
	l.Info(msg)
	return
}

func getFileOfLog(logName string) (f *os.File) {
	pathDirLog := conf.GetPathLog()
	file.MakeDir(pathDirLog)
	projectNameEn := conf.GetProjectNameEN()
	projectNameLower := strings.ToLower(projectNameEn)
	timeSuffix := times.GetTimeSuffixDate()
	fileName := fmt.Sprintf(
		"log_%s_%s_%s.log",
		projectNameLower,
		logName,
		timeSuffix)
	pathLog := pathDirLog + "/" + fileName
	f, errFileAppend := os.OpenFile(pathLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	err.ErrCheck(errFileAppend)
	return f
}

func closeFile(f *os.File) {
	errFileClose := f.Close()
	err.ErrCheck(errFileClose)
	return
}

func getSlogHandlerOptions() (opts *slog.HandlerOptions) {
	opts = new(slog.HandlerOptions)
	opts.AddSource = false
	opts.Level = slog.LevelInfo
	return opts
}
