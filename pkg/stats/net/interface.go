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

package net

import (
	"github.com/jedib0t/go-pretty/v6/table"
	netv3 "github.com/shirou/gopsutil/v3/net"
	"github.com/spf13/cobra"
	"os"

	"kube-tools/utils/common"
)

// executeNetInterfaceGroup 定义stats命令
func executeNetInterfaceGroup() *cobra.Command {
	var interfaceCommand = &cobra.Command{
		Use:   "interface",
		Short: "show the network interface of the system.",
		Long:  "show the network interface of the system.",
		// 定义根命令执行函数
		Run: netInterfaceGroup,
	}
	/* 注册interface相关子命令 */
	return interfaceCommand
}

func netInterfaceGroup(cmd *cobra.Command, args []string) {
	listNetInterface()
}

// 获取网卡信息
func listNetInterface() []netv3.InterfaceStat {
	// 初始化table对象
	t := table.NewWriter()
	// 设置输出到终端
	t.SetOutputMirror(os.Stdout)
	// 设置表头header
	t.AppendHeader(table.Row{"adapter name", "mtu", "adapter mac", "ip address", "net up", "multicast", "pointtopoint"})
	nets, _ := netv3.Interfaces()
	for _, net := range nets {
		// 添加表格行内容
		t.AppendRows([]table.Row{
			{
				net.Name,
				net.MTU,
				net.HardwareAddr,
				common.FormatAddr(net.Addrs),
				common.Find(net.Flags, "up"),
				common.Find(net.Flags, "multicast"),
				common.Find(net.Flags, "pointtopoint"),
			},
		})
		// 自动添加序号
		t.SetAutoIndex(true)
		t.AppendSeparator()
	}
	t.Render()
	return nets
}
