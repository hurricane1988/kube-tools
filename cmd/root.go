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

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wonderivan/logger"
	"os"

	"kube-tools/pkg/cpu"
	"kube-tools/pkg/iptables"
	"kube-tools/pkg/jwt"
	"kube-tools/pkg/memory"
	"kube-tools/pkg/metric"
	"kube-tools/pkg/ports"
	"kube-tools/pkg/stats"
	"kube-tools/pkg/traceroute"
)

// 定义根命令
var rootCmd = &cobra.Command{
	Use:   "kube-tools",
	Short: "kube-tools is a tool contains some thing",
	Long:  "kube-tools is a tool contains some thing",
	// 定义根命令执行函数
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("welcome to kube-tools")
	},
}

// Execute 定义cobra执行器
func Execute() {
	// 注册子端口扫描子命令到root根命令
	rootCmd.AddCommand(ports.ExecutePortScan())
	rootCmd.AddCommand(memory.ExecuteSimulateRaw())
	rootCmd.AddCommand(metric.ExecuteMetric())
	rootCmd.AddCommand(traceroute.ExecuteTraceroute())
	rootCmd.AddCommand(cpu.ExecuteCPUSimulateLoad())
	rootCmd.AddCommand(jwt.ExecuteJwtToken())
	rootCmd.AddCommand(iptables.ExecuteIptables())
	rootCmd.AddCommand(stats.ExecuteStatsGroup())
	if err := rootCmd.Execute(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
