package internal

import "os"

var DebugEnable = false

func init() {
	if os.Getenv("DEBUG") != "" {
		DebugEnable = true
	} else {
		DebugEnable = false
	}
}
