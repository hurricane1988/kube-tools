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

package stats

import (
	"github.com/spf13/cobra"

	"kube-tools/pkg/stats/cpu"
	"kube-tools/pkg/stats/net"
)

// ExecuteStatsGroup 定义stats命令
func ExecuteStatsGroup() *cobra.Command {
	var statsGroup = &cobra.Command{
		Use:   "stats",
		Short: "stats system cpu;memory;disk;process;docker;net information.",
		Long:  "stats system cpu;memory;disk;process;docker;net information.",
		// 定义根命令执行函数
		Run: func(cmd *cobra.Command, args []string) {},
	}

	/* 注册stats相关子命令 */
	statsGroup.AddCommand(cpu.ExecuteStatsCpu())
	statsGroup.AddCommand(net.ExecuteStatsNet())
	return statsGroup
}
