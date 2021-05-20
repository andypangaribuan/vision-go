package log

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/18
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
func codeStack(err error) string {
	if err == nil {
		return ""
	}

	exceptions := []string{
		"/github.com/labstack/echo/",
		"/libexec/src/net/http/",
	}
	isOnExceptions := func(filePath string) (isOnEx bool) {
		for _, ex := range exceptions {
			if strings.Contains(filePath, ex) {
				isOnEx = true
				break
			}
		}
		return
	}

	msg := fmt.Sprintf("%+v", err)
	arr := strings.Split(msg, "\n")
	files := make([][]string, 0)
	for _, v := range arr {
		if len(v) > 8 && v[:8] == "runtime." {
			break
		}

		if len(v) > 1 && v[:1] == "\t" {
			v = v[1:]
			vs := strings.Split(v, ":")
			if len(vs) == 2 {
				if _, err := strconv.Atoi(vs[1]); err == nil {
					if !isOnExceptions(vs[0]) {
						if _, err := os.Stat(vs[0]); !os.IsNotExist(err) {
							files = append(files, vs)
						}
					}
				}
			}
		}
	}

	if len(files) > 0 {
		codeStack := ""

		for _, vs := range files {
			lineAt, _ := strconv.Atoi(vs[1])
			if codes := readFile(vs[0], lineAt); codes != "" {
				if codeStack == "" {
					codeStack = ":: START CODE STACK\n\n"
				} else {
					codeStack += "\n\n\n"
				}

				codeStack += vs[0] + ":" + vs[1] + "\n"
				codeStack += codes
			}
		}

		if codeStack != "" {
			codeStack += "\n\n:: END CODE STACK"
			return codeStack
		}
	}

	return ""
}


func readFile(filePath string, lineAt int) string {
	includeTopLine := 4
	includeBottomLine := 4

	startAt := lineAt - includeTopLine
	stopAt := lineAt + includeBottomLine

	file, err := os.Open(filePath)
	if err != nil {
		return ""
	}

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	canStart := false
	lines := make([]string, 0)

	for scanner.Scan() {
		lineNumber++

		if canStart {
			if lineNumber > stopAt {
				break
			}

			line := scanner.Text()
			lines = append(lines, strconv.Itoa(lineNumber) + ": " + line)
		} else {
			if lineNumber == startAt {
				canStart = true
				line := scanner.Text()
				lines = append(lines, strconv.Itoa(lineNumber) + ": " + line)
			}
		}
	}

	n := len(lines)
	codes := ""
	if n > 0 {
		for i, v := range lines {
			codes += v
			if i + 1 != n {
				codes += "\n"
			}
		}
	}

	return codes
}
