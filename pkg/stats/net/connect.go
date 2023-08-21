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
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	netv3 "github.com/shirou/gopsutil/v3/net"
	"github.com/spf13/cobra"
	"os"
	"sort"
	"strings"

	"kube-tools/utils/common"
)

// executeConnectGroup 定义connect命令
func executeConnectGroup() *cobra.Command {
	var connCommand = &cobra.Command{
		Use:   "connect",
		Short: "show the connection of this system.",
		Long:  "show the connection of this system.",
		// 定义根命令执行函数
		Run: connGroup,
	}
	connCommand.Flags().StringP("type", "T", "all", "net type, support all,inet,inet4,inet6,tcp,udp,unix")
	connCommand.Flags().StringP("status", "S", "established", "connect status must be one of: listen,syn_sent,syn_recv,established,fin_wait1,fin_wait2,close_wait,closed,time_wait,last_ack,closing")
	return connCommand
}

func connGroup(cmd *cobra.Command, args []string) {
	netType, _ := cmd.Flags().GetString("type")
	connStatus, _ := cmd.Flags().GetString("status")
	// 判断输入的命令是否正常
	if common.Find(netTypes, netType) != true {
		fmt.Println("net type must be one of: all,inet,inet4,inet6,tcp,udp,unix")
	} else {
		listNetConnects(netType)
	}
	if common.Find(connectsStatus, strings.ToUpper(connStatus)) != true {
		fmt.Println("connect status must be one of: listen,syn_sent,syn_recv,established,fin_wait1,fin_wait2,close_wait,closed,time_wait,last_ack,closing")
	} else {
		sumRemoteIpMax(connStatus)
	}
}

func listNetConnects(netType string) []netv3.ConnectionStat {
	conns, _ := netv3.Connections(netType)
	// "all"：获取所有类型的连接，包括 TCP、UDP 等
	// "inet"：获取 IPv4 类型的连接
	// "inet4"：获取 IPv4 类型的连接（与 "inet" 相同）
	// "inet6"：获取 IPv6 类型的连接
	// "tcp"：获取 TCP 类型的连接。
	// "udp"：获取 UDP 类型的连接
	// "unix"：获取 Unix 域套接字类型的连接。
	// 初始化table对象
	t := table.NewWriter()
	// 设置输出到终端
	t.SetOutputMirror(os.Stdout)
	// 设置表头header
	t.AppendHeader(table.Row{"uid", "pid", "file descriptor", "family", "socket type", "local ip", "local port", "remote ip", "remote port", "status"})
	for _, c := range conns {
		t.AppendRows([]table.Row{
			{
				common.Int32ToString(c.Uids),
				c.Pid,
				c.Fd,
				familyType[c.Family],
				socketType[c.Type],
				c.Laddr.IP,
				c.Laddr.Port,
				c.Raddr.IP,
				c.Raddr.Port,
				c.Status,
			},
		})
	}
	t.Render()
	return conns
}

func sumRemoteIpMax(status string) {
	// 获取网络连接信息
	conns, err := netv3.Connections("all")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// 创建一个映射来存储每个远端IP的连接数量
	ipCountMap := make(map[string]int)

	// 遍历连接并统计每个远端IP的连接数量
	for _, conn := range conns {
		if conn.Status == strings.ToUpper(status) && conn.Raddr.IP != "" {
			ipCountMap[conn.Raddr.IP]++
		}
	}
	// 创建一个切片来对连接数量进行排序
	type IPCount struct {
		IP    string `json:"ip"`
		Count int    `json:"count"`
	}
	var ipCounts []IPCount
	for ip, count := range ipCountMap {
		ipCounts = append(ipCounts, IPCount{IP: ip, Count: count})
	}

	// 按连接数量从高到低对切片进行排序
	sort.Slice(ipCounts, func(i, j int) bool {
		return ipCounts[i].Count > ipCounts[j].Count
	})
	// 打印连接最多的远端IP
	t := table.NewWriter()
	// 设置输出到终端
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"remote ip", "connect count", "status"})
	for _, ip := range ipCounts {
		t.AppendRows([]table.Row{
			{
				ip.IP,
				ip.Count,
				strings.ToUpper(status),
			},
		})
	}
	t.SetAutoIndex(true)
	t.Render()
}
