package configuration

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Link struct {
	// 链接名称
	Name string
	// 链接别名
	Alias string
	// 链接路径
	Path string
}

func (t Link) String() string {
	return t.Name + ":" + t.Alias
}

type LinkBindItem struct {
	Name  string
	Alias string
}

func (t LinkBindItem) String() string {
	return t.Name + ":" + t.Alias
}

type BindsData map[string][]LinkBindItem

type configuration struct {
	DeclariedLinkNames []string
	Envs               []Link
	// Env.Name:Env.Alias -> Env
	Binds BindsData
}

var cache *configuration = nil

// 读取配置文件
func readConfig() configuration {
	if cache != nil {
		return *cache
	}
	_, err := os.Stat(configFilePath)
	if os.IsNotExist(err) {
		return configuration{
			DeclariedLinkNames: make([]string, 0),
			Envs:               make([]Link, 0),
			Binds:              map[string][]LinkBindItem{},
		}
	}
	content, err := os.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}
	var configuration configuration
	err = json.Unmarshal([]byte(content), &configuration)
	if err != nil {
		fmt.Println("Failed to read json config.")
		panic(err)
	}
	cache = &configuration
	return configuration
}

func saveConfig(configuration *configuration) {
	cache = configuration
	content, err := json.Marshal(configuration)
	if err != nil {
		fmt.Println("Failed to save json config.")
		panic(err)
	}
	os.WriteFile(configFilePath, content, 0b110_110_100)
}

// 添加一个环境变量声明
func AddEnvDeclarition(declaritionName string) {
	config := readConfig()
	config.DeclariedLinkNames = append(config.DeclariedLinkNames, declaritionName)
	saveConfig(&config)
}

func (th *configuration) isDeclaritionExist(declarationName string) bool {
	for _, v := range th.DeclariedLinkNames {
		if v == declarationName {
			return true
		}
	}
	return false
}

// 添加一个环境变量的值
func AddEnvValue(env *Link) error {
	config := readConfig()
	if !config.isDeclaritionExist(env.Name) {
		return errors.New("对应的环境变量没有声明")
	}
	config.Envs = append(config.Envs, *env)
	saveConfig(&config)
	return nil
}

func getBindMapKey(name, alias string) string {
	return name + ":" + alias
}

// target 绑定到 src
func AddLink(src, target *LinkBindItem) error {
	config := readConfig()
	key := getBindMapKey(src.Name, src.Alias)
	old, ok := config.Binds[key]
	if ok {
		config.Binds[key] = append(old, *target)
	} else {
		config.Binds[key] = []LinkBindItem{*target}
	}
	saveConfig(&config)
	return nil
}

func ListLinkNames() []string {
	return readConfig().DeclariedLinkNames
}

// 列出所有链接的值。当不传 name 时，返回所有的值
func ListLinkValues(name string) []Link {
	config := readConfig()
	if name == "" {
		return config.Envs
	}
	result := make([]Link, 0)
	for _, v := range config.Envs {
		if v.Name == name {
			result = append(result, v)
		}
	}
	return result
}

func FindEnvByNameAndAlias(name, aliase string) *Link {
	envs := ListLinkValues(name)
	for _, v := range envs {
		if v.Alias == aliase {
			return &v
		}
	}
	return nil
}

func ListBinds(item *LinkBindItem) []LinkBindItem {
	config := readConfig()

	value, ok := config.Binds[getBindMapKey(item.Name, item.Alias)]
	if !ok {
		return make([]LinkBindItem, 0)
	}
	return value
}

func GetAllBinds() BindsData {
	return readConfig().Binds
}
