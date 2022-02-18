package lib

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(InfoCmd)
}

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "打印CPU、内存信息",
	Run: func(cmd *cobra.Command, args []string) {
		// 表头
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"名字", "数量", "已使用量"})

		// cpu
		cpuPercent, _ := cpu.Percent(time.Second, false)
		cpuCount, _ := cpu.Counts(true)

		// 内存
		v, _ := mem.VirtualMemory()

		// 内容
		data := [][]string{
			[]string{"CPU", fmt.Sprintf("%d核", cpuCount), fmt.Sprintf("%.1f%%", cpuPercent[0])},
			[]string{"Memory", fmt.Sprintf("%dG", v.Total/1024/1024/1024), fmt.Sprintf("%.1f%%", v.UsedPercent)},
		}

		// 遍历
		for _, v := range data {
			table.Append(v)
		}
		table.Render()
	},
}