package displayer

import (
	"fmt"
	"strings"

	"github.com/symbolic-link-manager/internal/configuration"
)

func DisplayLinks(links ...configuration.Link) {
	if len(links) == 0 {
		fmt.Println("没有找到对应的链接")
		return
	}
	var builder strings.Builder
	for _, v := range links {
		builder.WriteString(v.Name)
		builder.WriteString(":")
		builder.WriteString(v.Alias)
		builder.WriteString(" => ")
		builder.WriteString(v.Path)
		builder.WriteString("\n")
	}
	fmt.Print(builder.String())
}

func DisplayBindsWithStringRoot(root string, binds ...configuration.LinkBindItem) {
	if len(binds) == 0 {
		fmt.Println("没有找到对应的绑定")
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

func DisplayBinds(root *configuration.LinkBindItem, binds ...configuration.LinkBindItem) {
	DisplayBindsWithStringRoot(root.String(), binds...)
}

func DisplayUsingLink(links []configuration.UsingLink) {
	if len(links) == 0 {
		fmt.Println("没有正在使用的链接")
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
