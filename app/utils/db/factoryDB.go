package db

import (
	_ "database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/ssh"
)

var driver string = "mysql"

/*
Init
*/

func MakeEntityOfDB() (entityDB *EntityDB) {
	entityDB = initEntityDB()
	return entityDB
}

func initEntityDB() (entityDB *EntityDB) {
	entityDB = new(EntityDB)
	entityDB.Init()
	return entityDB
}

func (self *EntityDB) Init() *EntityDB {
	self.setEntitySSH().setEntityConfigMysql().setDBSqlx().pingDBSqlx()
	return self
}

func (self *EntityDB) setEntitySSH() *EntityDB {
	isNeedSSH := conf.IsNeedSSH()
	self.IsNeedSSH = isNeedSSH
	if !isNeedSSH {
		return self
	}
	entitySSH := ssh.MakeEntityOfSSHForMysql()
	funcDialMysql := entitySSH.DialForMysql
	networkDialMysql := ssh.NetworkTCPANDSSH
	mysql.RegisterDialContext(networkDialMysql, funcDialMysql)
	self.EntitySSH = entitySSH
	return self
}

func (self *EntityDB) setEntityConfigMysql() *EntityDB {
	ec := conf.GetEntityConfigMysql()
	self.EntityConfigMysql = ec
	return self
}

func (self *EntityDB) setDBSqlx() *EntityDB {
	dsn := self.initDSN()
	m, errDB := sqlx.Open(driver, dsn)
	err.ErrCheck(errDB)
	self.DBSqlx = m
	return self
}

func (self *EntityDB) initDSN() (dsn string) {
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

func (self *EntityDB) pingDBSqlx() *EntityDB {
	errPing := self.DBSqlx.Ping()
	err.ErrCheck(errPing)
	return self
}

/*
Close
*/

func (self *EntityDB) CloseDB() {
	self.closeDBSqlx().closeSSH()
	return
}

func (self *EntityDB) closeDBSqlx() *EntityDB {
	if self.DBSqlx == nil {
		return self
	}
	self.DBSqlx.Close()
	self.DBSqlx = nil
	return self
}

func (self *EntityDB) closeSSH() *EntityDB {
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
