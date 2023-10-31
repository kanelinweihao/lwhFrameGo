package ip

import (
	"fmt"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"net"
)

func GetIP() (ip string) {
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
	ip := GetIP()
	msg := fmt.Sprintf(
		"%s\nThe ip is |%s| now\n%s",
		"----------",
		ip,
		"----------")
	fmt.Println(msg)
	return
}
