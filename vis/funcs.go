package vis

import (
	"github.com/andypangaribuan/vision-go/models"
	"github.com/labstack/echo/v4"
	"io"
	"sync"
	"time"
)



/* ============================================
	Created by andy pangaribuan on 2021/05/03
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
// package: code/db
type iDb interface {
	BuildMySQLConnStr(host string, port int, dbName, username, password string) string
	BuildPostgreSQLConnStr(host string, port int, dbName, username, password string, schema *string) string
	CreateInstance(dbType string, masterConnStr, slaveConnStr *string, maxLifeTimeConn, maxIdleConn, maxOpenConn int, autoRebind, unsafeCompatibility bool) (master models.IDbMaster, slave models.IDbSlave)
}


// package: core/env
type iEnv interface {
	GetStr(key string) string
	GetIntEnv(key string) int
	GetBoolEnv(key string) bool
}


// package: core/http
type iHttp interface {
	Get(url string, header map[string]string, params map[string]string, isSkipSecurityChecking bool, timeOut *time.Duration) ([]byte, int, error)
	Post(url string, header map[string]string, body map[string]interface{}, isSkipSecurityChecking bool, timeOut *time.Duration) ([]byte, int, error)
	PostData(url string, header map[string]string, data []byte, isSkipSecurityChecking bool, timeOut *time.Duration) ([]byte, int, error)
}


// package: core/term
type iTerm interface {
	Move(line, cursor int)
	MoveUp(inc int)
	MoveDown(inc int)
	MoveLeft(length int)
	MoveRight(length int)
	ClearScreen()
	EraseCurrentLine()
	ReplaceCurrentLine(format string, args ...interface{})
}


// package: core/json
type iJson interface {
	Marshal(obj interface{}) ([]byte, error)
	UnMarshal(data []byte, out interface{}) error
	Encode(obj interface{}) (string, error)
	Decode(jsonStr string, out interface{}) error
	MapToJson(maps map[string]interface{}) (string, error)
}


// package: core/log
type iLog interface {
	Printf(format string, args ...interface{})
	Println(format string, args ...interface{})
	Stack(args ...interface{}) (stack *string)
	BaseStack(skip int, args ...interface{}) (stack string)
}


// package: core/val
type iVal interface {
	StrTrimSpace(value string) string
	PtrStrTrimSpace(value *string) *string
	StrTrimMaxLength(value string, maxLength int) string
	PtrStrTrimMaxLength(value *string, maxLength int) *string
	StrSplitMaxDbText(dbType, value string) (txt string, arr []string)
	PtrStrSplitMaxDbText(dbType string, value *string) (txt *string, arr []string)
}


// package: core/convert
type iConvert interface {
	// fr: convert.go
	BytesLengthToHumanReadable(length int64, decimal int) string

	// fr: time.go
	RemoveUnIntTime(value string) string
	TimeNowUtcMicros() (dt time.Time, dts string)
	UnixMicro(val time.Time) int64
	UnixMicroNowTimeZone(zone int) int64
	TimeToStr(tm time.Time, format string) string
	TimeToStrDate(tm time.Time) string
	TimeToStrFull(tm time.Time) string
	TimeToStrMillis(tm time.Time) string
	TimeToStrMicros(tm time.Time) string
	StrToTime(layout string, value string) (tm time.Time, err error)
	StrToTimeDate(value string) (tm time.Time, err error)
	StrToTimeFull(value string) (tm time.Time, err error)
	StrToTimeMillis(value string) (tm time.Time, err error)
	StrToTimeMicros(value string) (tm time.Time, err error)
}


// package: core/util
type iUtil interface {
	// fr: api.go
	GetTraceId(c echo.Context) (traceId string)
	GetEchoRequestBody(c echo.Context) (body string, err error)
	BindEchoBodyToPars(c echo.Context, destination interface{}) error

	// fr: files.go
	GetFileExtension(fileName string) string
	ScanDir(includeDirectory, recursively bool, limit int, directoryPath string, extensions []string, condition func(model models.DirScan) (include bool)) (ch chan func()(*models.DirScan, error))
	CountFileLines(filePath string) (count int, err error)
	ScanFileLines(filePath string) (scan chan func()(isLast bool, index int, line string), err error)

	// fr: general.go
	RemoveAllUnusedStr(value string, args ...string) string
	GetUUID() string
	GetUniqueId() string
	GetId100() (id string)
	GetId10() (id string)
	GetId20() (id string)
	GetId30() (id string)
	RandomAlphabet(length int) string
	RandomLowerCaseAlphabet(length int) string
	RandomUpperCaseAlphabet(length int) string
	RandomNumeric(length int) string
	RandomAlphabetNumeric(length int) string
	RandomLowerCaseAlphabetNumeric(length int) string
	RandomUpperCaseAlphabetNumeric(length int) string
	Base64EncodeString(value string) string
	Base64DecodeString(value string) (decode string, err error)
	MD5Encryption(value string) string
	ParsValidationNullOrEmpty(names []string, args ...interface{}) (msg *string)
	StrIsDigitAndLetter(value string) bool
	Slicing(listLength, chunkSize int, funcBatch func(fr, to int))
	ConcurrentProcess(processSize, maxConcurrentProcess int, doFunc func(index int, wg *sync.WaitGroup), doneWaitFunc func() (continued bool))
	NonBlockingConcurrentProcess(processSize, maxConcurrentProcess int, doFunc func(index int, nb *models.NonBlockConcurrentProcessHolder))
	ReaderReadAll(r io.Reader) (data []byte, err error)
	RemoveInvalidChar(value string) string
}
