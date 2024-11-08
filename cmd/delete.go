package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/symbolic-link-manager/internal"
	"github.com/symbolic-link-manager/internal/configuration"
	"github.com/symbolic-link-manager/internal/logger"
)

func init() {

	var deleteCommand = &cobra.Command{
		Use:   "delete",
		Short: "删除资源。",
		Long:  "删除特定的资源。",
	}

	var deleteLink = &cobra.Command{
		Use:   "link LINK_NAME",
		Short: "删除链接定义",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			configuration.DeleteLink(args[0], "")
		},
	}

	var deleteLinkValue = &cobra.Command{
		Use:     "link-value LINK_NAME [ALIAS]",
		Aliases: []string{"lkv"},
		Short:   "删除链接的值",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			configuration.DeleteLink(args[0], args[1])
		},
	}

	var deleteBind = &cobra.Command{
		Use:   "bind LINK_NAME:ALIAS TARGET_LINK_NAME:ALIAS",
		Short: "删除链接绑定",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			srcName, srcAlias, err := internal.SplitVersion(args[0])
			if err != nil {
				logger.LogError(err)
				return
			}
			targetName, targetAlias, err := internal.SplitVersion(args[1])
			if err != nil {
				logger.LogError(err)
				return
			}
			result := configuration.DeleteBind(srcName, &configuration.LinkBindItem{
				CurrentAlias: srcAlias,
				TargetName:   targetName,
				TargetAlias:  targetAlias,
			})
			if result {
				fmt.Printf("删除成功")
			} else {
				fmt.Println("指定的绑定不存在")
			}
		},
	}

	deleteCommand.AddCommand(deleteLink)
	deleteCommand.AddCommand(deleteLinkValue)
	deleteCommand.AddCommand(deleteBind)
	rootCmd.AddCommand(deleteCommand)

}
