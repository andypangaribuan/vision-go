package vision

import (
	"github.com/andypangaribuan/vision-go/core/convert"
	"github.com/andypangaribuan/vision-go/core/db"
	"github.com/andypangaribuan/vision-go/core/env"
	"github.com/andypangaribuan/vision-go/core/http"
	"github.com/andypangaribuan/vision-go/core/json"
	"github.com/andypangaribuan/vision-go/core/log"
	"github.com/andypangaribuan/vision-go/core/term"
	"github.com/andypangaribuan/vision-go/core/util"
	"github.com/andypangaribuan/vision-go/core/val"
	"github.com/andypangaribuan/vision-go/vis"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/03
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
//goland:noinspection GoUnusedExportedFunction
func Initialize() {
	vis.Db = db.Db
	vis.Env = env.Env
	vis.Http = http.Http
	vis.Term = term.Term
	vis.Json = json.Json
	vis.Log = log.Log
	vis.Val = val.Val
	vis.Convert = convert.Convert
	vis.Util = util.Util
}
