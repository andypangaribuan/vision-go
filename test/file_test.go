package test

import (
	"github.com/andypangaribuan/vision-go"
	"github.com/andypangaribuan/vision-go/vis"
	"testing"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/03
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
func TestScanFiles(t *testing.T) {
	t.Log("START")
	vision.Initialize()

	dirPath := "/Users/legend/repo/github/vision-go/"
	files, err := vis.Util.ScanFiles(false, -1, dirPath, []string{}, nil)
	if err != nil {
		t.Errorf(":: ERROR\n%+v", err)
		return
	}

	if len(files) > 0 {
		t.Log("FILES:")
	}
	for _, file := range files {
		t.Logf("- %v", file.FilePath)
	}
}
