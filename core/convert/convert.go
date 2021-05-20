package convert

import (
	"fmt"
	"strconv"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/16
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
func (slf *convertStruct) BytesLengthToHumanReadable(length int64, decimal int) string {
	const unit = 1000
	if length < unit {
		return fmt.Sprintf("%d B", length)
	}

	div, exp := int64(unit), 0
	for n := length / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	format := "%f"
	if decimal > 0 {
		format = "%." + strconv.Itoa(decimal) + "f"
	}

	return fmt.Sprintf(format + " %cB", float64(length)/float64(div), "kMGTPE"[exp])
}
