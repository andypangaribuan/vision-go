package test

import (
	"fmt"
	"github.com/andypangaribuan/vision-go"
	"github.com/andypangaribuan/vision-go/models"
	"github.com/andypangaribuan/vision-go/vis"
	"strings"
	"testing"
	"time"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/03
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
func init() {
	vision.Initialize()
}


func Test_ScanDir(t *testing.T) {
	dirPath := getThisDir(t)

	t.Logf(":: SCAN ALL FILE")
	fns := vis.Util.ScanDir(false, true, -1, dirPath, []string{}, func(model models.DirScan) (include bool) {
		if strings.Contains(model.FilePath, vis.Const.PathSeparator + ".git" + vis.Const.PathSeparator) {
			return false
		}
		if strings.Contains(model.FilePath, vis.Const.PathSeparator + ".idea") {
			return false
		}
		return true
	})

	for fn := range fns {
		model, err := fn()
		if err != nil {
			t.Fatalf(":: ERROR\n%+v", err)
			return
		}

		size := vis.Convert.BytesLengthToHumanReadable(model.FileSize, 1)
		t.Logf("- %v \t%v", size, model.FilePath)
	}


	t.Log("")
	t.Log("")
	t.Logf(":: SCAN ALL DIRECTORY")
	fns = vis.Util.ScanDir(true, true, -1, dirPath, []string{}, func(model models.DirScan) (include bool) {
		if !model.IsDirectory {
			return false
		}

		if strings.Contains(model.DirPath, vis.Const.PathSeparator + ".git") {
			return false
		}
		if strings.Contains(model.DirPath, vis.Const.PathSeparator + ".idea") {
			return false
		}
		return true
	})

	for fn := range fns {
		model, err := fn()
		if err != nil {
			t.Fatalf(":: ERROR\n%+v", err)
			return
		}

		t.Logf("- %v", model.DirPath)
	}


	t.Log("")
	t.Log("")
	t.Logf(":: SCAN ALL")
	fns = vis.Util.ScanDir(true, true, -1, dirPath, []string{}, func(model models.DirScan) (include bool) {
		path := ""
		if model.IsDirectory {
			path = model.DirPath
		} else {
			path = model.FilePath
		}

		if strings.Contains(path, vis.Const.PathSeparator + ".git") {
			if strings.Contains(path, ".gitignore") {
				return true
			}
			return false
		}
		if strings.Contains(path, vis.Const.PathSeparator + ".idea") {
			return false
		}
		return true
	})

	for fn := range fns {
		model, err := fn()
		if err != nil {
			t.Fatalf(":: ERROR\n%+v", err)
			return
		}

		if model.IsDirectory {
			t.Logf("- D | %v", model.DirPath)
		} else {
			t.Logf("- F | %v", model.FilePath)
		}
	}
}


func Test_NonBlockingConcurrentProcess(t *testing.T) {
	var (
		processSize          = 50
		maxConcurrentProcess = 10
	)

	t.Logf("size: %v, concurrent: %v", processSize, maxConcurrentProcess)

	go func() {
		time.Sleep(time.Second * 2)
		t.Log("wait idx: 10, 20, 30, 40")
	}()

	vis.Util.NonBlockingConcurrentProcess(processSize, maxConcurrentProcess, func(index int, nb *models.NonBlockConcurrentProcessHolder) {
		if index != 0 && index % 10 == 0 {
			time.Sleep(time.Second * 3)
		}
		t.Logf("- idx: %v", index)
		nb.Done()
	})
}


func Test_CountFileLines(t *testing.T) {
	dirPath := getThisDir(t)
	filePath := dirPath + "LICENSE"

	count, err := vis.Util.CountFileLines(filePath)
	logFatal(err)

	if count == 22 {
		t.Log("success")
	} else {
		t.Logf("failed, count: %v", count)
		t.FailNow()
	}
}


func Test_ScanFileLines(t *testing.T) {
	dirPath := getThisDir(t)
	filePath := dirPath + "LICENSE"

	scan, err := vis.Util.ScanFileLines(filePath)
	logFatal(err)

	for fn := range scan {
		isLast, index, line := fn()
		if !isLast {
			line += "\n"
		}
		fmt.Printf("%v\t%v", index, line)
	}
	println()
}
