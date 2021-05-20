package api

import "github.com/labstack/echo/v4"


/* ============================================
	Created by andy pangaribuan on 2021/05/18
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
type echoContext struct {
	c echo.Context
}

type smLog struct {
	context *Context
}

type Context struct {
	echo *echoContext
	Log smLog
}

type EchoApi struct {
	port int
	e    *echo.Echo
}
