package configuration

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
)

var cache *configuration = nil
var lazyLoadConfigPath string

func getConfigPath() string {
	if lazyLoadConfigPath != "" {
		return lazyLoadConfigPath
	}
	p := path.Join(AppHome(), "configuration.json")
	lazyLoadConfigPath = p
	return p
}

// 读取配置文件
func readConfig() configuration {
	if cache != nil {
		return *cache
	}
	configFilePath := getConfigPath()
	_, err := os.Stat(configFilePath)
	if os.IsNotExist(err) {
		return configuration{
			DeclaredLinkNames: make([]string, 0),
			Links:             make([]Link, 0),
			Binds:             map[string][]LinkBindItem{},
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

// AddEnvDeclaration 添加一个环境变量声明
func AddEnvDeclaration(declarationName string) {
	config := readConfig()
	config.DeclaredLinkNames = append(config.DeclaredLinkNames, declarationName)
	saveConfig(&config)
}

func (th *configuration) isDeclarationExist(declarationName string) bool {
	p := th.findDeclaration(declarationName)
	if p == -1 {
		return false
	}
	return true
}

// 搜索变量声明的位置，如果没找到，返回 -1
func (th *configuration) findDeclaration(declarationName string) int {
	for pos, v := range th.DeclaredLinkNames {
		if v == declarationName {
			return pos
		}
	}
	return -1
}

// AddEnvValue 添加一个环境变量的值
func AddEnvValue(env *Link) error {
	config := readConfig()
	if !config.isDeclarationExist(env.Name) {
		return errors.New("对应的环境变量没有声明")
	}
	config.Links = append(config.Links, *env)
	saveConfig(&config)
	return nil
}

func getBindMapKey(name, alias string) string {
	return name + ":" + alias
}

// AddBind target 绑定到 src
func AddBind(srcName, srcAlias, targetName, targetAlias string) error {
	config := readConfig()

	old, ok := config.Binds[srcName]

	entity := LinkBindItem{
		CurrentTag:  srcAlias,
		TargetName:  targetName,
		TargetAlias: targetAlias,
	}
	if ok {
		config.Binds[srcName] = append(old, entity)
	} else {
		config.Binds[srcName] = []LinkBindItem{entity}
	}
	saveConfig(&config)
	return nil
}

func ListLinkNames() []string {
	return readConfig().DeclaredLinkNames
}

// ListLinkValues 列出所有链接的值。
//
// 当不传 [name] 时，返回所有的值
func ListLinkValues(name string) []Link {
	config := readConfig()
	if name == "" {
		return config.Links
	}
	result := make([]Link, 0)
	for _, v := range config.Links {
		if v.Name == name {
			result = append(result, v)
		}
	}
	return result
}

func FindLinkByNameAndAlias(name, alias string) *Link {
	envs := ListLinkValues(name)
	for _, v := range envs {
		if v.Tag == alias {
			return &v
		}
	}
	return nil
}

func ListBinds(item *LinkBindItem) []LinkBindItem {
	config := readConfig()

	value, ok := config.Binds[getBindMapKey(item.TargetName, item.TargetAlias)]
	if !ok {
		return make([]LinkBindItem, 0)
	}
	return value
}

func GetAllBinds() BindsData {
	return readConfig().Binds
}

func rebuildDeclaredLinks(links []Link) []string {
	var names []string
	set := make(map[string]struct{})
	for _, link := range links {
		_, ok := set[link.Name]
		if !ok {
			set[link.Name] = struct{}{}
			names = append(names, link.Name)
		}
	}
	return names
}

// DeleteLink 删除链接.
// 如果不提供第二个参数, 则删除全部.
// 返回被删除的元素
func DeleteLink(linkName, alias string) []Link {
	config := readConfig()

	var newLinks []Link
	var deleted []Link

	for _, link := range config.Links {
		if link.Name != linkName {
			newLinks = append(newLinks, link)
			continue
		}
		if alias == "" || link.Tag == alias {
			deleted = append(deleted, link)
			continue
		}
		newLinks = append(newLinks, link)
	}
	newDeclaredLinkNames := rebuildDeclaredLinks(newLinks)
	config.Links = newLinks
	config.DeclaredLinkNames = newDeclaredLinkNames
	return deleted
}

// DeleteBind 删除对应的绑定.
//
// 返回 [true] 表示删除成功
func DeleteBind(rootLinkName string, linkBindItem *LinkBindItem) bool {
	config := readConfig()

	result, ok := config.Binds[rootLinkName]
	if !ok {
		return false
	}
	for i, item := range result {
		if item.TargetName == linkBindItem.TargetName &&
			item.TargetAlias == linkBindItem.TargetAlias &&
			item.CurrentTag == linkBindItem.CurrentTag {
			result = append(result[:i], result[i+1:]...)
			config.Binds[rootLinkName] = result
			saveConfig(&config)
			return true
		}
	}
	return false
}

// RenameLinkDeclaration 重命名链接声明
// 如果没有找到旧的声明，将会返回一个错误.
func RenameLinkDeclaration(oldName, newName string) error {
	config := readConfig()
	pos := config.findDeclaration(oldName)
	if pos == -1 {
		return errors.New("旧链接声明不存在")
	}
	np := config.findDeclaration(newName)
	if np != -1 {
		return errors.New("新链接名称已经存在")
	}

	config.DeclaredLinkNames = append(config.DeclaredLinkNames[:pos], config.DeclaredLinkNames[pos+1:]...)

	for _, link := range config.Links {
		if link.Name == oldName {
			link.Name = newName
		}
	}

	_, ok := config.Binds[oldName]
	if ok {
		config.Binds[newName] = config.Binds[oldName]
		delete(config.Binds, oldName)
	}
	return nil
}

// UpdateLinkValue 更新链接值
func UpdateLinkValue(name, alias string, updateEntity Link) error {
	config := readConfig()
	link := FindLinkByNameAndAlias(name, alias)
	if link == nil {
		return errors.New("链接不存在")
	}
	if updateEntity.Tag != "" {
		link.Tag = updateEntity.Tag
	}
	if updateEntity.Path != "" {
		link.Path = updateEntity.Path
	}
	saveConfig(&config)
	return nil
}

type UpdateBindDTO struct {
	// Required
	SrcName string
	// Required
	SrcAlias string
	// Required
	TargetName string
	// Required
	TargetAlias string
	// Optional
	NewName string
	// Optional
	NewAlias string
}

// UpdateBind 更新绑定
func UpdateBind(dto UpdateBindDTO) error {
	config := readConfig()
	bind, ok := config.Binds[dto.SrcName]
	if !ok {
		return errors.New("绑定不存在")
	}
	for _, item := range bind {
		if item.TargetName == dto.TargetName && item.TargetAlias == dto.TargetAlias {
			if dto.NewName != "" {
				item.TargetName = dto.NewName
			}
			if dto.NewAlias != "" {
				item.TargetAlias = dto.NewAlias
			}
			saveConfig(&config)
			return nil
		}
	}
	return errors.New("未找到绑定的目标链接")
}
