package util

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/andypangaribuan/vision-go/models"
	"github.com/andypangaribuan/vision-go/vis"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"io"
	"math/rand"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"time"
	"unicode"
	"unicode/utf8"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/18
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
const charsetLowerCase = "abcdefghijklmnopqrstuvwxyz"
const charsetUpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const charsetNumeric = "0123456789"

var (
	strIsDigitAndLetterRegex = regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString
	seededRand               = rand.New(rand.NewSource(time.Now().UnixNano()))
	randMux                  sync.Mutex
)


func getRand(f func(rn *rand.Rand)) {
	randMux.Lock()
	defer randMux.Unlock()
	f(seededRand)
}


func (*utilStruct) RemoveAllUnusedStr(value string, args ...string) string {
	for _, arg := range args {
		value = strings.Replace(value, arg, "", -1)
	}
	return value
}



func (slf *utilStruct) GetUUID() string {
	id := uuid.NewV4().String()
	id = slf.RemoveAllUnusedStr(id, "-")
	return id // length: 32
}


func (slf *utilStruct) GetUniqueId() string {
	tm := vis.Convert.TimeToStrMicros(time.Now().UTC())
	tm = slf.RemoveAllUnusedStr(tm, "-", " ", ":", ".")

	y1 := tm[:2]
	y2 := tm[2:4]
	M := tm[4:6]
	d := tm[6:8]
	h := tm[8:10]
	m := tm[10:12]
	s := tm[12:14]
	mc1 := tm[14:17]
	mc2 := tm[17:]

	id := mc1 + s + y1 + h + M + m + y2 + d + mc2
	return id
}

func (slf *utilStruct) GetId100() (id string) {
	id = slf.GetUUID() + slf.GetUniqueId() + slf.RandomAlphabetNumeric(30) + slf.RandomNumeric(18)
	if len(id) > 100 {
		id = id[:100]
	}
	return
}

func (slf *utilStruct) GetId10() (id string) {
	id = slf.GetUUID()
	if len(id) > 10 {
		id = id[:10]
	}
	return
}

func (slf *utilStruct) GetId20() (id string) {
	id = slf.GetUUID()
	if len(id) > 20 {
		id = id[:20]
	}
	return
}

func (slf *utilStruct) GetId30() (id string) {
	id = slf.GetUUID()
	if len(id) > 30 {
		id = id[:30]
	}
	return
}



func (*utilStruct) RandomAlphabet(length int) string {
	return randomCharset(length, charsetLowerCase + charsetUpperCase)
}

func (*utilStruct) RandomLowerCaseAlphabet(length int) string {
	return randomCharset(length, charsetLowerCase)
}

func (*utilStruct) RandomUpperCaseAlphabet(length int) string {
	return randomCharset(length, charsetUpperCase)
}

func (*utilStruct) RandomNumeric(length int) string {
	return randomCharset(length, charsetNumeric)
}

func (*utilStruct) RandomAlphabetNumeric(length int) string {
	return randomCharset(length, charsetLowerCase + charsetUpperCase + charsetNumeric)
}

func (*utilStruct) RandomLowerCaseAlphabetNumeric(length int) string {
	return randomCharset(length, charsetLowerCase + charsetNumeric)
}

func (*utilStruct) RandomUpperCaseAlphabetNumeric(length int) string {
	return randomCharset(length, charsetUpperCase + charsetNumeric)
}

func randomCharset(length int, charset string) string {
	b := make([]byte, length)
	l := len(charset)
	getRand(func(rn *rand.Rand) {
		for i:= range b {
			b[i] = charset[rn.Intn(l)]
		}
	})
	return string(b)
}



func (*utilStruct) Base64EncodeString(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}

func (*utilStruct) Base64DecodeString(value string) (decode string, err error) {
	if blob, _err := base64.StdEncoding.DecodeString(value); _err == nil {
		decode = string(blob)
	} else {
		err = errors.WithStack(_err)
	}
	return
}

func (*utilStruct) MD5Encryption(value string) string {
	hash := md5.New()
	hash.Write([]byte(value))
	return hex.EncodeToString(hash.Sum(nil))
}



func (*utilStruct) ParsValidationNullOrEmpty(names []string, args ...interface{}) (msg *string) {
	pars := ""
	argsLength := len(args)

	for i, name := range names {
		isEmptyOrNull := true

		if i < argsLength {
			objRef := reflect.ValueOf(args[i])
			objKind := objRef.Kind()
			if objKind == reflect.Ptr {
				objRef = objRef.Elem()
				objKind = objRef.Kind()
			}

			if objKind != reflect.Invalid {
				switch objKind {
				case reflect.String:
					v := objRef.String()
					v = strings.TrimSpace(v)
					if v != "" {
						isEmptyOrNull = false
					}
					break
				case reflect.Slice, reflect.Array, reflect.Map:
					length := objRef.Len()
					if length > 0 {
						isEmptyOrNull = false
					}
					break
				default:
					isEmptyOrNull = false
					break
				}
			}
		}

		if isEmptyOrNull {
			if pars != "" {
				pars += ", "
			}
			pars += name
		}
	}

	if pars != "" {
		pars = "empty parameters: " + pars
		msg = &pars
	}
	return
}

func (*utilStruct) StrIsDigitAndLetter(value string) bool {
	return strIsDigitAndLetterRegex(value)
}



func (*utilStruct) Slicing(listLength, chunkSize int, funcBatch func(fr, to int)) {
	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}

	for i:=0; i<listLength; i+=chunkSize {
		funcBatch(i, min(i+chunkSize, listLength))
	}
}



func (*utilStruct) ConcurrentProcess(processSize, maxConcurrentProcess int, doFunc func(index int, wg *sync.WaitGroup), doneWaitFunc func() (continued bool)) {
	var wg sync.WaitGroup
	activeProcess := 0

	for i:=0; i<processSize; i++ {
		wg.Add(1)
		doFunc(i, &wg)
		activeProcess++

		wait := activeProcess == maxConcurrentProcess
		if !wait && i+1 == processSize {
			wait = true
		}

		if wait {
			wg.Wait()
			activeProcess = 0
			if !doneWaitFunc() {
				break
			}
		}
	}
}

func (*utilStruct) NonBlockingConcurrentProcess(processSize, maxConcurrentProcess int, doFunc func(index int, nb *models.NonBlockConcurrentProcessHolder)) {
	if processSize < 1 {
		return
	}

	finished := make(chan bool)
	var mux sync.Mutex
	var holder *models.NonBlockConcurrentProcessHolder
	onBlock := 0
	count := 0

	holder = &models.NonBlockConcurrentProcessHolder{
		Done: func() {
			mux.Lock()
			defer mux.Unlock()

			if count < processSize {
				count++
				go doFunc(count-1, holder)
			} else {
				onBlock--
			}

			if onBlock == 0 {
				finished <- true
			}
		},
	}

	startFirst := func() {
		mux.Lock()
		defer mux.Unlock()

		if maxConcurrentProcess < 1 {
			holder = &models.NonBlockConcurrentProcessHolder{}
			for i:=0; i<processSize; i++ {
				go doFunc(i, holder)
			}
		} else {
			for i:=0; i<processSize; i++ {
				if i == maxConcurrentProcess {
					break
				}

				onBlock++
				count++
				go doFunc(i, holder)
			}
		}
	}

	go startFirst()
	<-finished
}



func (*utilStruct) ReaderReadAll(r io.Reader) (data []byte, err error) {
	var buf bytes.Buffer
	defer buf.Reset()

	buf.Grow(512)
	_, err = buf.ReadFrom(r)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}

	data = buf.Bytes()
	return
}

func (*utilStruct) RemoveInvalidChar(value string) string {
	/* REMOVE INVALID UTF8 CHAR */
	if !utf8.ValidString(value) {
		v := make([]rune, 0, len(value))
		for i, r := range value {
			if r == utf8.RuneError {
				_, size := utf8.DecodeRuneInString(value[i:])
				if size == 1 {
					continue
				}
			}
			v = append(v, r)
		}
		value = string(v)
	}

	/* REMOVE NON PRINTABLE CHAR */
	clean := strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}

		sc := fmt.Sprintf("%q", r)
		scn := len(sc)

		// find hex character code
		// https://golang.org/pkg/regexp/syntax/
		if scn >= 6 && sc[:1] == "'" && sc[scn-1:] == "'" && sc[1:3] == "\\x" {
			return -1
		}

		return r
	}, value)

	return clean
}
