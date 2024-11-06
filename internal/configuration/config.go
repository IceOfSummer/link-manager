package configuration

import (
	"os"
	"path"

	"github.com/link-manager/internal/localizer"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var DENV_HOME string
var configFilePath string

func init() {
	denv, ok := os.LookupEnv("DENV_HOME")
	if !ok {
		panic(localizer.GetMessageAndIgnoreError(&i18n.LocalizeConfig{MessageID: "error.noenv"}))
	}
	DENV_HOME = denv
	configFilePath = path.Join(denv, "configuration.json")
}
