// Helper functions
package cmd

import (
	"github.com/stretchr/testify/assert"
	"github.com/symbolic-link-manager/internal/configuration"
	"os"
	"path"
	"runtime"
	"testing"
)

func SetUpTestEnvironment() {
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(path.Join(filename, ".."))

	err := os.Setenv(configuration.AppHomeEnvKey, path.Join(root, "tmp"))
	if err != nil {
		panic(err)
	}
	_ = os.Mkdir(configuration.AppHome(), 0b111_101_000)
}

func CleanUp() {
	_ = os.Remove(path.Join(configuration.AppHome(), "configuration.json"))
}

func ExecuteCommand(t *testing.T, args ...string) {
	rootCmd.SetArgs(args)
	assert.Nil(t, rootCmd.Execute())
}

func TestMain(m *testing.M) {
	SetUpTestEnvironment()
	code := m.Run()
	CleanUp()
	os.Exit(code)
}
