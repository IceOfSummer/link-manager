package configuration

import (
	"github.com/symbolic-link-manager/internal/localizer"
	"os"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var lazyLoadAppHome string
var AppHomeEnvKey = "SLINK_MANAGER_HOME"

// AppHome 获取 lazyLoadAppHome
func AppHome() string {
	if lazyLoadAppHome != "" {
		return lazyLoadAppHome
	}
	home, ok := os.LookupEnv(AppHomeEnvKey)
	if !ok {
		panic(localizer.GetMessage(&i18n.LocalizeConfig{MessageID: "error.noenv"}))
	}
	lazyLoadAppHome = home
	return home
}

func init() {
}

type Link struct {
	// 链接名称
	Name string
	// 链接标签
	Tag string
	// 链接路径
	Path string
}

func (t Link) String() string {
	return t.Name + ":" + t.Tag
}

// LinkBindItem 一个链接绑定.
//
// 首先由 [BindsData] 获取到 [Link.Name]. 之后即可创建一个完整的链接:
//
// [BindsData].key : [LinkBindItem.CurrentTag] ==> [LinkBindItem.TargetName] : [LinkBindItem.TargetAlias]
type LinkBindItem struct {
	CurrentTag  string
	TargetName  string
	TargetAlias string
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
