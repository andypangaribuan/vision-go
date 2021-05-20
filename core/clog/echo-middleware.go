package clog

import (
	"encoding/base64"
	"fmt"
	"github.com/andypangaribuan/vision-go/vis"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/18
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
func (slf EchoMiddlewareV1) Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		if !IsHaveCLog() {
			err = next(c)
			return
		}

		req := c.Request()
		res := c.Response()

		model := LogMiddlewareV1{
			ServiceName:      vis.Store.ServiceName,
			ServiceVersion:   vis.Store.ServiceVersion,
			TraceId:          vis.Util.GetTraceId(c),
			RequestIp:        c.RealIP(),
			RequestHost:      req.Host,
			RequestUri:       req.RequestURI,
			RequestMethod:    req.Method,
			RequestAgent:     req.UserAgent(),
			RequestTimeStart: time.Now().UTC(),
		}

		if !slf.ignoreThisLogRequestBody(model.RequestUri) {
			if val, err := vis.Util.GetEchoRequestBody(c); err == nil {
				model.RequestBody = &val
			}
		}


		err = next(c)
		if err != nil {
			c.Error(err)
			err = errors.WithStack(err)
		}

		if err != nil {
			val := fmt.Sprintf("%+v", err)
			model.RequestError = &val
		}


		getHeader := func(key string) (val string) {
			val = req.Header.Get(key)
			if val != "" {
				req.Header.Del(key)
			}
			return
		}

		if val := getHeader(vis.Conf.RequestFromServiceName()); val != "" {
			model.RequestFromServiceName = &val
		}
		if val := getHeader(vis.Conf.RequestFromServiceVersion()); val != "" {
			model.RequestFromServiceVersion = &val
		}
		if val := getHeader(vis.Conf.RequestUidKey()); val != "" {
			model.RequestUid = &val
		}

		if val := getHeader("ERQVA"); val != "" {
			if blob, err := base64.StdEncoding.DecodeString(val); err == nil {
				var maps map[string]interface{}
				if err := vis.Json.UnMarshal(blob, &maps); err == nil {
					for k,v := range maps {
						switch val := v.(type) {
						case string:
							switch k {
							case "1": model.RequestAppPackage = &val
							case "2": model.RequestAppVersion = &val
							case "4": model.RequestOsName = &val
							case "5": model.RequestOsVersion = &val
							}
						default:
							switch k {
							case "3": if val, err := strconv.Atoi(fmt.Sprintf("%v", v)); err == nil {
								model.RequestAppBuildNumber = &val
							}
							}
						}
					}
				}
			}
		}

		if val := getHeader("ERQVB"); val != "" {
			if blob, err := base64.StdEncoding.DecodeString(val); err == nil {
				var maps map[string]interface{}
				if err := vis.Json.UnMarshal(blob, &maps); err == nil {
					for k,v := range maps {
						switch val := v.(type) {
						case string:
							switch k {
							case "3": model.RequestFcmToken = &val
							}
						default:
							switch k {
							case "1":
								if val, ok := v.(float64); ok {
									model.RequestLocationLat = &val
								}
							case "2":
								if val, ok := v.(float64); ok {
									model.RequestLocationLong = &val
								}
							}
						}
					}
				}
			}
		}

		if val := getHeader("ERQVC"); val != "" {
			if blob, err := base64.StdEncoding.DecodeString(val); err == nil {
				var maps map[string]interface{}
				if err := vis.Json.UnMarshal(blob, &maps); err == nil {
					for k,v := range maps {
						switch val := v.(type) {
						case string:
							switch k {
							case "1": model.RequestDeviceId = &val
							case "2": model.RequestBrandName = &val
							case "3": model.RequestDeviceModel = &val
							}
						default:
							switch k {
							case "4": if val, err := strconv.Atoi(fmt.Sprintf("%v", v)); err == nil {
								model.RequestFromPhysicalDevice = &val
							}
							}
						}
					}
				}
			}
		}


		if val := fmt.Sprint(req.Header); val != "" {
			model.RequestHeader = &val
		}

		if v := req.Header.Get(echo.HeaderContentLength); v != "" {
			if val, err := strconv.ParseInt(v, 10, 64); err == nil {
				model.RequestBytes = val
			}
		}

		model.RequestTimeFinish = time.Now().UTC()
		model.ResponseStatus = res.Status
		model.ResponseBytes = res.Size
		model.LogDate = time.Now().UTC()

		go sendLogMiddlewareV1(model)

		return
	}
}


func (slf EchoMiddlewareV1) ignoreThisLogRequestBody(uri string) (ignore bool) {
	uri = strings.ToLower(uri)
	for _, ignoreUri := range slf.IgnoreLogRequestBodyUris {
		if uri == strings.ToLower(ignoreUri) {
			ignore = true
			break
		}
	}
	return
}
