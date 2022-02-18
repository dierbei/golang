package lib

import (
	"fmt"
	"github.com/spf13/cobra"

	"os"
)

// RootCmd
// 初始化全局Cmd，在包内任何变量和函数被外部调用的时候会调用init函数添加Cmd
// 最后调用Execute
var RootCmd = &cobra.Command{
	Use:   "jt-devops",
	Short: "运维小工具",
	Long:  `运维小工具`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.1.0")
	},
}
