package val

import (
	"strings"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/18
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
func (*valStruct) StrTrimSpace(value string) string {
	return strings.TrimSpace(value)
}


func (*valStruct) PtrStrTrimSpace(value *string) *string {
	var newValue *string

	if value != nil {
		v := *value
		v = strings.TrimSpace(v)
		newValue = &v
	}

	return newValue
}


func (*valStruct) StrTrimMaxLength(value string, maxLength int) string {
	if len(value) > maxLength {
		return value[:maxLength]
	}
	return value
}


func (*valStruct) PtrStrTrimMaxLength(value *string, maxLength int) *string {
	if value != nil && len(*value) > maxLength {
		v := *value
		v = v[:maxLength]
		return &v
	}
	return value
}


func (*valStruct) StrSplitMaxDbText(dbType, value string) (txt string, arr []string) {
	txt = value
	arr = make([]string, 0)

	if value == "" {
		return
	}

	switch dbType {
	case "mysql":
		val := value
		for {
			length := len(val)
			if length == 0 {
				break
			}

			split := val
			if length > 65000 {
				split = val[:65000]
			}
			arr = append(arr, split)
			val = val[len(split):]
		}
	}

	return
}


func (slf *valStruct) PtrStrSplitMaxDbText(dbType string, value *string) (txt *string, arr []string) {
	if value == nil {
		return
	}

	v, arr := slf.StrSplitMaxDbText(dbType, *value)
	return &v, arr
}
