package vis

import "time"


/* ============================================
	Created by andy pangaribuan on 2021/05/03
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
type constVariable struct {
	PathSeparator string
}

type storeStruct struct {
	ServiceName                   string
	ServiceVersion                string
	CLogBaseUrl                   string
	CLogMaxConcurrentConnection   int
	CLogQueueEngineLogPrintEnable bool
	LogPrintEnabled               bool
}

type confStruct struct {
	DefaultHttpRequestTimeout time.Duration // default: 3 minute
}
