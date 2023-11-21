package db

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/db/dbConnector"
)

/*
Init
*/

func MakeEntityDB() (entityDB *EntityDB) {
	entityDB = new(EntityDB)
	entityDB.Init()
	return entityDB
}

func (self *EntityDB) Init() *EntityDB {
	self.setEntityDBConnector()
	return self
}

func (self *EntityDB) setEntityDBConnector() *EntityDB {
	self.EntityDBConnector = dbConnector.MakeEntityDBConnector()
	return self
}

/*
Close
*/

func (self *EntityDB) CloseDB() {
	self.closeDBConnector()
	return
}

func (self *EntityDB) closeDBConnector() *EntityDB {
	if self.EntityDBConnector == nil {
		return self
	}
	self.EntityDBConnector.CloseDBConnector()
	self.EntityDBConnector = nil
	return self
}
