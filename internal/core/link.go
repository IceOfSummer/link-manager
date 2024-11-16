package core

import (
	"github.com/symbolic-link-manager/internal/logger"
	"github.com/symbolic-link-manager/internal/storage"
	"os"
	"path"
	"path/filepath"
)

const appDirectory = "app"

// UseLink 使用当前链接.
// 返回所有设置的链接，包括间接连接的。
func UseLink(linkname, tag string) []*storage.Tag {
	holder := make([]*storage.Tag, 0)
	useLink0(storage.FindTag(linkname, tag), &holder)
	return holder
}

func useLink0(tag *storage.Tag, resultHolder *[]*storage.Tag) {
	if tag == nil {
		return
	}
	*resultHolder = append(*resultHolder, tag)

	holder := filepath.FromSlash(path.Join(storage.AppHome(), appDirectory))
	_, err := os.Stat(holder)
	if os.IsNotExist(err) {
		logger.LogDebug("Creating 'app' directory.")
		err := os.Mkdir(holder, 0b111_101_101)
		if err != nil {
			panic(err)
		}
	}

	appPath := filepath.FromSlash(path.Join(holder, tag.Linkname))

	lk, _ := os.Readlink(appPath)
	if lk != "" {
		logger.LogDebug("Removing old tag file.")
		err := os.Remove(appPath)
		if err != nil {
			panic(err)
		}
	}

	err = createLink(appPath, tag.Path)
	if err != nil {
		panic(err)
	}

	// use all binds
	binds := ListBinds(tag.Linkname)

	for _, bind := range binds {
		if bind.Tag == tag.TagName {
			useLink0(storage.FindTag(bind.TargetLinkname, bind.TargetLinkname), resultHolder)
		}
	}
}

type UsingLink struct {
	Name string
	Path string
}

func ListUsing() ([]UsingLink, error) {
	base := path.Join(storage.AppHome(), "app")
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
