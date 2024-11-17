package storage

// BindsData
// Key 使用 [Tag.Linkname]
type BindsData map[string][]*LinkBindItem

type configuration struct {
	DeclaredLinkNames []string
	Tags              []*Tag
	Binds             BindsData
}

// LinkBindItem 一个链接绑定.
//
// 首先由 [BindsData] 获取到 [Tag.Linkname]. 之后即可创建一个完整的链接:
//
// [BindsData].key : [LinkBindItem.CurrentTag] ==> [LinkBindItem.TargetName] : [LinkBindItem.TargetTag]
type LinkBindItem struct {
	CurrentTag string
	TargetName string
	TargetTag  string
}

type Tag struct {
	// 链接名称
	Linkname string
	// 标签名称
	TagName string
	// 标签路径. **绝对路径**
	Path string
}
