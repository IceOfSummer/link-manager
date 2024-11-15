package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/symbolic-link-manager/internal/configuration"
)

func TestLinkDelete(t *testing.T) {
	linkName := "TestLinkDelete"
	ExecuteCommand(t, "add", "link", linkName)
	assert.Equal(t,
		true,
		Exist(configuration.ListLinkNames(), func(ele string) bool {
			return ele == linkName
		}),
	)

	ExecuteCommand(t, "delete", "link", linkName)
	assert.Equal(t,
		false,
		Exist(configuration.ListLinkNames(), func(ele string) bool {
			return ele == linkName
		}),
	)
}

func TestDeleteTag(t *testing.T) {
	linkName := "TestDeleteTag"
	tag := "foo"
	ExecuteCommand(t, "add", "link", linkName)
	ExecuteCommand(t, "add", "tag", linkName, tag, "/val")

	configuration.ListLinkTags(linkName)
	assert.Equal(t,
		true,
		Exist(configuration.ListLinkTags(linkName), func(ele configuration.Link) bool {
			return ele.Tag == tag
		}),
	)

	ExecuteCommand(t, "delete", "tag", linkName, tag)

	assert.Equal(t,
		false,
		Exist(configuration.ListLinkTags(linkName), func(ele configuration.Link) bool {
			return ele.Tag == tag
		}),
	)

}

func TestDeleteLinkBind(t *testing.T) {
	name, tag := "TestDeleteLinkBind", "foo"
	name1, tag1 := "TestDeleteLinkBind1", "foo1"
	ExecuteCommand(t, "add", "link", name)
	ExecuteCommand(t, "add", "tag", name, tag, "/val")

	ExecuteCommand(t, "add", "link", name1)
	ExecuteCommand(t, "add", "tag", name1, tag1, "/val")

	ExecuteCommand(t, "add", "bind", name+":"+tag, name1+":"+tag1)

	assert.Equal(t,
		true,
		Exist(configuration.ListBinds(name, tag), func(ele configuration.LinkBindItem) bool {
			return ele.TargetTag == tag1 && ele.TargetName == name1 && ele.CurrentTag == tag
		}),
	)

	ExecuteCommand(t, "delete", "bind", name+":"+tag, name1+":"+tag1)

	assert.Equal(t,
		false,
		Exist(configuration.ListBinds(name, tag), func(ele configuration.LinkBindItem) bool {
			return ele.TargetTag == tag1 && ele.TargetName == name1 && ele.CurrentTag == tag
		}),
	)
}
