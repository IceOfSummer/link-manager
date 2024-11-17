// Helper functions
package cmd

import (
	"github.com/symbolic-link-manager/internal/core"
	"github.com/symbolic-link-manager/internal/storage"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SetUpTestEnvironment() {
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(path.Join(filename, ".."))

	err := os.Setenv(storage.AppHomeEnvKey, path.Join(root, "tmp"))
	if err != nil {
		panic(err)
	}
	_ = os.Mkdir(storage.AppHome(), 0b111_101_000)
	CleanUp()
}

func CleanUp() {
	target := path.Join(storage.AppHome(), "configuration.json")
	stat, err := os.Stat(target)
	if err != nil {
		return
	}
	if stat.Size() == 0 {
		return
	}
	_ = os.WriteFile(target, []byte(""), 0b111_101_000)
}

func ExecuteCommand(t *testing.T, args ...string) {
	rootCmd.SetArgs(args)
	assert.NoError(t, rootCmd.Execute())
}

func TestMain(m *testing.M) {
	SetUpTestEnvironment()
	code := m.Run()
	CleanUp()
	os.Exit(code)
}

func Exist[E any](slice []E, searchFn func(ele E) bool) bool {
	for _, v := range slice {
		if searchFn(v) {
			return true
		}
	}
	return false
}

func LinkNameExist(linkName string) bool {
	return Exist(core.ListDeclaredLinkNames(), func(name string) bool {
		return name == linkName
	})
}

// TagExist
// 判断标签是否存在, 如果 [path] 参数为空，则不检查路径
func TagExist(linkName, tag, path string) bool {
	return Exist(core.ListTags(linkName), func(link *storage.Tag) bool {
		return link.TagName == tag && (path == link.Path || path == "")
	})
}

func BindExist(linkName, tag, targetLinkName, targetTag string) bool {
	return Exist(core.ListBinds(linkName), func(bind *core.BindVO) bool {
		return bind.Tag == tag && bind.TargetLinkname == targetLinkName && bind.TargetTag == targetTag
	})
}

// 准备测试使用文件夹
func prepareTestDirectory(linkName string) (string, error) {
	home := storage.AppHome()

	testRoot := path.Join(home, "test")
	_ = os.Mkdir(testRoot, 0b111_111_101)

	target := path.Join(testRoot, linkName)
	stat, err := os.Stat(target)
	if err != nil {
		if !os.IsNotExist(err) {
			return "", err
		}
	} else if stat.IsDir() {
		return filepath.FromSlash(target), nil
	}
	err = os.Mkdir(path.Join(testRoot, linkName), 0b111_111_101)
	if err != nil {
		return "", err
	}
	return filepath.FromSlash(target), nil
}

// 创建绑定
func CreateBind(t *testing.T, baseName string, useRealDirectory bool) (*storage.Tag, *storage.Tag) {
	name, tag := baseName, baseName+"_tag"
	name1, tag1 := baseName+"1", baseName+"_tag1"
	var path0, path1 string

	if useRealDirectory {
		p, err := prepareTestDirectory(name)
		assert.NoError(t, err)
		path0 = p
		p, err = prepareTestDirectory(name)
		assert.NoError(t, err)
		path1 = p
	} else {
		path0 = "/fake/" + name
		path1 = "/fake/" + name1
	}
	ExecuteCommand(t, "add", "link", name)
	ExecuteCommand(t, "add", "tag", name, tag, path0)

	ExecuteCommand(t, "add", "link", name1)
	ExecuteCommand(t, "add", "tag", name1, tag1, path1)

	ExecuteCommand(t, "add", "bind", name+":"+tag, name1+":"+tag1)
	return &storage.Tag{Linkname: name, TagName: tag, Path: path0},
		&storage.Tag{Linkname: name1, TagName: tag1, Path: path1}
}

func AbsPath(p string) string {
	abs, err := filepath.Abs(p)
	if err != nil {
		panic(err)
	}
	return abs
}
