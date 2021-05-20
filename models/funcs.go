package models


/* ============================================
	Created by andy pangaribuan on 2021/05/19
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
type IDbMaster interface {
	Ping() error
	Exec(sqlQuery string, sqlPars ...interface{}) error
	Transaction() (IDbTransaction, error)
}

type IDbSlave interface {
	Ping() error
	Count(sqlQuery string, sqlPars ...interface{}) (int, error)
	Select(out interface{}, sqlQuery string, sqlPars ...interface{}) (*DbUnsafeSelectError, error)
}

type IDbTransaction interface {
	SqlQueries() []string
	SqlPars() [][]interface{}
	Exec(sqlQuery string, sqlPars ...interface{})
	Commit() *DbTxError
}
