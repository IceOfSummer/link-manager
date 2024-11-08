// 设置环境变量的实现.
package configuration

import (
	"os"
	"path"
	"path/filepath"

	"github.com/symbolic-link-manager/internal/logger"
)

const appDirectory = "app"

// 使用当前链接.
// 返回所有设置的链接，包括间接连接的。
func UseLink(link *Link) []Link {
	holder := filepath.FromSlash(path.Join(DENV_HOME, appDirectory))
	_, err := os.Stat(holder)
	if os.IsNotExist(err) {
		logger.LogDebug("Creating 'app' directory.")
		err := os.Mkdir(holder, 0b111_101_101)
		if err != nil {
			panic(err)
		}
	}

	target := filepath.FromSlash(path.Join(holder, link.Name))

	lk, _ := os.Readlink(target)
	if lk != "" {
		logger.LogDebug("Removing old link file.")
		err := os.Remove(target)
		if err != nil {
			panic(err)
		}
	}

	err = createLink(link.Path, target)
	if err != nil {
		panic(err)
	}

	// use all binds
	binds := ListBinds(&LinkBindItem{Name: link.Name, Alias: link.Alias})

	result := make([]Link, 0)
	result = append(result, *link)
	for _, v := range binds {
		result = append(result, UseLink(FindEnvByNameAndAlias(v.Name, v.Alias))...)
	}
	return result
}

type UsingLink struct {
	Name string
	Path string
}

func ListUsing() ([]UsingLink, error) {
	base := path.Join(DENV_HOME, "app")
	entries, err := os.ReadDir(base)
	logger.LogDebug("Searching " + base)
	if err != nil {
		return nil, err
	}
	result := make([]UsingLink, 0)
	for _, file := range entries {
		target := filepath.FromSlash(path.Join(base, file.Name()))
		lk, err := os.Readlink(target)
		if err == nil {
			result = append(result, UsingLink{
				Name: file.Name(),
				Path: lk,
			})
		} else {
			logger.LogDebug("Failed to read link file: " + err.Error())
		}
	}
	return result, nil
}
