package core

import "github.com/symbolic-link-manager/internal/storage"

func AddLinkDeclaration(linkname string) error {
	err := EnsureLinkStatus(linkname, false)
	if err != nil {
		return err
	}
	storage.InsertLinkDeclaration(linkname)
	return nil
}

func AddTag(linkname string, tag string, path string) error {
	err := EnsureLinkStatus(linkname, true)
	if err != nil {
		return err
	}
	err = EnsureTagStatus(linkname, tag, false)
	if err != nil {
		return err
	}
	return storage.InsertTag(linkname, tag, path)
}

func AddBind(srcName, srcTag, targetName, targetTag string) error {
	err := EnsureTagStatus(srcName, srcTag, true)
	if err != nil {
		return err
	}
	err = EnsureTagStatus(targetName, targetTag, true)
	if err != nil {
		return err
	}
	storage.InsertBind(srcName, srcTag, targetName, targetTag)
	return nil
}

func ListDeclaredLinkNames() []string {
	return storage.ListDeclaredLinkNames()
}

// ListTags 列出所有标签
// 如果 [Linkname] 为空字符串，则返回所有标签
func ListTags(linkname string) []*storage.Tag {
	tags := storage.ListTags()
	if linkname == "" {
		return tags
	}

	result := make([]*storage.Tag, 0)

	for _, tag := range tags {
		if tag.Linkname == linkname {
			result = append(result, tag)
		}
	}
	return result
}

type BindVO struct {
	Linkname       string
	Tag            string
	TargetLinkname string
	TargetTag      string
}

func toBindVO(linkname string, binds []*storage.LinkBindItem) []*BindVO {
	result := make([]*BindVO, len(binds))

	for i, bind := range binds {
		result[i] = &BindVO{
			Linkname:       linkname,
			Tag:            bind.CurrentTag,
			TargetLinkname: bind.TargetName,
			TargetTag:      bind.TargetTag,
		}
	}
	return result
}

// ListBinds 列出所有绑定
// 如果 [Linkname] 为空字符串，则返回所有绑定
func ListBinds(linkname string) []*BindVO {
	bindMap := storage.ListBinds()
	if linkname == "" {
		result := make([]*BindVO, 0)
		for n, items := range bindMap {
			result = append(result, toBindVO(n, items)...)
		}
		return result
	}

	binds, ok := bindMap[linkname]
	if !ok {
		return make([]*BindVO, 0)
	}
	return toBindVO(linkname, binds)
}

func ListAllBinds() storage.BindsData {
	return storage.ListBinds()
}

func RemoveTag(linkname, tag string) (*storage.Tag, error) {
	err := EnsureTagStatus(linkname, tag, true)
	if err != nil {
		return nil, err
	}
	return storage.DeleteTag(linkname, tag), nil
}

func RemoveLink(linkname string) ([]*storage.Tag, error) {
	err := EnsureLinkStatus(linkname, true)
	if err != nil {
		return []*storage.Tag{}, err
	}

	return storage.DeleteLink(linkname), nil
}

func RemoveBind(linkname string, linkBindItem *storage.LinkBindItem) error {
	err := EnsureBindStatus(linkname, linkBindItem.CurrentTag, linkBindItem.TargetName, linkBindItem.TargetTag, true)
	if err != nil {
		return err
	}
	storage.DeleteBind(linkname, linkBindItem)
	return nil
}

func RenameLink(oldName, newName string) error {
	err := EnsureLinkStatus(oldName, true)
	if err != nil {
		return err
	}
	err = EnsureLinkStatus(newName, false)
	if err != nil {
		return err
	}
	storage.ReplaceLinkname(oldName, newName)
	return nil
}

func UpdateTag(entity *storage.Tag) error {
	err := EnsureTagStatus(entity.Linkname, entity.TagName, true)
	if err != nil {
		return err
	}
	storage.UpdateTag(entity)
	return nil
}

type UpdateBindDTO struct {
	// Required
	SrcName string
	// Required
	SrcTag string
	// Required
	TargetName string
	// Required
	TargetTag string
	// Optional
	NewTag string
}

func UpdateBind(entity *UpdateBindDTO) error {
	err := EnsureBindStatus(entity.SrcName, entity.SrcTag, entity.TargetName, entity.TargetTag, true)
	if err != nil {
		return err
	}
	err = EnsureTagStatus(entity.TargetName, entity.NewTag, true)
	if err != nil {
		return err
	}
	storage.UpdateBind(entity.SrcName, &storage.LinkBindItem{
		CurrentTag: entity.SrcTag,
		TargetName: entity.TargetName,
		TargetTag:  entity.TargetTag,
	}, entity.NewTag)
	return nil
}
