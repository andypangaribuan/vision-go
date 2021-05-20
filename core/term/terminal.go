package term

import (
	"fmt"
	"math"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/16
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
func (slf *termStruct) Move(line, cursor int) {
	fmt.Printf("\033[%d;%dH", line, cursor)
}

func (slf *termStruct) MoveUp(inc int) {
	fmt.Printf("\033[%dA", inc)
}

func (slf *termStruct) MoveDown(inc int) {
	fmt.Printf("\033[%dB", inc)
}

func (slf *termStruct) MoveLeft(length int) {
	if length == -1 {
		length = math.MaxInt32
	}
	fmt.Printf("\033[%dD", length)
}

func (slf *termStruct) MoveRight(length int) {
	if length == -1 {
		length = math.MaxInt32
	}
	fmt.Printf("\033[%dC", length)
}

func (slf *termStruct) ClearScreen() {
	fmt.Printf("\r\033[2J")
	slf.Move(1, 0)
}

func (slf *termStruct) EraseCurrentLine() {
	fmt.Printf("\r\033[2K")
}

func (slf *termStruct) ReplaceCurrentLine(format string, args ...interface{}) {
	slf.EraseCurrentLine()
	fmt.Printf(format, args...)
}
