/*
Copyright 2023 QKP Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cpu

import (
	"github.com/jedib0t/go-pretty/v6/table"
	cpuv3 "github.com/shirou/gopsutil/v3/cpu"
	"github.com/spf13/cobra"
	"os"
)

// ExecuteStatsCpu 定义端口扫描执行器
func ExecuteStatsCpu() *cobra.Command {
	// statCpu 定义端口扫描子命令
	var statCpu = &cobra.Command{
		Use:   "cpu",
		Short: "stats cpu information",
		Long:  "stats cpu information",
		Run:   runner,
	}
	statCpu.Flags().StringP("choice", "C", "sum", "stats cpu information")
	return statCpu
}

// runner 扫描端口
func runner(cmd *cobra.Command, args []string) {
	choice, _ := cmd.Flags().GetString("choice")
	// 初始化table对象
	t := table.NewWriter()
	// 设置输出到终端
	t.SetOutputMirror(os.Stdout)
	switch choice {
	case "sum":
		t.AppendHeader(table.Row{"CPU", "Vendor", "Physical Cores", "Logical Cores", "MHz"})
		// 获取CPU信息
		c, _ := cpuv3.Info()
		// 添加表格的数据行
		t.AppendRows([]table.Row{
			{"", c[0].ModelName, "", c[0].Cores, c[0].Mhz},
		})
		t.AppendSeparator()
	case "detail":
		t.AppendHeader(table.Row{
			"cpu-id",
			"user",
			"system",
			"idle",
			"nice",
			"iowait",
			"hardware-interrupt",
			"software-interrupt",
			"vm-steal-time",
			"vm-guest-time",
			"vm-guest-nice-time",
		})
		cpus, _ := cpuv3.Times(true)
		for _, c := range cpus {
			t.AppendRows([]table.Row{
				{c.CPU, c.User, c.System, c.Idle, c.Nice, c.Iowait, c.Irq, c.Softirq, c.Steal, c.Guest, c.GuestNice},
			})
		}
	}
	t.Render()
}
