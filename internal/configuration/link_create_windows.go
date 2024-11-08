package configuration

import (
	"errors"
	"os/exec"

	"github.com/axgle/mahonia"
	"github.com/symbolic-link-manager/internal/logger"
)

// Create link and bypass admin permission requirement on windows.
func createLink(src, des string) error {
	cmd := exec.Command("cmd", "/c", "mklink", "/J", des, src)

	logger.LogDebug("Running Command: " + cmd.String())
	out, err := cmd.CombinedOutput()
	if err != nil {
		decoder := mahonia.NewDecoder("gbk")
		return errors.New(string(decoder.ConvertString(string(out))))
	}
	return nil
}
