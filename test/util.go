package test

import (
	"fmt"
	"github.com/andypangaribuan/vision-go/vis"
	"os"
	"strings"
	"testing"
	"time"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/19
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
func timeNow() string {
	return vis.Convert.TimeToStrFull(time.Now())
}

func logFatal(err error) {
	if err != nil {
		fmt.Printf("\n\n:: ERROR\n%+v\n", err)
		os.Exit(1)
	}
}

func getThisDir(t *testing.T) string {
	dirPath, err := os.Getwd()
	if err != nil {
		t.Fatalf(":: ERROR\n%+v", err)
	}
	if idx := strings.LastIndex(dirPath, "test"); idx != -1 {
		dirPath = dirPath[:idx]
	}

	return dirPath
}
