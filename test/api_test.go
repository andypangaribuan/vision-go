package test

import (
	"fmt"
	"github.com/andypangaribuan/vision-go"
	"github.com/andypangaribuan/vision-go/core/api"
	"github.com/andypangaribuan/vision-go/vis"
	"net/http"
	"testing"
	"time"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/19
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
const (
	apiPort  = 3456
	postPath = "/private/v1"
	baseUrl  = "http://localhost"
)


func init() {
	vision.Initialize()
	vis.Conf.DefaultHttpRequestTimeout = time.Minute
}

func callPost(t *testing.T) {
	time.Sleep(time.Second * 3)
	url := fmt.Sprintf("%v:%v%v", baseUrl, apiPort, postPath)
	t.Logf("url: [post] %v", url)

	data, httpCode, err := vis.Http.Post(url, nil, nil, false, nil)
	logFatal(err)

	if httpCode == 200 {
		t.Logf("success, response: %v", string(data))
	} else {
		t.Logf("failed, code: %v, response: %v", httpCode, string(data))
		t.FailNow()
	}
}

func serverSixSeconds(e *api.EchoApi) {
	done := make(chan bool)
	go e.Serve()
	go func() {
		time.Sleep(time.Second * 6)
		done <- true
	}()
	<-done
}



func Test_ApiWithoutMiddleware(t *testing.T) {
	go callPost(t)

	e := api.BuildEcho(apiPort, nil)

	fn := func(c api.Context) error {
		timeNow := timeNow()
		t.Logf("called: %v", timeNow)
		return c.ResponseStr(http.StatusOK, timeNow)
	}

	// single post
	//e.POST(postPath, fn)

	// group post
	//e.Group("/private", func(g *api.GroupApi) {
	//	g.POST("/v1", fn)
	//})

	// group, using posts
	e.Group("/private", func(g *api.GroupApi) {
		g.POSTS(map[string]api.HandlerFunc{
			"/v1":              fn,
			"/data-management": fn,
		})
	})

	serverSixSeconds(e)
}
