// Top level module.
package logger

import "fmt"

var DebugEnable bool = false

func LogDebug(msg string) {
	if DebugEnable {
		fmt.Println("[DEBUG] " + msg)
	}
}
