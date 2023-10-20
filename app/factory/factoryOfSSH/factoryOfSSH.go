package factoryOfSSH

import (
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/ssh"
)

func MakeEntityOfSSHForMysql() (entitySSHMysql *ssh.EntitySSH) {
	entitySSHMysql = ssh.InitSSHForMysql()
	return entitySSHMysql
}

func MakeEntityOfSSHForRedis() (entitySSHRedis *ssh.EntitySSH) {
	entitySSHRedis = ssh.InitSSHForRedis()
	return entitySSHRedis
}
