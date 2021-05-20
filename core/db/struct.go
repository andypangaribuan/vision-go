package db

import (
	"github.com/jmoiron/sqlx"
	"time"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/18
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
type dbStruct struct { }

type masterStruct struct {
	conn *connStruct
}

type slaveStruct struct {
	conn *connStruct
}

type transactionStruct struct {
	master     *masterStruct
	sqlQueries []string
	sqlPars    [][]interface{}
}


type connStruct struct {
	dbType              string
	connStr             string
	instance            *sqlx.DB
	maxLifeTimeConn     time.Duration
	maxIdleConn         int
	maxOpenConn         int
	autoRebind          bool
	unsafeCompatibility bool
}
