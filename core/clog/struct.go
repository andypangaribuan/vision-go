package clog

import (
	"github.com/labstack/echo/v4"
	"time"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/18
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
type queuePackageModel struct {
	url  string
	data []byte
}


type EchoMiddleware interface {
	Logger(next echo.HandlerFunc) echo.HandlerFunc
}

type EchoMiddlewareV1 struct {
	IgnoreLogRequestBodyUris []string
	ProfilerIncluded         bool
}


type LogTraceV1 struct {
	ServiceName               string    `json:"service_name"`
	ServiceVersion            string    `json:"service_version"`
	TraceId                   string    `json:"trace_id"`
	RequestFromServiceName    *string   `json:"request_from_service_name"`
	RequestFromServiceVersion *string   `json:"request_from_service_version"`
	RequestUid                *string   `json:"request_uid"`
	LogType                   string    `json:"log_type"`
	LogMessage                string    `json:"log_message"`
	LogTrace                  *string   `json:"log_trace"`
	LogDate                   time.Time `json:"log_date" jdt:"micros"`
}


type LogErrorV1 struct {
	ServiceName               string    `json:"service_name"`
	ServiceVersion            string    `json:"service_version"`
	TraceId                   string    `json:"trace_id"`
	RequestFromServiceName    *string   `json:"request_from_service_name"`
	RequestFromServiceVersion *string   `json:"request_from_service_version"`
	RequestUid                *string   `json:"request_uid"`
	RequestIp                 *string   `json:"request_ip"`
	RequestHost               *string   `json:"request_host"`
	RequestUri                *string   `json:"request_uri"`
	RequestMethod             *string   `json:"request_method"`
	RequestAgent              *string   `json:"request_agent"`
	ErrorMessage              *string   `json:"error_message"`
	ErrorTrace                string    `json:"error_trace"`
	ResponseData              *string   `json:"response_data"`
	LogDate                   time.Time `json:"log_date" jdt:"micros"`
}


type LogQueryV1 struct {
	ServiceName    string    `json:"service_name"`
	ServiceVersion string    `json:"service_version"`
	TraceId        string    `json:"trace_id"`
	LogType        string    `json:"log_type"`
	SqlQuery       string    `json:"sql_query"`
	SqlParams      *string   `json:"sql_params"`
	LogMessage     *string   `json:"log_message"`
	LogTrace       *string   `json:"log_trace"`
	LogDate        time.Time `json:"log_date" jdt:"micros"`
}


type LogMiddlewareV1 struct {
	ServiceName               string    `json:"service_name"`
	ServiceVersion            string    `json:"service_version"`
	TraceId                   string    `json:"trace_id"`
	RequestFromServiceName    *string   `json:"request_from_service_name"`
	RequestFromServiceVersion *string   `json:"request_from_service_version"`
	RequestUid                *string   `json:"request_uid"`
	RequestIp                 string    `json:"request_ip"`
	RequestHost               string    `json:"request_host"`
	RequestUri                string    `json:"request_uri"`
	RequestMethod             string    `json:"request_method"`
	RequestAgent              string    `json:"request_agent"`
	RequestAppPackage         *string   `json:"request_app_package"`          // EXT-RQV-A-1
	RequestAppVersion         *string   `json:"request_app_version"`          // EXT-RQV-A-2
	RequestAppBuildNumber     *int      `json:"request_app_build_number"`     // EXT-RQV-A-3
	RequestOsName             *string   `json:"request_os_name"`              // EXT-RQV-A-4
	RequestOsVersion          *string   `json:"request_os_version"`           // EXT-RQV-A-5
	RequestLocationLat        *float64  `json:"request_location_lat"`         // EXT-RQV-B-1
	RequestLocationLong       *float64  `json:"request_location_long"`        // EXT-RQV-B-2
	RequestFcmToken           *string   `json:"request_fcm_token"`            // EXT-RQV-B-3
	RequestDeviceId           *string   `json:"request_device_id"`            // EXT-RQV-C-1
	RequestBrandName          *string   `json:"request_device_brand"`         // EXT-RQV-C-2
	RequestDeviceModel        *string   `json:"request_device_model"`         // EXT-RQV-C-3
	RequestFromPhysicalDevice *int      `json:"request_from_physical_device"` // EXT-RQV-C-4
	RequestHeader             *string   `json:"request_header"`
	RequestBody               *string   `json:"request_body"`
	RequestError              *string   `json:"request_error"`
	RequestBytes              int64     `json:"request_bytes"`
	RequestTimeStart          time.Time `json:"request_time_start" jdt:"micros"`
	RequestTimeFinish         time.Time `json:"request_time_finish" jdt:"micros"`
	ResponseStatus            int       `json:"response_status"`
	ResponseBytes             int64     `json:"response_bytes"`
	LogDate                   time.Time `json:"log_date" jdt:"micros"`
}
