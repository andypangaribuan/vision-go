package api

import (
	"github.com/andypangaribuan/vision-go/core/clog"
	"github.com/andypangaribuan/vision-go/models"
	"github.com/andypangaribuan/vision-go/vis"
	"time"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/18
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
func (slf *smLog) TraceV1(logType string, message string, trace *string) {
	if !clog.IsHaveCLog() {
		return
	}

	var reqFromServiceName *string
	var reqFromServiceVersion *string
	var reqUid *string


	req := slf.context.echo.c.Request()

	if v := req.Header.Get(vis.Conf.RequestFromServiceName()); v != "" {
		reqFromServiceName = &v
	}
	if v := req.Header.Get(vis.Conf.RequestFromServiceVersion()); v != "" {
		reqFromServiceVersion = &v
	}
	if v := req.Header.Get(vis.Conf.RequestUidKey()); v != "" {
		reqUid = &v
	}

	log := clog.LogTraceV1{
		ServiceName:               vis.Store.ServiceName,
		ServiceVersion:            vis.Store.ServiceVersion,
		TraceId:                   slf.context.GetTraceId(),
		RequestFromServiceName:    reqFromServiceName,
		RequestFromServiceVersion: reqFromServiceVersion,
		RequestUid:                reqUid,
		LogType:                   logType,
		LogMessage:                message,
		LogTrace:                  trace,
		LogDate:                   time.Now().UTC(),
	}

	go clog.SendLogTraceV1(log)
}


func (slf *smLog) ErrorV1(skipLevel int, message *string, gr *models.APIGeneralResult, args ...interface{}) {
	if !clog.IsHaveCLog() {
		return
	}

	trace := ""
	if len(args) > 0 {
		trace = vis.Log.BaseStack(skipLevel, args...)
	}

	var response *string
	if gr != nil {
		if encode, err := vis.Json.Encode(*gr); err == nil {
			response = &encode
		}
	}

	slf.BaseErrorV1(message, trace, response)
}


func (slf *smLog) QueryV1(logType, query string, params, message, trace *string) {
	if !clog.IsHaveCLog() {
		return
	}

	log := clog.LogQueryV1{
		ServiceName:    vis.Store.ServiceName,
		ServiceVersion: vis.Store.ServiceVersion,
		TraceId:        slf.context.GetTraceId(),
		LogType:        logType,
		SqlQuery:       query,
		SqlParams:      params,
		LogMessage:     message,
		LogTrace:       trace,
		LogDate:        time.Now().UTC(),
	}

	go clog.SendLogQuery(log)
}


func (slf *smLog) UnsafeQueryV1(unsafe models.DbUnsafeSelectError) {
	if !clog.IsHaveCLog() {
		return
	}

	pars, err := vis.Json.Encode(unsafe.SqlPars)
	if err != nil {
		pars = "json encode error: " + err.Error()
	}

	log := clog.LogQueryV1{
		ServiceName:    vis.Store.ServiceName,
		ServiceVersion: vis.Store.ServiceVersion,
		TraceId:        slf.context.GetTraceId(),
		LogType:        unsafe.LogType,
		SqlQuery:       unsafe.SqlQuery,
		SqlParams:      &pars,
		LogMessage:     unsafe.LogMessage,
		LogTrace:       unsafe.LogTrace,
		LogDate:        time.Time{},
	}

	go clog.SendLogQuery(log)
}


func (slf *smLog) BaseErrorV1(message *string, trace string, response *string) {
	if !clog.IsHaveCLog() {
		return
	}

	var reqFromServiceName *string
	var reqFromServiceVersion *string
	var reqUid *string
	var reqIp *string
	var reqHost *string
	var reqUri *string
	var reqMethod *string
	var reqAgent *string

	req := slf.context.echo.c.Request()

	if v := req.Header.Get(vis.Conf.RequestFromServiceName()); v != "" {
		reqFromServiceName = &v
	}
	if v := req.Header.Get(vis.Conf.RequestFromServiceVersion()); v != "" {
		reqFromServiceVersion = &v
	}
	if v := req.Header.Get(vis.Conf.RequestUidKey()); v != "" {
		reqUid = &v
	}
	if v := slf.context.echo.c.RealIP(); v != "" {
		reqIp = &v
	}
	if v := req.Host; v != "" {
		reqHost = &v
	}
	if v := req.RequestURI; v != "" {
		reqUri = &v
	}
	if v := req.Method; v != "" {
		reqMethod = &v
	}
	if v := req.UserAgent(); v != "" {
		reqAgent = &v
	}

	log := clog.LogErrorV1{
		ServiceName:               vis.Store.ServiceName,
		ServiceVersion:            vis.Store.ServiceVersion,
		TraceId:                   slf.context.GetTraceId(),
		RequestFromServiceName:    reqFromServiceName,
		RequestFromServiceVersion: reqFromServiceVersion,
		RequestUid:                reqUid,
		RequestIp:                 reqIp,
		RequestHost:               reqHost,
		RequestUri:                reqUri,
		RequestMethod:             reqMethod,
		RequestAgent:              reqAgent,
		ErrorMessage:              message,
		ErrorTrace:                trace,
		ResponseData:              response,
		LogDate:                   time.Now().UTC(),
	}

	go clog.SendLogErrorV1(log)
}
