package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/symbolic-link-manager/internal"
	"github.com/symbolic-link-manager/internal/configuration"
	"github.com/symbolic-link-manager/internal/logger"
)

func createUpdateLinkDeclareCmd() *cobra.Command {
	var newLinkName *string

	var updateLinkDeclareCmd = &cobra.Command{
		Use: "link LINK_NAME",
		Run: func(cmd *cobra.Command, args []string) {
			if newLinkName == nil || *newLinkName == "" {
				fmt.Println("未修改任何字段.")
				return
			}
			err := configuration.RenameLinkDeclaration(args[0], *newLinkName)
			if err != nil {
				logger.LogError(err)
				return
			}
			fmt.Printf(*newLinkName)
		},
	}
	newLinkName = updateLinkDeclareCmd.Flags().String("name", "", "想要更新的名称")

	return updateLinkDeclareCmd
}

func createUpdateLinkValueCmd() *cobra.Command {
	var newAlias *string
	var newPath *string
	var updateLinkCmd = &cobra.Command{
		Use:     "link-value LINK_NAME ALIAS",
		Aliases: []string{"lkv"},
		Run: func(cmd *cobra.Command, args []string) {
			changedCnt := 0
			var updateEntity configuration.Link
			if newAlias != nil && *newAlias != "" {
				changedCnt++
				updateEntity.Alias = *newAlias
			}
			if newPath != nil && *newPath != "" {
				changedCnt++
				updateEntity.Path = *newPath
			}
			if changedCnt == 0 {
				fmt.Println("未修改任何字段.")
				return
			}
			err := configuration.UpdateLinkValue(args[0], args[1], updateEntity)
			if err != nil {
				logger.LogError(err)
				return
			}
		},
		Args: cobra.ExactArgs(2),
	}
	newPath = updateLinkCmd.Flags().String("path", "", "新的路径")
	newAlias = updateLinkCmd.Flags().String("alias", "", "新的别名, 将会同步更新绑定.")

	return updateLinkCmd
}

func createUpdateBindCmd() *cobra.Command {
	var newTargetLinkName *string
	var newTargetLinkAlas *string

	var updateBindCmd = &cobra.Command{
		Use: "bind LINK_NAME:LINK_ALIAS TARGET_LINK_NAME:TARGET_LINK_ALIAS",
		Run: func(cmd *cobra.Command, args []string) {
			linkName, linkAlias, err := internal.SplitVersion(args[0])
			if err != nil {
				logger.LogError(err)
				return
			}
			targetName, targetAlias, err := internal.SplitVersion(args[1])
			if err != nil {
				logger.LogError(err)
				return
			}
			changedCnt := 0
			var dto = configuration.UpdateBindDTO{
				SrcName:     linkName,
				SrcAlias:    linkAlias,
				TargetName:  targetName,
				TargetAlias: targetAlias,
			}
			if newTargetLinkName != nil && *newTargetLinkName != "" {
				changedCnt++
				dto.NewName = *newTargetLinkName
			}
			if newTargetLinkAlas != nil && *newTargetLinkAlas != "" {
				changedCnt++
				dto.NewAlias = *newTargetLinkAlas
			}

			if changedCnt == 0 {
				fmt.Println("未修改任何字段.")
				return
			}
			err = configuration.UpdateBind(dto)
			if err != nil {
				logger.LogError(err)
				return
			}
		},
		Args: cobra.ExactArgs(2),
	}

	return updateBindCmd
}

func init() {
	var updateCommand = &cobra.Command{
		Use:   "update",
		Short: "更新资源。",
		Long:  "更新特定的资源。",
	}

	updateCommand.AddCommand(createUpdateLinkDeclareCmd())
	updateCommand.AddCommand(createUpdateLinkValueCmd())
	updateCommand.AddCommand(createUpdateBindCmd())
	rootCmd.AddCommand(updateCommand)
}
