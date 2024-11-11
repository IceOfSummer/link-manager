package cmd

import (
	"fmt"
	"github.com/symbolic-link-manager/internal"
	"github.com/symbolic-link-manager/internal/localizer"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "slm",
	Short: localizer.GetMessageWithoutParam(localizer.CommandRootShort),
	Long:  localizer.GetMessageWithoutParam(localizer.CommandRootLong),
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&internal.DebugEnable, "debug", internal.DebugEnable, "Enable debug logger.")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
