package main

import (
	"fmt"
	"time"

	"gkkyw/lib"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/spf13/cobra"
)

var cpuCMD = &cobra.Command{
	Use:   "cpu",
	Short: "打印CPU使用率",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			p, _ := cpu.Percent(time.Second, false)
			fmt.Printf("\rCPU:%.1f%%", p[0])
			time.Sleep(time.Second)
		}
	},
}

func main() {
	// 添加Cmd
	lib.RootCmd.AddCommand(cpuCMD)

	// 执行
	lib.Execute()

}
