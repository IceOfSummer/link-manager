package cmd

import (
	"fmt"

	"github.com/link-manager/internal/configuration"

	"github.com/spf13/cobra"
)

func init() {
	var addCommand = &cobra.Command{
		Use:   "add",
		Short: "添加资源。",
		Long:  "添加特定的资源。",
	}

	var envAddCommand = &cobra.Command{
		Use:   "link [link-name]",
		Short: "声明一个链接",
		Long:  "声明一个链接，仅声明，没有具体的值",
		Run: func(cmd *cobra.Command, args []string) {
			configuration.AddEnvDeclarition(args[0])
			fmt.Println("添加了新的环境变量声明: " + args[0])
		},
		Args: cobra.MinimumNArgs(1),
	}
	var envValueAddCommand = &cobra.Command{
		Use:     "link-value [link-name] [Alias] [Path]",
		Aliases: []string{"lkv"},
		Short:   "给动态环境变量添加一个值",
		Long:    "给动态环境变量添加一个值，动态环境变量名称需要提前声明",
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
		Args: cobra.MinimumNArgs(3),
	}

	addCommand.AddCommand(envAddCommand)
	addCommand.AddCommand(envValueAddCommand)
	rootCmd.AddCommand(addCommand)
}
