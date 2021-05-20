package db

import (
	"fmt"
	"github.com/andypangaribuan/vision-go/models"
	"github.com/andypangaribuan/vision-go/vis"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/18
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
func (slf *masterStruct) connect() (*sqlx.DB, error) {
	if slf.conn.instance != nil {
		return slf.conn.instance, nil
	}

	instance, err := sqlxConnect(slf.conn)
	if err == nil {
		slf.conn.instance = instance
	}

	return instance, err
}


func (slf *masterStruct) Ping() error {
	instance, err := slf.connect()
	if err == nil {
		err = instance.Ping()
		if err != nil {
			err = errors.WithStack(err)
		}
	}
	return err
}

func (slf *masterStruct) Exec(sqlQuery string, sqlPars ...interface{}) error {
	_, err := slf.connect()
	if err != nil {
		return err
	}

	query, pars := normalizeSqlQueryPars(slf.conn, sqlQuery, sqlPars)
	_, err = sqlxExec(slf.conn, query, pars...)
	return err
}

func (slf *masterStruct) Transaction() (models.IDbTransaction, error) {
	_, err := slf.connect()
	if err != nil {
		return nil, err
	}

	return &transactionStruct{
		master:     slf,
		sqlQueries: make([]string, 0),
		sqlPars:    make([][]interface{}, 0),
	}, nil
}




func (slf *transactionStruct) SqlQueries() []string {
	return slf.sqlQueries
}

func (slf *transactionStruct) SqlPars() [][]interface{} {
	return slf.sqlPars
}

func (slf *transactionStruct) Exec(sqlQuery string, sqlPars ...interface{}) {
	query, pars := normalizeSqlQueryPars(slf.master.conn, sqlQuery, sqlPars)
	slf.sqlQueries = append(slf.sqlQueries, query)
	slf.sqlPars = append(slf.sqlPars, pars)
}

func (slf *transactionStruct) Commit() *models.DbTxError {
	if len(slf.sqlQueries) == 0 {
		return nil
	}

	instance, err := slf.master.connect()
	if err != nil {
		return &models.DbTxError{Err: err}
	}

	tx, err := instance.Begin()
	if err != nil {
		return &models.DbTxError{Err: errors.WithStack(err)}
	}

	errMsg := ""
	for i, query := range slf.sqlQueries {
		_, err = tx.Exec(query, slf.sqlPars[i]...)
		if err != nil {
			pars := slf.sqlPars[i]
			msg := ":: START ERROR ON SQL TRANSACTION"
			msg += "\n# error at query no. %v"
			msg += "\n# query:\n%v"
			if len(pars) == 0 {
				msg = fmt.Sprintf(msg, i+1, query)
			} else {
				ms := make(map[int]interface{}, 0)
				for i, v := range pars {
					ms[i+1] = v
				}

				if encode, err := vis.Json.Encode(ms); err == nil {
					msg = fmt.Sprintf(msg + "\n# pars:\n%v", i+1, query, encode)
				} else {
					msg = fmt.Sprintf(msg + "\n# pars:\n%+v", i+1, query, ms)
				}
			}

			msg += "\n\n:: ALL QUERY"
			for _i, _q := range slf.sqlQueries {
				p := slf.sqlPars[_i]
				m := "\n# query no. %v:"
				m += "\n$ %v"
				if len(p) == 0 {
					m = fmt.Sprintf(m, _i+1, _q)
				} else {
					if p1, _e := vis.Json.Encode(p); _e == nil {
						m = fmt.Sprintf(m + "\n$ %v", _i+1, _q, p1)
					} else {
						m = fmt.Sprintf(m + "\n$ sql pars encoding error: %v", _i+1, _q, _e)
					}
				}
				m += "\n"
			}
			msg += ":: END ERROR ON SQL TRANSACTION"
			errMsg = msg

			err = errors.WithStack(err)
			_ = tx.Rollback()
			break
		}
	}

	if err == nil {
		err = tx.Commit()
		if err != nil {
			err = errors.WithStack(err)
		}
	}

	if err != nil {
		return &models.DbTxError{
			Err: err,
			Msg: errMsg,
		}
	}

	return nil
}
