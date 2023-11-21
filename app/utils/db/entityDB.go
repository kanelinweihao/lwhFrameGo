package db

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/db/dbConnector"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/db/dbReader"
)

type EntityDB struct {
	BoxFromDB         base.AttrS3
	BoxForDB          base.BoxData
	EntityDBConnector *dbConnector.EntityDBConnector
	EntityDBReader    *dbReader.EntityDBReader
}

func (self *EntityDB) BatchGetDataFromDB(boxForDB base.BoxData) (boxFromDB base.AttrS3) {
	self.BoxForDB = boxForDB
	entityDBConnector := self.EntityDBConnector
	entityDBReader := dbReader.MakeEntityDBReader(entityDBConnector, boxForDB)
	self.EntityDBReader = entityDBReader
	boxFromDB = entityDBReader.BatchReadDB()
	self.BoxFromDB = boxFromDB
	return boxFromDB
}
