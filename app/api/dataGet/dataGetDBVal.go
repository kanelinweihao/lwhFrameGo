package dataGet

import (
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
)

var queryUID string = `SELECT
id AS 'UID',
mobile_no AS '手机号'
FROM users
WHERE id = ?
;
`
var queryOID string = `SELECT
user_id AS 'UID',
org_id AS '机构编号'
FROM users_org
WHERE user_id = ?
;
`

type TypeEntityData interface {
	EntityUID | EntityOID
}

type EntityUID struct {
	UserId   int    `db:"UID"`
	MobileNo string `db:"手机号"`
}

type EntityOID struct {
	UserId int    `db:"UID"`
	OrgId  string `db:"机构编号"`
}
