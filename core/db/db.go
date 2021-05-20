package db


/* ============================================
	Created by andy pangaribuan on 2021/05/18
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
var Db *dbStruct

const (
	dbTypePostgres = "postgresql"
	dbTypeMySql    = "mysql"
)

func init() {
	Db = &dbStruct{}
}
