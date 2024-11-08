package configuration

import (
	"os"
	"path"

	"github.com/symbolic-link-manager/internal/localizer"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var APP_HOME string
var configFilePath string

func init() {
	home, ok := os.LookupEnv("SLINK_MANAGER_HOME")
	if !ok {
		panic(localizer.GetMessageAndIgnoreError(&i18n.LocalizeConfig{MessageID: "error.noenv"}))
	}
	APP_HOME = home
	configFilePath = path.Join(home, "configuration.json")
}

type Link struct {
	// 链接名称
	Name string
	// 链接别名
	Alias string
	// 链接路径
	Path string
}

func (t Link) String() string {
	return t.Name + ":" + t.Alias
}

// LinkBindItem 一个链接绑定.
//
// 首先由 [BindsData] 获取到 [Link.Name]. 之后即可创建一个完整的链接:
//
// [BindsData].key : [LinkBindItem.CurrentAlias] ==> [LinkBindItem.TargetName] : [LinkBindItem.TargetAlias]
type LinkBindItem struct {
	CurrentAlias string
	TargetName   string
	TargetAlias  string
}

func (t LinkBindItem) String() string {
	return t.TargetName + ":" + t.TargetAlias
}

// BindsData
// Key 使用 [Link.Name]
type BindsData map[string][]LinkBindItem

type configuration struct {
	DeclaredLinkNames []string
	Links             []Link
	Binds             BindsData
}
