package cmd

import (
	"fmt"
	"strings"

	"github.com/symbolic-link-manager/internal/configuration"

	"github.com/spf13/cobra"
)

func init() {
	var useCommand = &cobra.Command{
		Use:   "use [link-name] [alias]",
		Short: "使用环境变量",
		Long:  "使用特定的环境变量",
		Run: func(cmd *cobra.Command, args []string) {
			env := configuration.FindLinkByNameAndAlias(args[0], args[1])
			if env == nil {
				fmt.Println("链接或别名不存在")
				return
			}
			r := configuration.UseLink(env)
			var builder strings.Builder
			for _, v := range r {
				builder.WriteString("切换为: ")
				builder.WriteString(v.String())
				builder.WriteString("\n")
			}
			fmt.Print(builder.String())
		},
		Args: cobra.MinimumNArgs(1),
	}

	rootCmd.AddCommand(useCommand)
}
