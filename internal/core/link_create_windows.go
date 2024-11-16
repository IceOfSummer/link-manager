package core

import (
	"errors"
	"os/exec"

	"github.com/axgle/mahonia"
	"github.com/symbolic-link-manager/internal/logger"
)

// 在 src 处创建一个指向 des 的软连接
func createLink(src, des string) error {
	// Create link and bypass admin permission requirement on windows.
	cmd := exec.Command("cmd", "/c", "mklink", "/J", src, des)

	logger.LogDebug("Running Command: " + cmd.String())
	out, err := cmd.CombinedOutput()
	if err != nil {
		decoder := mahonia.NewDecoder("gbk")
		return errors.New(string(decoder.ConvertString(string(out))))
	}
	return nil
}
