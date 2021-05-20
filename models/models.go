package models


/* ============================================
	Created by andy pangaribuan on 2021/05/03
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
type APIGeneralResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}


type DirScan struct {
	IsDirectory              bool
	FileName                 string // e.g: file-x.yz
	FileExtension            string // e.g: yz
	FilePath                 string // e.g: /dir-a/dir-b/file-x.yz
	FileNameWithoutExtension string // e.g: file-x
	FileSize                 int64  // in bytes
	DirName                  string // e.g: dir-b
	DirPath                  string // e.g: /dir-a/dir-b
}


type DbTxError struct {
	Err error
	Msg string
}

type DbUnsafeSelectError struct {
	LogType    string
	SqlQuery   string
	SqlPars    []interface{}
	LogMessage *string
	LogTrace   *string
}

type NonBlockConcurrentProcessHolder struct {
	Done func()
}
