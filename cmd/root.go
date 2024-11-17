package cmd

import (
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
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&internal.DebugEnable, "debug", internal.DebugEnable, "Enable debug logger.")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		println(err.Error())
		println()
		println(localizer.GetMessageWithoutParam(localizer.MessageHelp))
		os.Exit(1)
	}
}
