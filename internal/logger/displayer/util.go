package displayer

import (
	"fmt"
	"github.com/symbolic-link-manager/internal/core"
	"github.com/symbolic-link-manager/internal/localizer"
	"github.com/symbolic-link-manager/internal/storage"
	"strings"
)

func DisplayLinks(links ...*storage.Tag) {
	if len(links) == 0 {
		fmt.Println(localizer.GetMessageWithoutParam(localizer.NothingFound))
		return
	}
	var builder strings.Builder
	for _, v := range links {
		builder.WriteString(v.Linkname)
		builder.WriteString(":")
		builder.WriteString(v.TagName)
		builder.WriteString(" => ")
		builder.WriteString(v.Path)
		builder.WriteString("\n")
	}
	fmt.Print(builder.String())
}

func DisplayBindsVO(binds []*core.BindVO) {
	if len(binds) == 0 {
		fmt.Println(localizer.GetMessageWithoutParam(localizer.NothingFound))
		return
	}
	var builder strings.Builder

	for _, v := range binds {
		builder.WriteString(v.Linkname)
		builder.WriteString(":")
		builder.WriteString(v.Tag)
		builder.WriteString(" => ")
		builder.WriteString(v.TargetLinkname)
		builder.WriteString(":")
		builder.WriteString(v.TargetTag)
		builder.WriteString("\n")
	}
	fmt.Print(builder.String())
}

func DisplayBindsWithStringRoot(root string, binds ...*storage.LinkBindItem) {
	if len(binds) == 0 {
		fmt.Println(localizer.GetMessageWithoutParam(localizer.NothingFound))
		return
	}
	var builder strings.Builder

	for _, v := range binds {
		builder.WriteString(root)
		builder.WriteString(":")
		builder.WriteString(v.CurrentTag)
		builder.WriteString(" => ")
		builder.WriteString(v.TargetName)
		builder.WriteString(":")
		builder.WriteString(v.TargetTag)
		builder.WriteString("\n")
	}
	fmt.Print(builder.String())
}

func DisplayUsingLink(links []core.UsingLink) {
	if len(links) == 0 {
		fmt.Println(localizer.GetMessageWithoutParam(localizer.NothingFound))
		return
	}
	var builder strings.Builder

	for _, v := range links {
		builder.WriteString(v.Name)
		builder.WriteString(" => ")
		builder.WriteString(v.Path)
		builder.WriteString("\n")
	}
	fmt.Print(builder.String())
}
