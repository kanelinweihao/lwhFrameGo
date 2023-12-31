package ssh

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/pack"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/times"
	"golang.org/x/crypto/ssh"
)

func InitEntitySSH() (entitySSH *EntitySSH) {
	entitySSH = new(EntitySSH)
	entitySSH.Init()
	return entitySSH
}

func (self *EntitySSH) Init() *EntitySSH {
	self.setEntityConfigSSH().dialSSH()
	return self
}

func (self *EntitySSH) setEntityConfigSSH() *EntitySSH {
	ec := conf.GetEntityConfigSSH()
	self.EntityConfigSSH = ec
	return self
}

func (self *EntitySSH) dialSSH() *EntitySSH {
	ec := self.EntityConfigSSH
	sshHost := ec.Host
	sshPort := ec.Port
	sshUser := ec.User
	sshTypeAuth := ec.TypeAuth
	sshPassword := ec.Password
	sshPathPrivateKey := ec.PathPrivateKey
	network := NetworkTCP
	address := getAddressToDial(
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
	if errDial == nil {
		self.ClientDial = clientDial
		self.Network = network
		self.Address = address
		return self
	}
	times.Sleep(2, "s")
	clientDial, errDialAgain := ssh.Dial(
		network,
		address,
		configSSHClient)
	err.ErrCheck(errDialAgain)
	self.ClientDial = clientDial
	self.Network = network
	self.Address = address
	return self
}

func getAddressToDial(sshHost string, sshPort string) (address string) {
	address = fmt.Sprintf(
		"%s:%s",
		sshHost,
		sshPort)
	return address
}

func getConfigSSHClient(sshUser string, sshTypeAuth string, sshPassword string, sshPathPrivateKey string) (entitySSHConfig *ssh.ClientConfig) {
	hostKeyCallBack := ssh.InsecureIgnoreHostKey()
	timeoutSSH := TimeoutSSH
	sshAuth := getConfigSSHAuth(
		sshTypeAuth,
		sshPassword,
		sshPathPrivateKey)
	entitySSHConfig = &ssh.ClientConfig{
		User:            sshUser,
		HostKeyCallback: hostKeyCallBack,
		Timeout:         timeoutSSH,
		Auth:            sshAuth,
	}
	return entitySSHConfig
}

func getConfigSSHAuth(sshTypeAuth string, sshPassword string, sshPathPrivateKey string) (sshAuth []ssh.AuthMethod) {
	switch sshTypeAuth {
	case "Password":
		sshAuth = []ssh.AuthMethod{
			ssh.Password(sshPassword),
		}
	case "Key":
		arrBytePrivateKey := pack.ReadFileEmbedAsArrayByte(sshPathPrivateKey)
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

func (self *EntitySSH) CloseSSH() {
	self.ClientDial.Close()
	return
}
