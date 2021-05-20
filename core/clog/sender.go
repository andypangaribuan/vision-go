package clog

import (
	"github.com/andypangaribuan/vision-go/vis"
	"log"
	"sync"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/18
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
var mux sync.RWMutex
var queuePackage []queuePackageModel
var queueActive bool


func IsHaveCLog() bool {
	if vis.Store.CLogBaseUrl == "" || vis.Store.CLogBaseUrl == "-" {
		return false
	}
	return true
}


func sendLogMiddlewareV1(model LogMiddlewareV1) {
	if IsHaveCLog() {
		send(model, "/save/log-middleware/v1")
	}
}


func SendLogTraceV1(model LogTraceV1) {
	if IsHaveCLog() {
		send(model, "/save/log-trace/v1")
	}
}


func SendLogErrorV1(model LogErrorV1) {
	if IsHaveCLog() {
		send(model, "/save/log-error/v1")
	}
}


func SendLogQuery(model LogQueryV1) {
	if IsHaveCLog() {
		send(model, "/save/log-query/v1")
	}
}



func send(model interface{}, subUrl string) {
	if !IsHaveCLog() {
		return
	}

	var errTrace *string
	url := vis.Store.CLogBaseUrl + subUrl

	if data, err := vis.Json.Marshal(model); err != nil {
		errTrace = vis.Log.Stack(url, err)
	} else {
		if vis.Store.CLogMaxConcurrentConnection <= 0 {
			errTrace = postData(url, data, nil)
		} else {
			doQueue(url, data)
		}
	}

	if errTrace != nil {
		vis.Log.Println(*errTrace)
	}
}


func postData(url string, data []byte, wg *sync.WaitGroup) (errTrace *string) {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	httpBlob, httpCode, err := vis.Http.PostData(url, nil, data, true, nil)
	if err != nil {
		errTrace = vis.Log.Stack(url, data, err)
		vis.Log.Println(*errTrace)
	} else {
		if httpCode != 200 {
			vis.Log.Printf("url: %v\nhttp-code: %v\nresponse: %v\n", url, httpCode, string(httpBlob))
		}
	}
	return
}


func doQueue(url string, data []byte) {
	mux.Lock()
	defer mux.Unlock()
	queuePackage = append(queuePackage, queuePackageModel{
		url:  url,
		data: data,
	})

	if !queueActive {
		queueActive = true
		go queueEngine()
	}
}


func queueEngine() {
	printLogQueueEngine("QueueEngine: Active\n")
	for {
		packages := getQueuePackage()
		if len(packages) == 0 {
			break
		}

		vis.Util.ConcurrentProcess(len(packages), vis.Store.CLogMaxConcurrentConnection,
			func(index int, wg *sync.WaitGroup) {
				m := packages[index]
				printLogQueueEngine("QueueEngine: Sent: %v\n", m.url)
				go postData(m.url, m.data, wg)
			},
			func() (continued bool) {
				return true
			})
	}
	printLogQueueEngine("QueueEngine: Shutdown\n\n\n")
}


func getQueuePackage() (packages []queuePackageModel) {
	mux.Lock()
	defer mux.Unlock()

	packages = make([]queuePackageModel, 0)
	max := vis.Store.CLogMaxConcurrentConnection * 10
	lenQueuePackage := len(queuePackage)
	if lenQueuePackage == 0 {
		queueActive = false
	} else {
		if lenQueuePackage < max {
			max = lenQueuePackage
		}

		packages = append(packages, queuePackage[0:max]...)
		if lenQueuePackage > max {
			queuePackage = queuePackage[max:]
		} else {
			queuePackage = make([]queuePackageModel, 0)
		}
	}
	return
}


func printLogQueueEngine(format string, args ...interface{}) {
	if vis.Store.CLogQueueEngineLogPrintEnable {
		log.Printf(format, args...)
	}
}
