package time

import (
	"fmt"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"time"
)

var (
	t             time.Time
	formateBase   string = "2006-01-02 15:04:05"
	formateSuffix string = "20060102150405"
	now           string
	suffix        string
)

func getLocal() (localCn *time.Location) {
	localCn, errTime := time.LoadLocation("Asia/Shanghai")
	err.ErrCheck(errTime)
	return localCn
}

func GetNow() (now string) {
	localCn := getLocal()
	now = time.Now().UTC().In(localCn).Format(formateBase)
	return now
}

func ShowTime() {
	now = GetNow()
	fmt.Println(now)
	return
}

func ShowTimeAndMsg(msg string) {
	now = GetNow()
	msgFull := now + "    " + msg + "\n"
	fmt.Println(msgFull)
	return
}

func GetTimeSuffix() (suffix string) {
	localCn := getLocal()
	suffix = time.Now().UTC().In(localCn).Format(formateSuffix)
	return suffix
}

func Sleep(sleepNum int, sleepUnit string) {
	timeUnit := time.Millisecond
	switch sleepUnit {
	case "ns":
		timeUnit = time.Nanosecond
	case "us":
		timeUnit = time.Microsecond
	case "ms":
		timeUnit = time.Millisecond
	case "s":
		timeUnit = time.Second
	case "m":
		timeUnit = time.Minute
	case "h":
		timeUnit = time.Hour
	default:
		msgError := fmt.Sprintf(
			"The sleepUnit |%s| is invalid of |%s|",
			sleepUnit,
			"timeSleep")
		err.ErrPanic(msgError)
	}
	timeNum := time.Duration(sleepNum)
	d := timeUnit * timeNum
	time.Sleep(d)
	return
}
