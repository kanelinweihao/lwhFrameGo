package dbConnector

import (
	"github.com/jmoiron/sqlx"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/ssh"
)

type EntityDBConnector struct {
	DBSqlx            *sqlx.DB
	EntityConfigMysql *conf.EntityConfigMysql
	IsNeedSSH         bool
	EntitySSH         *ssh.EntitySSH
}
