package storage

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/symbolic-link-manager/internal/localizer"
	"os"
	"path"
)

var lazyLoadAppHome string
var AppHomeEnvKey = "SLINK_MANAGER_HOME"

// AppHome 获取 lazyLoadAppHome
func AppHome() string {
	if lazyLoadAppHome != "" {
		return lazyLoadAppHome
	}
	home, ok := os.LookupEnv(AppHomeEnvKey)
	if !ok {
		panic(localizer.GetMessage(&i18n.LocalizeConfig{MessageID: "error.noenv"}))
	}
	lazyLoadAppHome = home
	return home
}

var cache *configuration = nil
var lazyLoadConfigPath string

// getConfigPath 获取配置文件路径
func getConfigPath() string {
	if lazyLoadConfigPath != "" {
		return lazyLoadConfigPath
	}
	p := path.Join(AppHome(), "configuration.json")
	lazyLoadConfigPath = p
	return p
}

// readConfig 读取配置文件, 如果有修改了数据, 应该调用 [saveConfig] 进行持久化。
func readConfig() configuration {
	if cache != nil {
		return *cache
	}
	configFilePath := getConfigPath()
	_, err := os.Stat(configFilePath)
	if os.IsNotExist(err) {
		return configuration{
			DeclaredLinkNames: make([]string, 0),
			Tags:              make([]*Tag, 0),
			Binds:             map[string][]*LinkBindItem{},
		}
	}
	content, err := os.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}
	var configuration = configuration{
		DeclaredLinkNames: make([]string, 0),
		Tags:              make([]*Tag, 0),
		Binds:             map[string][]*LinkBindItem{},
	}
	err = json.Unmarshal(content, &configuration)
	if err != nil && len(content) > 0 {
		fmt.Println("Failed to read json config.")
		panic(err)
	}
	cache = &configuration
	return configuration
}

// saveConfig 保存配置
func saveConfig(configuration *configuration) {
	configFilePath := getConfigPath()
	cache = configuration
	content, err := json.Marshal(configuration)
	if err != nil {
		fmt.Println("Failed to save json config.")
		panic(err)
	}
	err = os.WriteFile(configFilePath, content, 0b110_110_100)
	if err != nil {
		panic(err)
	}
}
