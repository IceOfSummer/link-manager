package cmd

import (
	"fmt"
	"github.com/symbolic-link-manager/internal/core"
	"strings"

	"github.com/symbolic-link-manager/internal/localizer"

	"github.com/symbolic-link-manager/internal/logger/displayer"

	"github.com/spf13/cobra"
)

func init() {
	var getCommand = &cobra.Command{
		Use:   "get",
		Short: localizer.GetMessageWithoutParam(localizer.CommandGetShort),
	}

	var getLinks = &cobra.Command{
		Use:   "links",
		Short: localizer.GetMessageWithoutParam(localizer.CommandGetLinksShort),
		Run: func(cmd *cobra.Command, args []string) {
			names := core.ListDeclaredLinkNames()
			fmt.Print(strings.Join(names, "\n"))
		},
	}

	var getTags = &cobra.Command{
		Use:   localizer.GetMessageWithoutParam(localizer.CommandGetTagUse),
		Short: localizer.GetMessageWithoutParam(localizer.CommandGetTagShort),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				displayer.DisplayLinks(core.ListTags("")...)
			} else {
				displayer.DisplayLinks(core.ListTags(args[0])...)
			}
		},
	}

	var getBound = &cobra.Command{
		Use:   localizer.GetMessageWithoutParam(localizer.CommandGetBindUse),
		Short: localizer.GetMessageWithoutParam(localizer.CommandGetBindShort),
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) == 0 {
				displayer.DisplayBindsVO(core.ListBinds(""))
			} else {
				displayer.DisplayBindsVO(core.ListBinds(args[0]))
			}
		},
		Args: cobra.MaximumNArgs(1),
	}

	var getUsing = &cobra.Command{
		Use:   "using",
		Short: localizer.GetMessageWithoutParam(localizer.CommandGetUsing),
		RunE: func(cmd *cobra.Command, args []string) error {
			using, err := core.ListUsing()
			if err != nil {
				return err
			}
			displayer.DisplayUsingLink(using)
			return nil
		},
	}

	getCommand.AddCommand(getLinks)
	getCommand.AddCommand(getTags)
	getCommand.AddCommand(getBound)
	getCommand.AddCommand(getUsing)

	rootCmd.AddCommand(getCommand)
}
