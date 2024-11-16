package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/symbolic-link-manager/internal"
	"github.com/symbolic-link-manager/internal/configuration"
	"github.com/symbolic-link-manager/internal/localizer"
)

func createUpdateLinkDeclareCmd() *cobra.Command {
	var newLinkName *string

	var updateLinkDeclareCmd = &cobra.Command{
		Use:   localizer.GetMessageWithoutParam(localizer.CommandUpdateLinkUse),
		Short: localizer.GetMessageWithoutParam(localizer.CommandUpdateLinkShort),
		RunE: func(cmd *cobra.Command, args []string) error {
			if newLinkName == nil || *newLinkName == "" {
				return localizer.CreateError(localizer.NothingChanged)
			}
			err := configuration.RenameLinkDeclaration(args[0], *newLinkName)
			if err != nil {
				return err
			}
			fmt.Printf(*newLinkName)
			return nil
		},
	}
	newLinkName = updateLinkDeclareCmd.Flags().String(
		"name",
		"",
		localizer.GetMessageWithoutParam(localizer.UpdateFlagName),
	)

	return updateLinkDeclareCmd
}

func createUpdateLinkValueCmd() *cobra.Command {
	var newTag *string
	var newPath *string
	var updateLinkCmd = &cobra.Command{
		Use:   localizer.GetMessageWithoutParam(localizer.CommandUpdateLKVUse),
		Short: localizer.GetMessageWithoutParam(localizer.CommandUpdateLKVShort),
		RunE: func(cmd *cobra.Command, args []string) error {
			changedCnt := 0
			var updateEntity configuration.Link
			if newTag != nil && *newTag != "" {
				changedCnt++
				updateEntity.Tag = *newTag
			}
			if newPath != nil && *newPath != "" {
				changedCnt++
				updateEntity.Path = *newPath
			}
			if changedCnt == 0 {
				return localizer.CreateError(localizer.NothingChanged)
			}
			err := configuration.UpdateTag(args[0], args[1], updateEntity)
			if err != nil {
				return err
			}
			return nil
		},
		Args: cobra.ExactArgs(2),
	}
	newPath = updateLinkCmd.Flags().String("path", "", localizer.GetMessageWithoutParam(localizer.UpdateFlagPath))
	newTag = updateLinkCmd.Flags().String("tag", "", localizer.GetMessageWithoutParam(localizer.UpdateFlagTag))

	return updateLinkCmd
}

func createUpdateBindCmd() *cobra.Command {
	var newTargetLinkName *string
	var newTargetLinkTag *string

	var updateBindCmd = &cobra.Command{
		Use:   localizer.GetMessageWithoutParam(localizer.CommandUpdateBindUse),
		Short: localizer.GetMessageWithoutParam(localizer.CommandUpdateBindShort),
		RunE: func(cmd *cobra.Command, args []string) error {
			linkName, linkAlias, err := internal.SplitVersion(args[0])
			if err != nil {
				return err
			}
			targetName, targetAlias, err := internal.SplitVersion(args[1])
			if err != nil {
				return err
			}
			changedCnt := 0
			var dto = configuration.UpdateBindDTO{
				SrcName:    linkName,
				SrcTag:     linkAlias,
				TargetName: targetName,
				TargetTag:  targetAlias,
			}
			if newTargetLinkName != nil && *newTargetLinkName != "" {
				changedCnt++
				dto.NewName = *newTargetLinkName
			}
			if newTargetLinkTag != nil && *newTargetLinkTag != "" {
				changedCnt++
				dto.NewAlias = *newTargetLinkTag
			}

			if changedCnt == 0 {
				return localizer.CreateError(localizer.NothingChanged)
			}
			err = configuration.UpdateBind(dto)
			if err != nil {
				return err
			}
			return nil
		},
		Args: cobra.ExactArgs(2),
	}
	// TODO: i18n
	newTargetLinkTag = updateBindCmd.Flags().String("targetTag", "", "target tag")
	newTargetLinkName = updateBindCmd.Flags().String("targetName", "", "target name")
	return updateBindCmd
}

func init() {
	var updateCommand = &cobra.Command{
		Use:   "update",
		Short: localizer.GetMessageWithoutParam(localizer.CommandUpdateShort),
		Long:  localizer.GetMessageWithoutParam(localizer.CommandUpdateLong),
	}

	updateCommand.AddCommand(createUpdateLinkDeclareCmd())
	updateCommand.AddCommand(createUpdateLinkValueCmd())
	updateCommand.AddCommand(createUpdateBindCmd())
	rootCmd.AddCommand(updateCommand)
}
