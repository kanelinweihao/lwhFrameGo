package ssh

import (
	"context"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"go.lwh.com/linweihao/lwhFrameGo/app/conf"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/err"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/file"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/ip"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/pack"
	"golang.org/x/crypto/ssh"
	"net"
	"time"
	// ttt "go.lwh.com/linweihao/lwhFrameGo/app/utils/time"
)

var NetworkTCP string = "tcp"
var NetworkTCPSSH string = "tcp+ssh"
var TimeoutSSH time.Duration = time.Second * time.Duration(5)

type EntitySSH struct {
	clientDial *ssh.Client
	network    string
	address    string
}

////
// Init
////

func InitSSHForMysql() (entitySSH *EntitySSH) {
	entitySSH = initSSH()
	entitySSH.RegisterDialToMysql()
	return entitySSH
}

func InitSSHForRedis() (entitySSH *EntitySSH) {
	entitySSH = initSSH()
	return entitySSH
}

func initSSH() (entitySSH *EntitySSH) {
	// ip.ShowIP()
	s := conf.GetEntityConfigSSH()
	// dd.DD(s.TypeAuth)
	clientDial, network, address := dialSSH(s)
	// dd.DD(clientDial)
	entitySSH = &EntitySSH{
		clientDial: clientDial,
		network:    network,
		address:    address,
	}
	return entitySSH
}

func dialSSH(s *conf.EntityConfigSSH) (clientDial *ssh.Client, network string, address string) {
	sshHost := s.Host
	sshPort := s.Port
	sshUser := s.User
	sshTypeAuth := s.TypeAuth
	sshPassword := s.Password
	sshPathPrivateKey := s.PathPrivateKey
	network = "tcp"
	address = getAddressToDial(
		sshHost,
		sshPort)
	configSSHClient := getConfigSSHClient(
		sshUser,
		sshTypeAuth,
		sshPassword,
		sshPathPrivateKey)
	clientDial, errDial := ssh.Dial(
		network,
		address,
		configSSHClient)
	err.ErrCheck(errDial)
	// ttt.ShowTimeAndMsg("Dial set client success")
	// ip.ShowIP()
	return clientDial, network, address
}

func getAddressToDial(sshHost string, sshPort string) (address string) {
	address = fmt.Sprintf(
		"%s:%s",
		sshHost,
		sshPort)
	return address
}

func getConfigSSHClient(sshUser string, sshTypeAuth string, sshPassword string, sshPathPrivateKey string) (configSSHClient *ssh.ClientConfig) {
	hostKeyCallBack := ssh.InsecureIgnoreHostKey()
	timeoutSSH := TimeoutSSH
	entitySSHConfig := ssh.ClientConfig{
		User:            sshUser,
		HostKeyCallback: hostKeyCallBack,
		Timeout:         timeoutSSH,
	}
	sshAuth := getConfigSSHAuth(
		sshTypeAuth,
		sshPassword,
		sshPathPrivateKey)
	entitySSHConfig.Auth = sshAuth
	configSSHClient = &entitySSHConfig
	return configSSHClient
}

func getConfigSSHAuth(sshTypeAuth string, sshPassword string, sshPathPrivateKey string) (sshAuth []ssh.AuthMethod) {
	switch sshTypeAuth {
	case "Password":
		sshAuth = []ssh.AuthMethod{
			ssh.Password(sshPassword),
		}
	case "Key":
		// arrBytePrivateKey := file.ReadFileAsArrayByte(sshPathPrivateKey)
		arrBytePrivateKey := pack.GetArrayBytePrivateKey(sshPathPrivateKey)
		signer, errParse := ssh.ParsePrivateKey(arrBytePrivateKey)
		err.ErrCheck(errParse)
		sshAuth = []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		}
	default:
		msgError := fmt.Sprintf(
			"The typeAuth |%s| is invalid of |%s|",
			sshTypeAuth,
			"dialSSH")
		err.ErrPanic(msgError)
	}
	return sshAuth
}

////
// Exec
////

func (self *EntitySSH) CloseSSH() {
	clientDial := self.GetClientDial()
	clientDial.Close()
	return
}

////
// Exec
////

func (self *EntitySSH) GetClientDial() (clientDial *ssh.Client) {
	return self.clientDial
}

func (self *EntitySSH) GetNetwork() (network string) {
	return self.network
}

func (self *EntitySSH) GetAddress() (address string) {
	return self.address
}

func (self *EntitySSH) SetAddress(address string) {
	self.address = address
	return
}

func (self *EntitySSH) RegisterDialToMysql() {
	funcDialMysql := self.DialForMysql
	networkDialMysql := NetworkTCPSSH
	mysql.RegisterDialContext(
		networkDialMysql,
		funcDialMysql)
	// ttt.ShowTimeAndMsg("Dial register mysql success")
	// ip.ShowIP()
	return
}

func (self *EntitySSH) DialForMysql(context context.Context, address string) (coon net.Conn, errDial error) {
	// dd.DD(context)
	// dd.DD(address)
	network := self.network
	coon, errDial = self.clientDial.Dial(network, address)
	err.ErrCheck(errDial)
	return coon, errDial
}

func (self *EntitySSH) DialForRedis() (coon net.Conn, errDial error) {
	network := self.network
	address := self.address
	// dd.DD(network)
	// dd.DD(address)
	coon, errDial = self.clientDial.Dial(network, address)
	err.ErrCheck(errDial)
	return coon, errDial
}
