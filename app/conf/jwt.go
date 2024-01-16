package conf

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/ip"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/times"
	"net/http"
	"time"
)

func GetJwtSecret() (jwtSecret []byte) {
	strJwtSecret := getPathFromEnv("JwtSecret")
	jwtSecret = []byte(strJwtSecret)
	return jwtSecret
}

func GetJwtHoursExpires() (h time.Duration) {
	strHours := getPathFromEnv("JwtHoursExpires")
	intHours := conv.ToIntFromStr(strHours)
	h = time.Duration(intHours)
	return h
}

func GetJwtIssuer() (iss string) {
	ipServer := ip.GetIPOfServer()
	iss = ipServer
	return iss
}

func GetJwtSubject() (sub string) {
	sub = GetProjectNameEN()
	return sub
}

func GetJwtAudience(req *http.Request) (aud []string) {
	ipClient := ip.GetIPOfClient(req)
	aud = make([]string, 1)
	aud = append(aud, ipClient)
	return aud
}

func GetJwtID() (jti string) {
	suffix := times.GetTimeSuffixSecond()
	jti = suffix
	return jti
}
