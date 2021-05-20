package vis

import (
	"os"
	"time"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/03
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
//goland:noinspection GoUnusedGlobalVariable
var (
	Conf  *confStruct
	Const *constVariable
	Store *storeStruct

	Db      iDb
	Env     iEnv
	Http    iHttp
	Term    iTerm
	Json    iJson
	Log     iLog
	Val     iVal
	Convert iConvert
	Util    iUtil
)


func init() {
	Conf = &confStruct{
		DefaultHttpRequestTimeout: time.Minute * 3,
	}

	Const = &constVariable{
		PathSeparator: string(os.PathSeparator),
	}

	Store = &storeStruct{
		ServiceName:                   "vision:default",
		ServiceVersion:                "0.0.0",
		CLogBaseUrl:                   "-",
		CLogMaxConcurrentConnection:   0,
		CLogQueueEngineLogPrintEnable: false,
		LogPrintEnabled:               true,
	}
}
