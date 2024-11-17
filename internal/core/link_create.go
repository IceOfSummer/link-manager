//go:build linux || darwin

package core

import (
	"os"

	"github.com/symbolic-link-manager/internal/logger"
)

// 在 src 处创建一个指向 des 的软连接
func createLink(src, des string) error {
	logger.LogDebug("Creating link, src = " + src + "; des: " + des)
	return os.Symlink(des, src)
}
