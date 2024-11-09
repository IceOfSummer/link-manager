package displayer

import (
	"fmt"
	"github.com/symbolic-link-manager/internal/localizer"
	"strings"

	"github.com/symbolic-link-manager/internal/configuration"
)

func DisplayLinks(links ...configuration.Link) {
	if len(links) == 0 {
		fmt.Println(localizer.GetMessageWithoutParam(localizer.NothingFound))
		return
	}
	var builder strings.Builder
	for _, v := range links {
		builder.WriteString(v.Name)
		builder.WriteString(":")
		builder.WriteString(v.Tag)
		builder.WriteString(" => ")
		builder.WriteString(v.Path)
		builder.WriteString("\n")
	}
	fmt.Print(builder.String())
}

func DisplayBindsWithStringRoot(root string, binds ...configuration.LinkBindItem) {
	if len(binds) == 0 {
		fmt.Println(localizer.GetMessageWithoutParam(localizer.NothingFound))
		return
	}
	var builder strings.Builder

	for _, v := range binds {
		builder.WriteString(root)
		builder.WriteString(" => ")
		builder.WriteString(v.String())
		builder.WriteString("\n")
	}
	fmt.Print(builder.String())
}

func DisplayUsingLink(links []configuration.UsingLink) {
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
