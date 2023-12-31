package dbConnector

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/ssh"
)

func InitEntityDBConnector() (entityDBConnector *EntityDBConnector) {
	entityDBConnector = new(EntityDBConnector)
	entityDBConnector.Init()
	return entityDBConnector
}

func (self *EntityDBConnector) Init() *EntityDBConnector {
	self.setEntitySSH().setEntityConfigMysql().setDBSqlx().pingDBSqlx()
	return self
}

func (self *EntityDBConnector) setEntitySSH() *EntityDBConnector {
	isNeedSSH := conf.IsNeedSSH()
	self.IsNeedSSH = isNeedSSH
	if !isNeedSSH {
		return self
	}
	entitySSH := ssh.InitEntitySSH()
	funcDialMysql := entitySSH.DialForMysql
	networkDialMysql := ssh.NetworkTCPANDSSH
	mysql.RegisterDialContext(networkDialMysql, funcDialMysql)
	self.EntitySSH = entitySSH
	return self
}

func (self *EntityDBConnector) setEntityConfigMysql() *EntityDBConnector {
	ec := conf.GetEntityConfigMysql()
	self.EntityConfigMysql = ec
	return self
}

func (self *EntityDBConnector) setDBSqlx() *EntityDBConnector {
	driver := conf.DriverDB
	dsn := self.initDSN()
	m, errDB := sqlx.Open(driver, dsn)
	err.ErrCheck(errDB)
	self.DBSqlx = m
	return self
}

func (self *EntityDBConnector) initDSN() (dsn string) {
	ec := self.EntityConfigMysql
	mysqlHost := ec.Host
	mysqlPort := ec.Port
	mysqlUser := ec.User
	mysqlPassword := ec.Password
	mysqlDBName := ec.DBName
	mysqlCharset := ec.Charset
	network := ssh.NetworkTCP
	isNeedSSH := self.IsNeedSSH
	if isNeedSSH {
		network = ssh.NetworkTCPANDSSH
	}
	networkFull := fmt.Sprintf(
		"%s(%s:%s)",
		network,
		mysqlHost,
		mysqlPort)
	dsn = fmt.Sprintf(
		"%s:%s@%s/%s?charset=%s",
		mysqlUser,
		mysqlPassword,
		networkFull,
		mysqlDBName,
		mysqlCharset)
	return dsn
}

func (self *EntityDBConnector) pingDBSqlx() *EntityDBConnector {
	errPing := self.DBSqlx.Ping()
	err.ErrCheck(errPing)
	return self
}

func (self *EntityDBConnector) CloseDBConnector() {
	self.closeDBSqlx().closeSSH()
	return
}

func (self *EntityDBConnector) closeDBSqlx() *EntityDBConnector {
	if self.DBSqlx == nil {
		return self
	}
	self.DBSqlx.Close()
	self.DBSqlx = nil
	return self
}

func (self *EntityDBConnector) closeSSH() *EntityDBConnector {
	if !self.IsNeedSSH {
		self.EntitySSH = nil
		return self
	}
	if self.EntitySSH == nil {
		return self
	}
	self.EntitySSH.CloseSSH()
	self.EntitySSH = nil
	return self
}
