package cmd

import (
	"fmt"
	"github.com/symbolic-link-manager/internal/localizer"
	"strings"

	"github.com/symbolic-link-manager/internal/configuration"
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
			names := configuration.ListLinkNames()
			fmt.Print(strings.Join(names, "\n"))
		},
	}

	var getEnvValue = &cobra.Command{
		Use:   localizer.GetMessageWithoutParam(localizer.CommandGetLKVUse),
		Short: localizer.GetMessageWithoutParam(localizer.CommandGetLKVShort),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				displayer.DisplayLinks(configuration.ListLinkValues("")...)
			} else {
				displayer.DisplayLinks(configuration.ListLinkValues(args[0])...)
			}
		},
	}

	var getBound = &cobra.Command{
		Use:   localizer.GetMessageWithoutParam(localizer.CommandGetBindUse),
		Short: localizer.GetMessageWithoutParam(localizer.CommandGetBindShort),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				binds := configuration.GetAllBinds()
				for k, v := range binds {
					displayer.DisplayBindsWithStringRoot(k, v...)
				}
				return
			}

			result := configuration.ListBinds(args[0], args[1])
			displayer.DisplayBindsWithStringRoot(args[0], result...)
		},
	}

	var getUsing = &cobra.Command{
		Use:   "using",
		Short: localizer.GetMessageWithoutParam(localizer.CommandGetUsing),
		Run: func(cmd *cobra.Command, args []string) {
			using, err := configuration.ListUsing()
			if err != nil {
				panic(err)
			}
			displayer.DisplayUsingLink(using)
		},
	}

	getCommand.AddCommand(getLinks)
	getCommand.AddCommand(getEnvValue)
	getCommand.AddCommand(getBound)
	getCommand.AddCommand(getUsing)

	rootCmd.AddCommand(getCommand)
}
