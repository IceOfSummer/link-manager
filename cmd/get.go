package cmd

import (
	"fmt"
	"strings"

	"github.com/symbolic-link-manager/internal/configuration"
	"github.com/symbolic-link-manager/internal/logger/displayer"

	"github.com/spf13/cobra"
)

func init() {
	var getCommand = &cobra.Command{
		Use:   "get",
		Short: "列出指定资源",
		Long:  "列出指定资源",
	}

	var getEnv = &cobra.Command{
		Use:   "links",
		Short: "列出所有声明的链接",
		Run: func(cmd *cobra.Command, args []string) {
			names := configuration.ListLinkNames()
			fmt.Print(strings.Join(names, "\n"))
		},
	}

	var getEnvValue = &cobra.Command{
		Use:     "link-values [LINK_NAME]",
		Short:   "列出所有链接的值",
		Aliases: []string{"lkv"},
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				displayer.DisplayLinks(configuration.ListLinkValues("")...)
			} else {
				displayer.DisplayLinks(configuration.ListLinkValues(args[0])...)
			}
		},
	}

	var getBound = &cobra.Command{
		Use:   "bind ([LINK_NAME] [LINK_ALIAS])",
		Short: "获取链接所有的绑定",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				binds := configuration.GetAllBinds()
				for k, v := range binds {
					displayer.DisplayBindsWithStringRoot(k, v...)
				}
				return
			}

			root := &configuration.LinkBindItem{
				Name:  args[0],
				Alias: args[1],
			}
			result := configuration.ListBinds(root)
			displayer.DisplayBinds(root, result...)
		},
	}

	var getUsing = &cobra.Command{
		Use:   "using",
		Short: "获取所有当前正在使用的链接",
		Run: func(cmd *cobra.Command, args []string) {
			using, err := configuration.ListUsing()
			if err != nil {
				panic(err)
			}
			displayer.DisplayUsingLink(using)

		},
	}

	getCommand.AddCommand(getEnv)
	getCommand.AddCommand(getEnvValue)
	getCommand.AddCommand(getBound)
	getCommand.AddCommand(getUsing)

	rootCmd.AddCommand(getCommand)
}
