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
	assert.True(t, TagExist(linkName, tag, "/val"))

	ExecuteCommand(t, "delete", "tag", linkName, tag)

	assert.False(t, TagExist(linkName, tag, "/val"))
}

func TestDeleteLinkBind(t *testing.T) {
	name, tag := "TestDeleteLinkBind", "foo"
	name1, tag1 := "TestDeleteLinkBind1", "foo1"
	ExecuteCommand(t, "add", "link", name)
	ExecuteCommand(t, "add", "tag", name, tag, "/val")

	ExecuteCommand(t, "add", "link", name1)
	ExecuteCommand(t, "add", "tag", name1, tag1, "/val")

	ExecuteCommand(t, "add", "bind", name+":"+tag, name1+":"+tag1)
	assert.True(t, BindExist(name, tag, name1, tag1))

	ExecuteCommand(t, "delete", "bind", name+":"+tag, name1+":"+tag1)
	assert.False(t, BindExist(name, tag, name1, tag1))
}
