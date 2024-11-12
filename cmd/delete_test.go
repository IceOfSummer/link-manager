package cmd

import (
	"github.com/stretchr/testify/assert"
	"github.com/symbolic-link-manager/internal/configuration"
	"testing"
)

func TestLinkDelete(t *testing.T) {
	linkName := "TestLinkDelete"
	ExecuteCommand(t, "add", "link", linkName)
	old := configuration.ListLinkNames()
	assert.Equal(t, 1, len(old))

	ExecuteCommand(t, "delete", "link", linkName)
	names := configuration.ListLinkNames()
	assert.Equal(t, 0, len(names))
}
