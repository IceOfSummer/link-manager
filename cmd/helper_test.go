// Helper functions
package cmd

import (
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/symbolic-link-manager/internal/configuration"
)

func SetUpTestEnvironment() {
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(path.Join(filename, ".."))

	err := os.Setenv(configuration.AppHomeEnvKey, path.Join(root, "tmp"))
	if err != nil {
		panic(err)
	}
	_ = os.Mkdir(configuration.AppHome(), 0b111_101_000)
}

func CleanUp() {
	_ = os.WriteFile(path.Join(configuration.AppHome(), "configuration.json"), []byte(""), 0b111_101_000)
}

func ExecuteCommand(t *testing.T, args ...string) {
	rootCmd.SetArgs(args)
	assert.Nil(t, rootCmd.Execute())
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
	return Exist(configuration.ListLinkNames(), func(name string) bool {
		return name == linkName
	})
}

// TagExsit
// 判断标签是否存在, 如果 [path] 参数为空，则不检查路径
func TagExsit(linkName, tag, path string) bool {
	return Exist(configuration.ListLinkTags(linkName), func(link configuration.Link) bool {
		return link.Tag == tag && (path == link.Path || path == "")
	})
}

func BindExist(linkName, tag, targetLinkName, targetTag string) bool {
	return Exist(configuration.ListBinds(linkName, tag), func(bind configuration.LinkBindItem) bool {
		return bind.CurrentTag == tag && bind.TargetName == targetLinkName && bind.TargetTag == targetTag
	})
}

// 创建绑定
func CreateBind(t *testing.T, baseName string) (*configuration.Link, *configuration.Link) {
	name, tag, path := baseName, baseName+"_tag", baseName+"/path"
	name1, tag1, path1 := baseName+"1", baseName+"_tag1", baseName+"/path1"
	ExecuteCommand(t, "add", "link", name)
	ExecuteCommand(t, "add", "tag", name, tag, path)

	ExecuteCommand(t, "add", "link", name1)
	ExecuteCommand(t, "add", "tag", name1, tag1, path1)

	ExecuteCommand(t, "add", "bind", name+":"+tag, name1+":"+tag1)
	return &configuration.Link{Name: name, Tag: tag, Path: path},
		&configuration.Link{Name: name1, Tag: tag1, Path: path1}
}
