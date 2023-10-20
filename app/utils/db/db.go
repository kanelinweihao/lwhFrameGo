package db

import (
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.lwh.com/linweihao/lwhFrameGo/app/conf"
	"go.lwh.com/linweihao/lwhFrameGo/app/factory/factoryOfSSH"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/base"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/conv"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/err"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/goroutine"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/ssh"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/time"
)

var driver string = "mysql"

// struct
type EntityDB struct {
	DBSqlx    *sqlx.DB
	EntitySSH *ssh.EntitySSH
}
type TypeEntityData interface{}

////
// Init
////

func InitDB() (entityDB *EntityDB) {
	entitySSH := getEntitySSH()
	dbSqlx := getDBSqlx()
	entityDB = &EntityDB{
		DBSqlx:    dbSqlx,
		EntitySSH: entitySSH,
	}
	return entityDB
}

func getEntitySSH() (entitySSH *ssh.EntitySSH) {
	entitySSH = nil
	isNeedSSH := isNeedSSH()
	if isNeedSSH {
		// entitySSH = ssh.InitSSHForMysql()
		entitySSH = factoryOfSSH.MakeEntityOfSSHForMysql()
	}
	return entitySSH
}

func isNeedSSH() (isNeedSSH bool) {
	isNeedSSH = conf.IsNeedSSH()
	return isNeedSSH
}

func getDBSqlx() (dbSqlx *sqlx.DB) {
	m := conf.GetEntityConfigMysql()
	dsn := getDSN(m)
	dbSqlx, errDB := sqlx.Open(driver, dsn)
	err.ErrCheck(errDB)
	// dd.DD(dbSqlx)
	// ping := dbSqlx.Ping
	// dd.DD(ping)
	return dbSqlx
}

func getDSN(m *conf.EntityConfigMysql) (dsn string) {
	mysqlHost := m.Host
	mysqlPort := m.Port
	mysqlUser := m.User
	mysqlPassword := m.Password
	mysqlDBName := m.DBName
	mysqlCharset := m.Charset
	isNeedSSH := isNeedSSH()
	network := ssh.NetworkTCP
	if isNeedSSH {
		network = ssh.NetworkTCPSSH
	}
	networkFull := fmt.Sprintf(
		"%s(%s:%s)",
		network,
		mysqlHost,
		mysqlPort)
	// dd.DD(tcp)
	dsn = fmt.Sprintf(
		"%s:%s@%s/%s?charset=%s",
		mysqlUser,
		mysqlPassword,
		networkFull,
		mysqlDBName,
		mysqlCharset)
	// dd.DD(dsn)
	return dsn
}

////
// Close
////

func (self *EntityDB) CloseDB() {
	dbSqlx := self.DBSqlx
	dbSqlx.Close()
	// time.ShowTimeAndMsg("DB close success")
	isNeedSSH := conf.IsNeedSSH()
	if isNeedSSH {
		entitySSH := self.EntitySSH
		entitySSH.CloseSSH()
		// time.ShowTimeAndMsg("SSH close success")
	}
	return
}

////
// Exec
////

func GetArrAttrForExcel[T interface{}](self *EntityDB, arrEntity []T, query string, queryValue ...interface{}) (attrS2 base.AttrS2) {
	dbSqlx := self.DBSqlx
	// dd.DD(query)
	errSqlxSelect := dbSqlx.Select(
		&arrEntity,
		query,
		queryValue...)
	err.ErrCheck(errSqlxSelect)
	// dd.DD(arrEntity)
	attrT2 := conv.ToAttrT2FromArrEntity(arrEntity)
	// dd.DD(attrT2)
	attrS2 = conv.ToAttrS2FromAttrT2(attrT2)
	// dd.DD(attrS2)
	return attrS2
}

func GetArrAttrForExcelUseGoroutine[T interface{}](entityChannel *goroutine.EntityChannel, self *EntityDB, arrEntity []T, query string, queryValue ...interface{}) {
	attrS2 := GetArrAttrForExcel(
		self,
		arrEntity,
		query,
		queryValue...)
	entityChannel.WriteOnce(attrS2)
	// time.ShowTimeAndMsg("channel write success")
	return
}
