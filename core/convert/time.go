package convert

import (
	"github.com/andypangaribuan/vision-go/vis"
	"github.com/pkg/errors"
	"strings"
	"time"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/18
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
const (
	layoutTimeDate = "2006-01-02"
	layoutTimeFull = "2006-01-02 15:04:05"
	layoutTimeMillis = "2006-01-02 15:04:05.000"
	layoutTimeMicros = "2006-01-02 15:04:05.000000"
)

func (slf *convertStruct) RemoveUnIntTime(value string) string {
	return vis.Util.RemoveAllUnusedStr(value, "-", " ", ":", ".")
}



func (slf *convertStruct) TimeNowUtcMicros() (dt time.Time, dts string) {
	dt = time.Now().UTC()
	dts = slf.TimeToStrMicros(dt)
	return
}

func (*convertStruct) UnixMicro(val time.Time) int64 {
	return int64(time.Nanosecond) * val.UnixNano() / int64(time.Microsecond)
}

func (slf *convertStruct) UnixMicroNowTimeZone(zone int) int64 {
	return slf.UnixMicro(time.Now().UTC().Add(time.Hour * time.Duration(zone)))
}



func (*convertStruct) TimeToStr(tm time.Time, format string) string {
	replacer := [][]string{
		{"yyyy", "2006"},
		{"MM", "01"},
		{"dd", "02"},
		{"HH", "15"},
		{"mm", "04"},
		{"ss", "05"},
		{"SSSSSS", "000000"},
		{"SSSSS", "00000"},
		{"SSSS", "0000"},
		{"SSS", "000"},
		{"SS", "00"},
		{"S", "0"},
	}

	for _, arr := range replacer  {
		format = strings.Replace(format, arr[0], arr[1], -1)
	}

	return tm.Format(format)
}

func (*convertStruct) TimeToStrDate(tm time.Time) string {
	return tm.Format(layoutTimeDate)
}

func (*convertStruct) TimeToStrFull(tm time.Time) string {
	return tm.Format(layoutTimeFull)
}

func (*convertStruct) TimeToStrMillis(tm time.Time) string {
	return tm.Format(layoutTimeMillis)
}

func (*convertStruct) TimeToStrMicros(tm time.Time) string {
	return tm.Format(layoutTimeMicros)
}



func (*convertStruct) StrToTime(layout string, value string) (tm time.Time, err error) {
	tm, err = time.Parse(layout, value)
	if err != nil {
		err = errors.WithStack(err)
	}
	return
}

func (slf *convertStruct) StrToTimeDate(value string) (tm time.Time, err error) {
	return slf.StrToTime(layoutTimeDate, value)
}

func (slf *convertStruct) StrToTimeFull(value string) (tm time.Time, err error) {
	return slf.StrToTime(layoutTimeFull, value)
}

func (slf *convertStruct) StrToTimeMillis(value string) (tm time.Time, err error) {
	return slf.StrToTime(layoutTimeMillis, value)
}

func (slf *convertStruct) StrToTimeMicros(value string) (tm time.Time, err error) {
	return slf.StrToTime(layoutTimeMicros, value)
}
