//go:build linux || darwin

package core

import "os"

// 在 src 处创建一个指向 des 的软连接
func createLink(src, des string) error {
	return os.Symlink(des, src)
}
