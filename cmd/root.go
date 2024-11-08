package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/symbolic-link-manager/internal/logger"
)

const longDoc = `一个管理软连接的工具，通常用于快速切换工具 SDK 版本。

完整文档: https://github.com/IceOfSummer/symbolic-link-manager
`

var rootCmd = &cobra.Command{
	Use:   "slm",
	Short: "symbolic-link-manager 是一个管理系统软连接的工具",
	Long:  longDoc,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&logger.DebugEnable, "debug", false, "Enable debug logger.")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
