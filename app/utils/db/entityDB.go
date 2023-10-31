package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/goroutine"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/ssh"
)

type EntityDB struct {
	DBSqlx            *sqlx.DB
	EntityConfigMysql *conf.EntityConfigMysql
	IsNeedSSH         bool
	EntitySSH         *ssh.EntitySSH
}

func GetArrAttrForExcel[T interface{}](self *EntityDB, arrEntity []T, query string, queryValue ...interface{}) (attrS2 base.AttrS2) {
	d := self.DBSqlx
	errSqlxSelect := d.Select(
		&arrEntity,
		query,
		queryValue...)
	err.ErrCheck(errSqlxSelect)
	attrT2 := conv.ToAttrT2FromArrEntity(arrEntity)
	attrS2 = conv.ToAttrS2FromAttrT2(attrT2)
	return attrS2
}

func GetArrAttrForExcelUseGoroutine[T interface{}](entityChannel *goroutine.EntityChannel, self *EntityDB, arrEntity []T, query string, queryValue ...interface{}) {
	attrS2 := GetArrAttrForExcel(
		self,
		arrEntity,
		query,
		queryValue...)
	entityChannel.WriteOnce(attrS2)
	return
}
