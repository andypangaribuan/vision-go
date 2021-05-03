package vis

import "github.com/andypangaribuan/vision-go/models"


/* ============================================
	Created by andy pangaribuan on 2021/05/03
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
type constVariable struct {
	PathSeparator string
}


type iUtil interface {
	GetFileExtension(fileName string) string
	ScanFiles(recursively bool, limit int, dirPath string, extensions []string, condition func(file models.FileScan) (include bool)) (files []models.FileScan, err error)
}
