package db

import (
	"github.com/andypangaribuan/vision-go/models"
	"github.com/andypangaribuan/vision-go/vis"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/18
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
func (slf *slaveStruct) connect() (*sqlx.DB, error) {
	if slf.conn.instance != nil {
		return slf.conn.instance, nil
	}

	instance, err := sqlxConnect(slf.conn)
	if err == nil {
		slf.conn.instance = instance
	}

	return instance, err
}


func (slf *slaveStruct) Ping() error {
	instance, err := slf.connect()
	if err == nil {
		err = instance.Ping()
		if err != nil {
			err = errors.WithStack(err)
		}
	}
	return err
}


func (slf *slaveStruct) Count(sqlQuery string, sqlPars ...interface{}) (int, error) {
	instance, err := slf.connect()
	if err != nil {
		return 0, err
	}

	query, pars := normalizeSqlQueryPars(slf.conn, sqlQuery, sqlPars)
	var count int
	if err := instance.Get(&count, query, pars...); err != nil {
		return 0, errors.WithStack(err)
	}
	return count, nil
}


func (slf *slaveStruct) Select(out interface{}, sqlQuery string, sqlPars ...interface{}) (*models.DbUnsafeSelectError, error) {
	instance, err := slf.connect()
	if err != nil {
		return nil, err
	}

	query, pars := normalizeSqlQueryPars(slf.conn, sqlQuery, sqlPars)
	err = instance.Select(out, query, pars...)
	if err != nil {
		err = errors.WithStack(err)
		if slf.conn.unsafeCompatibility {
			msg := err.Error()
			unsafe := models.DbUnsafeSelectError{
				LogType:    "error",
				SqlQuery:   query,
				SqlPars:    pars,
				LogMessage: &msg,
				LogTrace:   vis.Log.Stack(query, pars, err),
			}

			err = instance.Unsafe().Select(out, query, pars...)
			if err != nil {
				err = errors.WithStack(err)
			}
			return &unsafe, err
		}
	}

	return nil, err
}
