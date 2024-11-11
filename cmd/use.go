package cmd

import (
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/symbolic-link-manager/internal/configuration"
	"github.com/symbolic-link-manager/internal/localizer"

	"github.com/spf13/cobra"
)

func init() {
	var useCommand = &cobra.Command{
		Use:   localizer.GetMessageWithoutParam(localizer.CommandUse),
		Short: localizer.GetMessageWithoutParam(localizer.CommandUseShort),
		RunE: func(cmd *cobra.Command, args []string) error {
			env := configuration.FindLinkByNameAndAlias(args[0], args[1])
			if env == nil {
				return localizer.CreateNoSuchLinkValueError(args[0], args[1])
			}
			recover()
			r := configuration.UseLink(env)
			for _, v := range r {
				fmt.Println(localizer.GetMessage(&i18n.LocalizeConfig{
					MessageID: localizer.CommandUseSuccess,
					TemplateData: map[string]string{
						"LinkName": v.Name,
						"Tag":      v.Alias,
					},
				}))
			}
			return nil
		},
		Args: cobra.MinimumNArgs(1),
	}

	rootCmd.AddCommand(useCommand)
}
