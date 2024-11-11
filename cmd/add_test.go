package cmd

import (
	"os"
	"testing"
)

func TestLinkDeclarationAdd(t *testing.T) {
	linkName := "TestLinkDeclarationAdd"
	rootCmd.SetArgs([]string{"add", "link", linkName})
	err := rootCmd.Execute()
	if err != nil {
		t.Error(err)
	}

}

func TestMain(m *testing.M) {
	SetUpTestEnvironment()
	code := m.Run()
	CleanUp()
	os.Exit(code)
}
