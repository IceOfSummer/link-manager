package internal

import "os"

var DebugEnable = false

func init() {
	if os.Getenv("DEBUG_ENABLED") == "true" {
		DebugEnable = true
	} else {
		DebugEnable = false
	}
}
