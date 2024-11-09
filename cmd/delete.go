package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/symbolic-link-manager/internal"
	"github.com/symbolic-link-manager/internal/configuration"
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
			deleted := configuration.DeleteLink(args[0], "")
			if len(deleted) == 0 {
				return localizer.CreateNoSuchLinkError(args[0])
			}
			fmt.Println(localizer.GetMessageWithoutParam(localizer.MessageDeleteSuccessPrefix))
			displayer.DisplayLinks(deleted...)
			return nil
		},
	}

	var deleteLinkValue = &cobra.Command{
		Use:   localizer.GetMessageWithoutParam(localizer.CommandDeleteLKVUse),
		Short: localizer.GetMessageWithoutParam(localizer.CommandDeleteLKVShort),
		Long:  localizer.GetMessageWithoutParam(localizer.CommandDeleteLKVLong),
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			deleted := configuration.DeleteLink(args[0], "")
			if len(deleted) == 0 {
				return localizer.CreateNoSuchLinkError(args[0])
			}
			fmt.Println(localizer.GetMessageWithoutParam(localizer.MessageDeleteSuccessPrefix))
			displayer.DisplayLinks(deleted...)
			return nil
		},
	}

	var deleteBind = &cobra.Command{
		Use:   localizer.GetMessageWithoutParam(localizer.CommandDeleteBindUse),
		Short: localizer.GetMessageWithoutParam(localizer.CommandDeleteBindShort),
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			srcName, srcAlias, err := internal.SplitVersion(args[0])
			if err != nil {
				return err
			}
			targetName, targetAlias, err := internal.SplitVersion(args[1])
			if err != nil {
				return err
			}
			item := configuration.LinkBindItem{
				CurrentTag:  srcAlias,
				TargetName:  targetName,
				TargetAlias: targetAlias,
			}
			result := configuration.DeleteBind(srcName, &item)
			if result {
				fmt.Println(localizer.GetMessageWithoutParam(localizer.MessageDeleteSuccessPrefix))
				displayer.DisplayBindsWithStringRoot(srcName, item)
			} else {
				return localizer.CreateNoSuchBindError()
			}
			return nil
		},
	}

	deleteCommand.AddCommand(deleteLink)
	deleteCommand.AddCommand(deleteLinkValue)
	deleteCommand.AddCommand(deleteBind)
	rootCmd.AddCommand(deleteCommand)

}
