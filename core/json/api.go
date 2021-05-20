package json

import (
	"github.com/json-iterator/go"
	"github.com/pkg/errors"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/18
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
var api jsoniter.API

func init() {
	api = ConfigWithCustomTimeFormat

	SetDefaultTimeFormat("2006-01-02 15:04:05.000000", nil)
	AddLocaleAlias("-", nil)

	AddTimeFormatAlias("date", "2006-01-02")
	AddTimeFormatAlias("time", "15:04:05")
	AddTimeFormatAlias("dt", "2006-01-02 15:04:05")
	AddTimeFormatAlias("millis", "2006-01-02 15:04:05.000")
	AddTimeFormatAlias("micros", "2006-01-02 15:04:05.000000")
}


func (*jsonStruct) Marshal(obj interface{}) ([]byte, error) {
	data, err := api.Marshal(obj)
	if err != nil {
		err = errors.WithStack(err)
	}
	return data, err
}

func (*jsonStruct) UnMarshal(data []byte, out interface{}) error {
	err := api.Unmarshal(data, &out)
	if err != nil {
		err = errors.WithStack(err)
	}
	return err
}

func (*jsonStruct) Encode(obj interface{}) (string, error) {
	val, err := api.MarshalToString(obj)
	if err != nil {
		err = errors.WithStack(err)
	}
	return val, err
}

func (*jsonStruct) Decode(jsonStr string, out interface{}) error {
	err := api.UnmarshalFromString(jsonStr, &out)
	if err != nil {
		err = errors.WithStack(err)
	}
	return err
}

func (*jsonStruct) MapToJson(maps map[string]interface{}) (string, error) {
	val, err := api.MarshalToString(maps)
	if err != nil {
		err = errors.WithStack(err)
	}
	return val, err
}
