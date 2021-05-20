package vis


/* ============================================
	Created by andy pangaribuan on 2021/05/18
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
func (*confStruct) TraceIdKey() string {
	return "VISION_API_TRACE_ID"
}

func (*confStruct) RequestUidKey() string {
	return "VISION_API_REQUEST_UID"
}

func (*confStruct) RequestFromServiceName() string {
	return "VISION_API_REQUEST_FROM_SERVICE_NAME"
}

func (*confStruct) RequestFromServiceVersion() string {
	return "VISION_API_REQUEST_FROM_SERVICE_VERSION"
}
