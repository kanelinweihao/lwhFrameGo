package ssh

import (
	"context"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"golang.org/x/crypto/ssh"
	"net"
	"time"
)

var NetworkTCP string = "tcp"
var NetworkTCPANDSSH string = "tcp+ssh"
var TimeoutSSH time.Duration = time.Second * time.Duration(5)

type EntitySSH struct {
	ClientDial      *ssh.Client
	Network         string
	Address         string
	EntityConfigSSH *conf.EntityConfigSSH
}

func (self *EntitySSH) DialForMysql(ctx context.Context, address string) (coon net.Conn, errDial error) {
	network := self.Network
	coon, errDial = self.ClientDial.Dial(network, address)
	err.ErrCheck(errDial)
	return coon, errDial
}

func (self *EntitySSH) DialForRedis(ctx context.Context, network string, address string) (coon net.Conn, errDial error) {
	coon, errDial = self.ClientDial.Dial(network, address)
	err.ErrCheck(errDial)
	return coon, errDial
}
