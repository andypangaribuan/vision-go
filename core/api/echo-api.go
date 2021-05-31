package api

import (
	"github.com/andypangaribuan/vision-go/core/clog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/18
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
//noinspection GoUnusedExportedFunction
func BuildEcho(port int, logMiddleware clog.EchoMiddleware) *EchoApi {
	e := echo.New()

	if logMiddleware != nil {
		e.Use(logMiddleware.Logger)
	}

	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	})

	if logMiddleware != nil {
		if v, ok := logMiddleware.(clog.EchoMiddlewareV1); ok {
			if v.ProfilerIncluded {
				g := e.Group("/debug/pprof")
				g.GET("/*", echo.WrapHandler(http.DefaultServeMux))
			}
		}
	}

	sm := EchoApi{port, e}
	return &sm
}


func (slf *EchoApi) Serve() {
	slf.e.Logger.Fatal(slf.e.Start(":" + strconv.Itoa(slf.port)))
}


func (slf *EchoApi) POST(path string, handler HandlerFunc) {
	slf.e.POST(path, func(c echo.Context) error {
		ctx := Context{
			echo: &echoContext{c: c},
		}
		ctx.Log = smLog{context: &ctx}

		return handler(ctx)
	})
}

func (slf *EchoApi) POSTS(endpoint map[string]HandlerFunc) {
	for k, v := range endpoint {
		slf.POST(k, v)
	}
}


func (slf *EchoApi) Group(path string, group func(g *GroupApi)) {
	g := &GroupApi{
		ea:   slf,
		path: path,
	}
	group(g)
}

func (slf *GroupApi) POST(path string, handler HandlerFunc) {
	slf.ea.POST(slf.path + path, handler)
}

func (slf *GroupApi) POSTS(endpoint map[string]HandlerFunc) {
	for k, v := range endpoint {
		slf.POST(k, v)
	}
}

func (slf *GroupApi) Group(path string, group func(g *GroupApi)) {
	g := &GroupApi{
		ea:   slf.ea,
		path: slf.path + path,
	}
	group(g)
}
