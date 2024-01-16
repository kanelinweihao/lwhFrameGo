package ip

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"net"
	"net/http"
	"strings"
)

func GetIPOfServer() (ip string) {
	addrs, errAddrs := net.InterfaceAddrs()
	err.ErrCheck(errAddrs)
	for _, addr := range addrs {
		ipnet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		isLoopBack := ipnet.IP.IsLoopback()
		if isLoopBack {
			continue
		}
		if ipnet.IP.To4() != nil {
			ip = ipnet.IP.String()
		}
	}
	return ip
}

func ShowIP() {
	ip := GetIPOfServer()
	msg := fmt.Sprintf(
		"%s\nThe ip of server is |%s| now\n%s",
		"----------",
		ip,
		"----------")
	fmt.Println(msg)
	return
}

func GetIPOfClient(req *http.Request) (ip string) {
	xForwardedFor := req.Header.Get("X-Forwarded-For")
	ip = strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	xRealIp := req.Header.Get("X-Real-Ip")
	ip = strings.TrimSpace(xRealIp)
	if ip != "" {
		return ip
	}
	ip, _, errS := net.SplitHostPort(strings.TrimSpace(req.RemoteAddr))
	if errS == nil {
		return ip
	}
	return ""
}
