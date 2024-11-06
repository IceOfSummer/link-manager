package cmd

import (
	"fmt"

	"github.com/link-manager/internal/configuration"

	"github.com/spf13/cobra"
)

func init() {
	var useCommand = &cobra.Command{
		Use:   "use [link-name] [alias]",
		Short: "使用环境变量",
		Long:  "使用特定的环境变量",
		Run: func(cmd *cobra.Command, args []string) {
			env := configuration.FindEnvByNameAndAlias(args[0], args[1])
			if env == nil {
				return
			}
			configuration.UseEnv(env)
			fmt.Println("设置成功")
		},
		Args: cobra.MinimumNArgs(1),
	}

	rootCmd.AddCommand(useCommand)
}
