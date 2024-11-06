package localizer

import (
	"embed"
	"os"
	"os/exec"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle = i18n.NewBundle(language.English)

//go:embed messages.zh.toml
var localZhFs embed.FS

//go:embed messages.en.toml
var localEnFs embed.FS

func init() {
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	_, err := bundle.LoadMessageFileFS(localZhFs, "messages.zh.toml")
	if err != nil {
		panic(err)
	}
	_, err = bundle.LoadMessageFileFS(localEnFs, "messages.en.toml")
	if err != nil {
		panic(err)
	}
}

// 获取当前语言.
func getLocale() string {
	envlang, ok := os.LookupEnv("LANG")
	if ok {
		return strings.Split(envlang, ".")[0]
	}

	// Exec powershell Get-Culture on Windows.
	cmd := exec.Command("powershell", "Get-Culture | select -exp Name")
	output, err := cmd.Output()
	if err == nil {
		return strings.Trim(string(output), "\r\n")
	}

	return "en_US"
}

func GetLocalizer() *i18n.Localizer {
	return i18n.NewLocalizer(bundle, getLocale())
}

func LocalizeAndIgnoreError(localizer *i18n.Localizer, config *i18n.LocalizeConfig) string {
	r, err := localizer.Localize(config)
	if err == nil {
		return r
	}
	return config.MessageID
}

func GetMessageAndIgnoreError(config *i18n.LocalizeConfig) string {
	return LocalizeAndIgnoreError(GetLocalizer(), config)
}
