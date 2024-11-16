package core

import (
	"github.com/symbolic-link-manager/internal/localizer"
	"github.com/symbolic-link-manager/internal/storage"
)

func EnsureLinkStatus(linkname string, expectExist bool) error {
	names := storage.ListDeclaredLinkNames()
	find := false
	for _, name := range names {
		if name == linkname {
			find = true
			break
		}
	}
	if find && expectExist {
		return nil
	} else if !find && !expectExist {
		return nil
	}
	if expectExist {
		return localizer.CreateNoSuchLinkError(linkname)
	}
	return localizer.CreateLinkNameAlreadyExistError(linkname)
}

func EnsureTagStatus(linkname string, tagName string, expectExist bool) error {
	tag := storage.FindTag(linkname, tagName)

	if tag == nil && !expectExist {
		return nil
	} else if tag != nil && expectExist {
		return nil
	}
	if expectExist {
		return localizer.CreateNoSuchTagError(linkname, tagName)
	}
	return localizer.CreateTagAlreadyExistError(linkname, tagName)
}

func EnsureBindStatus(srcName, srcTag, targetName, targetTag string, expectExist bool) error {
	err := EnsureTagStatus(srcName, srcTag, expectExist)
	if err != nil {
		return err
	}
	err = EnsureTagStatus(targetName, targetTag, expectExist)
	if err != nil {
		return err
	}

	bind := storage.FindBind(srcName, srcTag, targetName, targetTag)

	if bind == nil && !expectExist {
		return nil
	} else if bind != nil && expectExist {
		return nil
	}
	if expectExist {
		return localizer.CreateError(localizer.ErrorBindNotExist)
	}
	return localizer.CreateError(localizer.ErrorBindAlreadyExist)
}
