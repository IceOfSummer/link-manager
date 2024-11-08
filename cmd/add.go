package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/symbolic-link-manager/internal/configuration"

	"github.com/spf13/cobra"
)

// 分割字符串中的冒号
func splitVersion(nameWithVersion string) (*configuration.LinkBindItem, error) {
	sp := strings.Split(nameWithVersion, ":")
	if len(sp) != 2 {
		return nil, errors.New("` " + nameWithVersion + "` 存在多个 `:`，无法解析别名!")
	}
	return &configuration.LinkBindItem{Name: sp[0], Alias: sp[1]}, nil
}

func init() {
	var addCommand = &cobra.Command{
		Use:   "add",
		Short: "添加资源。",
		Long:  "添加特定的资源。",
	}

	var envAddCommand = &cobra.Command{
		Use:   "link LINK_NAME",
		Short: "声明一个链接",
		Long:  "声明一个链接，仅声明，没有具体的值",
		Run: func(cmd *cobra.Command, args []string) {
			configuration.AddEnvDeclarition(args[0])
			fmt.Println("添加了新的环境变量声明: " + args[0])
		},
		Args: cobra.ExactArgs(1),
	}
	var envValueAddCommand = &cobra.Command{
		Use:     "link-value LINK_NAME ALIAS PATH",
		Aliases: []string{"lkv"},
		Short:   "给动态环境变量添加一个值. 在添加时请注意路径之间的空格!",
		Long:    "给动态环境变量添加一个值(动态环境变量名称需要提前声明)，在添加时请注意路径之间的空格，如果路径之间有空格，则需要用引号包裹。",
		Example: `  lkm add link-value java 8 "C:\Program Files\Java\jdk-17.0.12+7"`,
		Run: func(cmd *cobra.Command, args []string) {
			err := configuration.AddEnvValue(&configuration.Link{
				Name:  args[0],
				Alias: args[1],
				Path:  args[2],
			})
			if err != nil {
				LogError("环境变量或别名不存在!")
				return
			}
			fmt.Println("设置成功!")
		},
		Args: cobra.ExactArgs(3),
	}

	var bindAddCommand = &cobra.Command{
		Use:   "bind LINK_NAME:ALIAS TARGET_LINK_NAME:ALIAS",
		Short: "单向绑定两个链接",
		Long:  "单向绑定两个链接，当切换到 `LINK_NAME:ALIAS` 后会自动切换到 `TARGET_LINK_NAME:ALIAS`",
		Run: func(cmd *cobra.Command, args []string) {
			// errors were checked in Args function.
			src, _ := splitVersion(args[0])
			target, _ := splitVersion(args[1])
			configuration.AddLink(src, target)
			fmt.Println("Bound: " + src.Name + ":" + src.Alias + " ==>" + target.Name + ":" + target.Alias)
		},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return errors.New("需要 2 个参数，但是提供了" + string(len(args)) + "个")
			}
			_, err := splitVersion(args[0])
			if err != nil {
				return err
			}
			_, err = splitVersion(args[1])
			if err != nil {
				return err
			}
			return nil
		},
	}

	addCommand.AddCommand(envAddCommand)
	addCommand.AddCommand(envValueAddCommand)
	addCommand.AddCommand(bindAddCommand)
	rootCmd.AddCommand(addCommand)
}
