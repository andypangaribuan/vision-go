package db

import (
	"database/sql"
	"fmt"
	"github.com/andypangaribuan/vision-go/models"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"strings"
	"time"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/18
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
func (*dbStruct) BuildMySQLConnStr(host string, port int, dbName, username, password string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?parseTime=true", username, password, host, port, dbName)
}

func (*dbStruct) BuildPostgreSQLConnStr(host string, port int, dbName, username, password string, schema *string) string {
	if schema == nil {
		return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbName)
	} else {
		return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable", host, port, username, password, dbName, *schema)
	}
}

func (*dbStruct) CreateInstance(dbType string, masterConnStr, slaveConnStr *string, maxLifeTimeConn, maxIdleConn, maxOpenConn int, autoRebind, unsafeCompatibility bool) (master models.IDbMaster, slave models.IDbSlave) {
	dbType = strings.ToLower(dbType)
	if dbType != dbTypeMySql && dbType != dbTypePostgres {
		return
	}

	var (
		_maxLifeTimeConn = time.Minute * 3
		_maxIdleConn     = 3
		_maxOpenConn     = 100
	)

	if maxLifeTimeConn > 0 {
		_maxLifeTimeConn = time.Minute * time.Duration(maxLifeTimeConn)
	}
	if maxIdleConn > 0 {
		_maxIdleConn = maxIdleConn
	}
	if maxOpenConn > 0 {
		_maxOpenConn = maxOpenConn
	}

	conn := func(connStr string) *connStruct {
		return &connStruct{
			dbType:              dbType,
			connStr:             connStr,
			instance:            nil,
			maxLifeTimeConn:     _maxLifeTimeConn,
			maxIdleConn:         _maxIdleConn,
			maxOpenConn:         _maxOpenConn,
			autoRebind:          autoRebind,
			unsafeCompatibility: unsafeCompatibility,
		}
	}

	if masterConnStr != nil {
		master = &masterStruct{conn: conn(*masterConnStr)}
	}
	if slaveConnStr != nil {
		slave = &slaveStruct{conn: conn(*slaveConnStr)}
	}

	return
}



func normalizeSqlQueryPars(conn *connStruct, sqlQuery string, sqlPars []interface{}) (string, []interface{}) {
	query := sqlQuery
	pars := sqlPars

	if len(sqlPars) == 1 {
		switch data := sqlPars[0].(type) {
		case []interface{}: pars = data
		}
	}

	if conn.dbType == dbTypePostgres && conn.autoRebind && len(pars) > 0 {
		query = sqlx.Rebind(sqlx.DOLLAR, sqlQuery)
	}
	return query, pars
}


func sqlxConnect(conn *connStruct) (*sqlx.DB, error) {
	driverName := conn.dbType
	if conn.dbType == dbTypePostgres {
		driverName = "postgres"
	}

	instance, err := sqlx.Connect(driverName, conn.connStr)
	if err == nil {
		instance.SetConnMaxLifetime(conn.maxLifeTimeConn)
		instance.SetMaxIdleConns(conn.maxIdleConn)
		instance.SetMaxOpenConns(conn.maxOpenConn)
		err = instance.Ping()
	}
	if err != nil {
		err = errors.WithStack(err)
	}

	return instance, err
}


func sqlxExec(conn *connStruct, sqlQuery string, sqlPars ...interface{}) (sql.Result, error) {
	if conn.dbType == dbTypePostgres && conn.autoRebind && len(sqlPars) > 0 {
		sqlQuery = sqlx.Rebind(sqlx.DOLLAR, sqlQuery)
	}

	stmt, err := conn.instance.Prepare(sqlQuery)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}
	defer func() { _ = stmt.Close() }()

	res, err := stmt.Exec(sqlPars...)
	if err != nil {
		err = errors.WithStack(err)
	}
	return res, err
}
