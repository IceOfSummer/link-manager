package cmd

import (
	"fmt"
	"github.com/symbolic-link-manager/internal/core"
	"github.com/symbolic-link-manager/internal/storage"

	"github.com/spf13/cobra"
	"github.com/symbolic-link-manager/internal"
	"github.com/symbolic-link-manager/internal/localizer"
	"github.com/symbolic-link-manager/internal/logger/displayer"
)

func init() {

	var deleteCommand = &cobra.Command{
		Use:   "delete",
		Short: localizer.GetMessageWithoutParam(localizer.CommandDeleteShort),
		Long:  localizer.GetMessageWithoutParam(localizer.CommandDeleteLong),
	}

	var deleteLink = &cobra.Command{
		Use:   localizer.GetMessageWithoutParam(localizer.CommandDeleteLinkUse),
		Short: localizer.GetMessageWithoutParam(localizer.CommandDeleteLinkShort),
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			deleted, err := core.RemoveLink(args[0])
			if err != nil {
				return err
			}
			fmt.Println(localizer.GetMessageWithoutParam(localizer.MessageDeleteSuccessPrefix))
			displayer.DisplayLinks(deleted...)
			return nil
		},
	}

	var deleteTag = &cobra.Command{
		Use:   localizer.GetMessageWithoutParam(localizer.CommandDeleteTagUse),
		Short: localizer.GetMessageWithoutParam(localizer.CommandDeleteTagShort),
		Long:  localizer.GetMessageWithoutParam(localizer.CommandDeleteTagLong),
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			deleted, err := core.RemoveTag(args[0], args[1])
			if err != nil {
				return err
			}
			fmt.Println(localizer.GetMessageWithoutParam(localizer.MessageDeleteSuccessPrefix))
			displayer.DisplayLinks(deleted)
			return nil
		},
	}

	var deleteBind = &cobra.Command{
		Use:   localizer.GetMessageWithoutParam(localizer.CommandDeleteBindUse),
		Short: localizer.GetMessageWithoutParam(localizer.CommandDeleteBindShort),
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			srcName, currentTag, err := internal.SplitVersion(args[0])
			if err != nil {
				return err
			}
			targetName, targetTag, err := internal.SplitVersion(args[1])
			if err != nil {
				return err
			}

			item := &storage.LinkBindItem{
				CurrentTag: currentTag,
				TargetName: targetName,
				TargetTag:  targetTag,
			}
			err = core.RemoveBind(srcName, item)
			if err != nil {
				return err
			}
			fmt.Println(localizer.GetMessageWithoutParam(localizer.MessageDeleteSuccessPrefix))
			displayer.DisplayBindsWithStringRoot(srcName, item)
			return nil
		},
	}

	deleteCommand.AddCommand(deleteLink)
	deleteCommand.AddCommand(deleteTag)
	deleteCommand.AddCommand(deleteBind)
	rootCmd.AddCommand(deleteCommand)

}
