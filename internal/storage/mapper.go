// Package storage 操作持久化数据. 所有数据校验应该在调用该文件的方法前进行，若碰到异常参数，则会直接 panic
package storage

import "path/filepath"

// FindTag 搜素对应的标签
// 如果找到，返回对应的值和索引。否则返回空和 [-1]
func (config *configuration) FindTag(linkname, tagName string) (*Tag, int) {
	tags := config.Tags
	for pos, tag := range tags {
		if tag.Linkname == linkname && tag.TagName == tagName {
			return tag, pos
		}
	}
	return nil, -1
}

// FindBind 搜素对应的所有绑定
// 如果找到，返回对应的值和索引。否则返回空和 [-1]
func (config *configuration) FindBind(linkname, tagName, targetLinkname, targetTagName string) (*LinkBindItem, int) {
	arr, ok := config.Binds[linkname]
	if !ok {
		return nil, -1
	}
	for i, item := range arr {
		if item.CurrentTag == tagName && item.TargetTag == targetTagName && item.TargetName == targetLinkname {
			return item, i
		}
	}
	return nil, -1
}

func (config *configuration) ListBinds(linkname, tagName string) []*LinkBindItem {
	arr, ok := config.Binds[linkname]
	if !ok {
		return make([]*LinkBindItem, 0)
	}
	result := make([]*LinkBindItem, 0)
	for _, item := range arr {
		if item.CurrentTag == tagName {
			result = append(result, item)
		}
	}
	return result
}

func ListTags() []*Tag {
	return readConfig().Tags
}

func FindTag(linkname, tagName string) *Tag {
	config := readConfig()
	tag, _ := config.FindTag(linkname, tagName)
	return tag
}

func ListBinds() BindsData {
	return readConfig().Binds
}

func FindBind(srcName, srcTag, targetName, targetTag string) *LinkBindItem {
	bindMap := ListBinds()
	binds, ok := bindMap[srcName]
	if !ok {
		return nil
	}
	for _, bind := range binds {
		if bind.CurrentTag == srcTag && bind.TargetTag == targetTag && bind.TargetName == targetName {
			return bind
		}
	}
	return nil
}

func ListDeclaredLinkNames() []string {
	return readConfig().DeclaredLinkNames
}

func InsertLinkDeclaration(linkname string) {
	config := readConfig()
	config.DeclaredLinkNames = append(config.DeclaredLinkNames, linkname)
	saveConfig(&config)
}

func InsertTag(linkname, tagName, tagPath string) error {
	config := readConfig()
	absPath, err := filepath.Abs(tagPath)
	if err != nil {
		return err
	}
	config.Tags = append(config.Tags, &Tag{Linkname: linkname, TagName: tagName, Path: absPath})
	saveConfig(&config)
	return nil
}

func InsertBind(srcName, srcTag, targetName, targetTag string) {
	config := readConfig()
	binds, ok := config.Binds[srcName]
	item := &LinkBindItem{CurrentTag: srcTag, TargetTag: targetTag, TargetName: targetName}
	if !ok {
		config.Binds[srcName] = []*LinkBindItem{item}
	} else {
		config.Binds[srcName] = append(binds, item)
	}
	saveConfig(&config)
}

// rebuildDeclaredLinks
// 根据 [Link] 重新创建已经声明的链接
func rebuildDeclaredLinks(tags []*Tag) []string {
	var names = make([]string, 0)
	set := make(map[string]struct{})
	for _, tag := range tags {
		_, ok := set[tag.Linkname]
		if !ok {
			set[tag.Linkname] = struct{}{}
			names = append(names, tag.Linkname)
		}
	}
	return names
}

func DeleteTag(linkname, tag string) *Tag {
	config := readConfig()
	deleted, pos := config.FindTag(linkname, tag)
	config.Tags = append(config.Tags[:pos], config.Tags[pos+1:]...)
	config.DeclaredLinkNames = rebuildDeclaredLinks(config.Tags)
	saveConfig(&config)
	return deleted
}

func DeleteBind(linkname string, linkBindItem *LinkBindItem) {
	config := readConfig()

	_, pos := config.FindBind(linkname, linkBindItem.CurrentTag, linkBindItem.TargetName, linkBindItem.TargetTag)
	if pos == -1 {
		panic("Bind not exist.")
	}
	config.Binds[linkname] = append(config.Binds[linkname][:pos], config.Binds[linkname][pos+1:]...)
	saveConfig(&config)
}

func ReplaceLinkname(oldName, newName string) {
	config := readConfig()

	newTags := make([]*Tag, 0)
	for _, tag := range config.Tags {
		if tag.Linkname == oldName {
			tag.Linkname = newName
		}
		newTags = append(newTags, tag)
	}
	config.Tags = newTags
	config.DeclaredLinkNames = rebuildDeclaredLinks(config.Tags)
	v, ok := config.Binds[oldName]
	if ok {
		config.Binds[newName] = v
		delete(config.Binds, oldName)
	}
	saveConfig(&config)
}

// UpdateTag 更新标签
// [newTag] 可以使用空字符串表示不更新.
func UpdateTag(entity *Tag) {
	config := readConfig()

	link, _ := config.FindTag(entity.Linkname, entity.TagName)
	if link == nil {
		panic("Tag not found.")
	}
	changed := 0
	if entity.Path != "" {
		changed++
		link.Path = entity.Path
	}
	if changed > 0 {
		saveConfig(&config)
	}
}

func UpdateBind(linkname string, item *LinkBindItem, newTag string) {
	config := readConfig()
	bind, _ := config.FindBind(linkname, item.CurrentTag, item.TargetName, item.TargetTag)
	if bind == nil {
		panic("Bind not exist.")
	}
	if updateBind0(bind, newTag) {
		saveConfig(&config)
	}
}

// 更新绑定
// 如果有字段更新，返回 true
func updateBind0(item *LinkBindItem, newTag string) bool {
	changed := 0
	if newTag != "" {
		changed++
		item.TargetTag = newTag
	}

	return changed > 0
}

// DeleteLink 删除链接
// 返回被删除的 Tag
func DeleteLink(linkname string) []*Tag {
	config := readConfig()
	newTags := make([]*Tag, 0)
	deleted := make([]*Tag, 0)
	for _, tag := range config.Tags {
		if tag.Linkname == linkname {
			deleted = append(deleted, tag)
		} else {
			newTags = append(newTags, tag)
		}
	}
	config.Tags = newTags
	config.DeclaredLinkNames = rebuildDeclaredLinks(config.Tags)
	saveConfig(&config)
	return deleted
}
