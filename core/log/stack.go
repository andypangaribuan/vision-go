package log

import (
	"fmt"
	"github.com/andypangaribuan/vision-go/models"
	"github.com/andypangaribuan/vision-go/vis"
	"os"
	"reflect"
	"runtime"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/18
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
func (slf *logStruct) Stack(args ...interface{}) (stack *string) {
	v := slf.BaseStack(2, args...)
	stack = &v
	return
}


func (*logStruct) BaseStack(skip int, args ...interface{}) (stack string) {
	pc, filePath, lineNumber, _ := runtime.Caller(skip)
	funcName := runtime.FuncForPC(pc).Name()

	format := ":: %s \n:: %s:%d"
	data := []interface{}{funcName, filePath, lineNumber}

	for _, arg := range args {
		switch v := arg.(type) {
		case error:
			format += "\n:: %v"
			data = append(data, v)
		}
	}

	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		if codes := readFile(filePath, lineNumber); codes != "" {
			format += "\n:: START CODE STACK\n"
			format += "%v"
			format += "\n:: END CODE STACK"

			data = append(data, codes)
		}
	}

	for _, v := range args {
		format += "\n\n>_\n"
		format += "%+v"

		data = append(data, getLogValue(v))
	}

	format += "\n\n|=|"
	stack = fmt.Sprintf(format, data...)
	stack = vis.Util.RemoveInvalidChar(stack)
	return
}


func getLogValue(obj interface{}) string {
	if obj == nil {
		return "nil"
	}

	objRef := reflect.ValueOf(obj)
	objKind := objRef.Kind()
	if objKind == reflect.Ptr {
		objRef = objRef.Elem()
		objKind = objRef.Kind()
	}

	if objKind == reflect.Invalid {
		return "nil"
	}
	obj = objRef.Interface()

	value := ""
	switch data := obj.(type) {
	case string: value = data
	case []byte: value = string(data)
	case error: value = codeStack(data) + "\n\n" + fmt.Sprintf("%+v", data)
	case models.DbTxError:
		if data.Msg != "" {
			value = data.Msg
		}
		if data.Err != nil {
			if value != "" {
				value += "\n\n"
			}
			value += codeStack(data.Err) + "\n\n" + fmt.Sprintf("%+v", data.Err)
		}
	default:
		if content, err := vis.Json.Encode(data); err == nil && content != "" {
			value = content
		}
		if value == "" || value == "{}" {
			value = fmt.Sprintf("%+v", data)
		}
	}

	return value
}
