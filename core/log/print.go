package log

import (
	"github.com/andypangaribuan/vision-go/vis"
	"log"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/18
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
func (*logStruct) Printf(format string, args ...interface{}) {
	if vis.Store.LogPrintEnabled {
		log.Printf(format, args...)
	}
}

func (slf *logStruct) Println(format string, args ...interface{}) {
	slf.Printf(format + "\n", args...)
}
