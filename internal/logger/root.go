// Top level module.
package logger

import (
	"fmt"
	"github.com/symbolic-link-manager/internal"
)

func LogDebug(msg string) {
	if internal.DebugEnable {
		fmt.Println("[DEBUG] " + msg)
	}
}

func LogError(err error) {
	fmt.Println(err.Error())
}
