package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/symbolic-link-manager/internal/configuration"
)

func TestLinkDeclarationAdd(t *testing.T) {
	linkName := "TestLinkDeclarationAdd"
	rootCmd.SetArgs([]string{"add", "link", linkName})
	err := rootCmd.Execute()
	if err != nil {
		t.Error(err)
	}
	names := configuration.ListLinkNames()
	assert.Equal(t, 1, len(names))
	assert.Equal(t, linkName, names[0])
}

func TestLinkTagAdd(t *testing.T) {
	linkName := "TestLinkTagAdd"
	rootCmd.SetArgs([]string{"add", "link", linkName})
	err := rootCmd.Execute()
	if err != nil {
		t.Error(err)
	}

	path := "/opt"
	tag := "ttt"
	rootCmd.SetArgs([]string{"add", "tag", linkName, tag, path})
	err = rootCmd.Execute()
	assert.Nil(t, err)

	tags := configuration.ListLinkValues(linkName)
	assert.Equal(t, len(tags), 1)
	assert.Equal(t, path, tags[0].Path)
	assert.Equal(t, linkName, tags[0].Name)
	assert.Equal(t, tag, tags[0].Tag)
}

func TestBindAdd(t *testing.T) {
	linkName := "TestBindAdd"
	linkName2 := "TestBindAdd2"
	commonTag := "tag"
	commonTag2 := "tag2"
	ExecuteCommand(t, "add", "link", linkName)
	ExecuteCommand(t, "add", "link", linkName2)
	ExecuteCommand(t, "add", "tag", linkName, commonTag, "/foo/bar")
	ExecuteCommand(t, "add", "tag", linkName2, commonTag2, "/foo/bar")

	ExecuteCommand(t, "add", "bind", linkName+":"+commonTag, linkName2+":"+commonTag2)

	binds := configuration.ListBinds(linkName, commonTag)
	assert.Equal(t, 1, len(binds))
	assert.Equal(t, binds[0].TargetName, linkName2)
	assert.Equal(t, binds[0].TargetTag, commonTag2)
}
