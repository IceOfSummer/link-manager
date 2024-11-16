package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkDelete(t *testing.T) {
	linkName := "TestLinkDelete"
	ExecuteCommand(t, "add", "link", linkName)

	assert.True(t, LinkNameExist(linkName))

	ExecuteCommand(t, "delete", "link", linkName)
	assert.False(t, LinkNameExist(linkName))
}

func TestDeleteTag(t *testing.T) {
	linkName := "TestDeleteTag"
	tag := "foo"

	ExecuteCommand(t, "add", "link", linkName)
	absPath := AbsPath("/val")
	ExecuteCommand(t, "add", "tag", linkName, tag, absPath)

	assert.True(t, TagExist(linkName, tag, absPath))

	ExecuteCommand(t, "delete", "tag", linkName, tag)

	assert.False(t, TagExist(linkName, tag, absPath))
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
