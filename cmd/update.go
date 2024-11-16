package cmd

import (
	"fmt"
	"github.com/symbolic-link-manager/internal/core"
	"github.com/symbolic-link-manager/internal/storage"

	"github.com/spf13/cobra"
	"github.com/symbolic-link-manager/internal"
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
			err := core.RenameLink(args[0], *newLinkName)
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

func createUpdateTagCmd() *cobra.Command {
	var newPath *string
	var updateLinkCmd = &cobra.Command{
		Use:   localizer.GetMessageWithoutParam(localizer.CommandUpdateTagUse),
		Short: localizer.GetMessageWithoutParam(localizer.CommandUpdateTagShort),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := core.UpdateTag(&storage.Tag{
				Linkname: args[0],
				TagName:  args[1],
				Path:     *newPath,
			})
			if err != nil {
				return err
			}
			fmt.Printf(localizer.GetMessageWithoutParam(localizer.MessageSuccess))
			return nil
		},
		Args: cobra.ExactArgs(2),
	}
	newPath = updateLinkCmd.Flags().String("path", "", localizer.GetMessageWithoutParam(localizer.UpdateFlagPath))

	return updateLinkCmd
}

func createUpdateBindCmd() *cobra.Command {
	var newTargetLinkTag *string

	var updateBindCmd = &cobra.Command{
		Use:   localizer.GetMessageWithoutParam(localizer.CommandUpdateBindUse),
		Short: localizer.GetMessageWithoutParam(localizer.CommandUpdateBindShort),
		RunE: func(cmd *cobra.Command, args []string) error {
			if *newTargetLinkTag == "" {
				return localizer.CreateError(localizer.NothingChanged)
			}

			srcName, srcTag, err := internal.SplitVersion(args[0])
			if err != nil {
				return err
			}
			targetName, targetTag, err := internal.SplitVersion(args[1])
			if err != nil {
				return err
			}

			err = core.UpdateBind(&core.UpdateBindDTO{
				SrcName:    srcName,
				SrcTag:     srcTag,
				TargetName: targetName,
				TargetTag:  targetTag,
				NewTag:     *newTargetLinkTag,
			})
			if err != nil {
				return err
			}
			fmt.Printf(localizer.GetMessageWithoutParam(localizer.MessageSuccess))
			return nil
		},
		Args: cobra.ExactArgs(2),
	}
	newTargetLinkTag = updateBindCmd.Flags().String("targetTag", "", "New target tag")
	return updateBindCmd
}

func init() {
	var updateCommand = &cobra.Command{
		Use:   "update",
		Short: localizer.GetMessageWithoutParam(localizer.CommandUpdateShort),
		Long:  localizer.GetMessageWithoutParam(localizer.CommandUpdateLong),
	}

	updateCommand.AddCommand(createUpdateLinkDeclareCmd())
	updateCommand.AddCommand(createUpdateTagCmd())
	updateCommand.AddCommand(createUpdateBindCmd())
	rootCmd.AddCommand(updateCommand)
}
