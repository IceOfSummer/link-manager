package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateLinkName(t *testing.T) {
	oldName, tag, path := "TestUpdateLinkName", "tag", AbsPath("/foo/bar")
	targetName, tag1, path1 := "TestUpdateLinkName1", "tag1", AbsPath("/foo/bar1")
	ExecuteCommand(t, "add", "link", oldName)
	ExecuteCommand(t, "add", "tag", oldName, tag, path)

	ExecuteCommand(t, "add", "link", targetName)
	ExecuteCommand(t, "add", "tag", targetName, tag1, path1)

	ExecuteCommand(t, "add", "bind", oldName+":"+tag, targetName+":"+tag1)

	newName := oldName + "_new"
	ExecuteCommand(t, "update", "link", oldName, "--name="+newName)

	assert.True(t, LinkNameExist(newName))
	assert.False(t, LinkNameExist(oldName))
	assert.True(t, TagExist(newName, tag, path))
	assert.True(t, BindExist(newName, tag, targetName, tag1))
}

func TestUpdateTag(t *testing.T) {
	linkName, tag, path := "TestUpdateLinkName", "tag", AbsPath("/foo/bar")

	ExecuteCommand(t, "add", "link", linkName)
	ExecuteCommand(t, "add", "tag", linkName, tag, path)

	newPath := AbsPath(path + "/new")
	ExecuteCommand(t, "update", "tag", linkName, tag, "--path="+newPath)

	assert.True(t, TagExist(linkName, tag, newPath))
}

func TestUpdateBind(t *testing.T) {
	cur, target := CreateBind(t, "TestUpdateBind", false)

	newTag, path := "TestUpdateBind_new_tag", "/foo/bar"
	ExecuteCommand(t, "add", "tag", target.Linkname, newTag, path)

	ExecuteCommand(t, "update", "bind", cur.Linkname+":"+cur.TagName, target.Linkname+":"+target.TagName, "--targetTag="+newTag)

	assert.True(t, BindExist(cur.Linkname, cur.TagName, target.Linkname, newTag))
	assert.False(t, BindExist(cur.Linkname, cur.TagName, target.Linkname, target.TagName))
}
