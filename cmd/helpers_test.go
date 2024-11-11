// Helper functions
package cmd

import (
	"github.com/symbolic-link-manager/internal/configuration"
	"os"
	"path"
	"runtime"
)

func SetUpTestEnvironment() {
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(path.Join(filename, ".."))

	err := os.Setenv(configuration.APP_HOME_ENV_KEY, path.Join(root, "tmp"))
	if err != nil {
		panic(err)
	}
	_ = os.Mkdir(configuration.AppHome(), 0b111_101_000)
}

func CleanUp() {
	_ = os.Remove(path.Join(configuration.AppHome(), "configuration.json"))
}
