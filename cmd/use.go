package cmd

import (
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/symbolic-link-manager/internal/core"
	"github.com/symbolic-link-manager/internal/localizer"

	"github.com/spf13/cobra"
)

func init() {
	var useCommand = &cobra.Command{
		Use:   localizer.GetMessageWithoutParam(localizer.CommandUse),
		Short: localizer.GetMessageWithoutParam(localizer.CommandUseShort),
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := core.UseLink(args[0], args[1])
			if err != nil {
				return err
			}
			for _, v := range r {
				fmt.Println(localizer.GetMessage(&i18n.LocalizeConfig{
					MessageID: localizer.CommandUseSuccess,
					TemplateData: map[string]string{
						"LinkName": v.Linkname,
						"Tag":      v.TagName,
					},
				}))
			}
			return nil
		},
		Args: cobra.ExactArgs(2),
	}

	rootCmd.AddCommand(useCommand)
}
