package cmd

import (
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/symbolic-link-manager/internal/localizer"
	"strings"

	"github.com/symbolic-link-manager/internal/configuration"

	"github.com/spf13/cobra"
)

// 分割字符串中的冒号
//
// 返回: (Name Alias error)
func splitVersion(nameWithVersion string) (string, string, error) {
	sp := strings.Split(nameWithVersion, ":")
	if len(sp) != 2 {
		return "", "", &localizer.LocalizedError{
			Config: &i18n.LocalizeConfig{
				MessageID: localizer.ErrorInvalidNTPair,
				TemplateData: map[string]string{
					"Raw": nameWithVersion,
				},
			},
		}
	}
	return sp[0], sp[1], nil
}

func init() {
	var addCommand = &cobra.Command{
		Use:   "add",
		Short: localizer.GetMessageWithoutParam(localizer.CommandAddShort),
		Long:  localizer.GetMessageWithoutParam(localizer.CommandAddLong),
	}

	var envAddCommand = &cobra.Command{
		Use:   localizer.GetMessageWithoutParam(localizer.CommandAddLinkUse),
		Short: localizer.GetMessageWithoutParam(localizer.CommandAddLinkShort),
		Long:  localizer.GetMessageWithoutParam(localizer.CommandAddLinkLong),
		Run: func(cmd *cobra.Command, args []string) {
			configuration.AddEnvDeclaration(args[0])
			fmt.Println(localizer.GetMessage(&i18n.LocalizeConfig{
				MessageID: localizer.CommandAddLinkSuccess,
				TemplateData: map[string]string{
					"LinkName": args[0],
				},
			}))
		},
		Args: cobra.ExactArgs(1),
	}
	var envValueAddCommand = &cobra.Command{
		Use:     localizer.GetMessageWithoutParam(localizer.CommandAddLKVUse),
		Aliases: []string{"lkv"},
		Short:   localizer.GetMessageWithoutParam(localizer.CommandAddLKVShort),
		Long:    localizer.GetMessageWithoutParam(localizer.CommandAddLKVLong),
		Example: `  slm add link-value java 17 "C:\Program Files\Java\jdk-17.0.12+7"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := configuration.AddEnvValue(&configuration.Link{
				Name:  args[0],
				Alias: args[1],
				Path:  args[2],
			})
			if err != nil {
				return err
			}
			fmt.Println(localizer.GetMessageWithoutParam(localizer.MessageSuccess))
			return nil
		},
		Args: cobra.ExactArgs(3),
	}

	var bindAddCommand = &cobra.Command{
		Use:   localizer.GetMessageWithoutParam(localizer.CommandAddBindUse),
		Short: localizer.GetMessageWithoutParam(localizer.CommandAddBindShort),
		Long:  localizer.GetMessageWithoutParam(localizer.CommandAddBindLong),
		RunE: func(cmd *cobra.Command, args []string) error {
			// errors were checked in Args function.
			srcName, srcAlias, err := splitVersion(args[0])
			if err != nil {
				return err
			}
			targetName, targetAlias, err := splitVersion(args[1])
			if err != nil {
				return err
			}
			err = configuration.AddBind(srcName, srcAlias, targetName, targetAlias)
			if err != nil {
				return err
			}
			fmt.Println(localizer.GetMessage(&i18n.LocalizeConfig{
				MessageID: localizer.CommandAddBindSuccess,
				TemplateData: map[string]string{
					"SrcName":    srcName,
					"SrcTag":     srcAlias,
					"TargetName": targetName,
					"TargetTag":  targetAlias,
				},
			}))
			return nil
		},
		Args: cobra.ExactArgs(2),
	}

	addCommand.AddCommand(envAddCommand)
	addCommand.AddCommand(envValueAddCommand)
	addCommand.AddCommand(bindAddCommand)
	rootCmd.AddCommand(addCommand)
}
