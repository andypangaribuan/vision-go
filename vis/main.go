package vis

import (
	"os"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/03
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
//goland:noinspection GoUnusedGlobalVariable
var (
	Util  iUtil
	Const *constVariable
)


func init() {
	Const = &constVariable{
		PathSeparator: string(os.PathSeparator),
	}
}
