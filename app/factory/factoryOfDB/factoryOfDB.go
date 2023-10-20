package factoryOfDB

import (
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/db"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
)

func MakeEntityOfDB() (entityDB *db.EntityDB) {
	entityDB = db.InitDB()
	return entityDB
}
