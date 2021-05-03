package models


/* ============================================
	Created by andy pangaribuan on 2021/05/03
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
type FileScan struct {
	FileName                 string	// e.g: file-x.yz
	Extension                string	// e.g: yz
	FilePath                 string // e.g: /dir-a/dir-b/file-x.yz
	FileNameWithoutExtension string	// e.g: file-x
	DirName                  string	// e.g: dir-b
	DirPath                  string	// e.g: /dir-a/dir-b
}
