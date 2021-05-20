package http

import "net/http"


/* ============================================
	Created by andy pangaribuan on 2021/05/19
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
var Http *httpStruct


func init() {
	Http = &httpStruct{}
}


type httpStruct struct { }


type visHttpClient struct {
	client http.Client
}
