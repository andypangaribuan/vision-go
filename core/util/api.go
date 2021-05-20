package util

import (
	"bytes"
	"github.com/andypangaribuan/vision-go/vis"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"io/ioutil"
	"reflect"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/18
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
func (slf *utilStruct) GetTraceId(c echo.Context) (traceId string) {
	req := c.Request()
	traceId = req.Header.Get(vis.Conf.TraceIdKey())
	if traceId == "" {
		traceId = slf.GetId100()
		req.Header.Add(vis.Conf.TraceIdKey(), traceId)
	}
	return
}


func (*utilStruct) GetEchoRequestBody(c echo.Context) (body string, err error) {
	blob, _err := getEchoRequestBodyBlob(c)
	if _err != nil {
		err = _err
		return
	}

	body = string(blob)
	return
}


func (slf *utilStruct) BindEchoBodyToPars(c echo.Context, destination interface{}) error {
	destValue := reflect.ValueOf(destination)
	if destValue.Kind() != reflect.Ptr {
		return errors.New("destination must pass a pointer")
	}
	if destValue.IsNil() {
		return errors.New("can not pass a nil pointer")
	}

	blob, err := getEchoRequestBodyBlob(c)
	if err != nil {
		return err
	}

	err = vis.Json.UnMarshal(blob, destination)
	if err != nil {
		err = errors.WithStack(err)
	}
	return err
}


func getEchoRequestBodyBlob(c echo.Context) (blob []byte, err error) {
	req := c.Request()
	reader := req.Body
	if reader == nil {
		err = errors.New("the request body reader are nil")
		return
	}

	blob, err = ioutil.ReadAll(reader)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	req.Body = ioutil.NopCloser(bytes.NewBuffer(blob))
	return
}
