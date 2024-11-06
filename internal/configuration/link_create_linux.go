package configuration

import "os"

func createLink(src, des string) error {
	return os.Symlink(src, des)
}
