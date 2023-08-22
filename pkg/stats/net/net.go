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
	netv3 "github.com/shirou/gopsutil/v3/net"
	"github.com/spf13/cobra"
	"github.com/wonderivan/logger"
)

// 定义字典常量
var familyType = map[uint32]string{
	30: "ipv6",
	2:  "ipv4",
}

// 定义链接状态map常量
var connectsStatus = []string{
	"LISTEN",
	"ESTABLISHED",
	"SYN_SENT",
	"SYN_RECV",
	"FIN_WAIT1",
	"FIN_WAIT2",
	"TIME_WAIT",
	"CLOSED",
	"CLOSE_WAIT",
	"CLOSING",
	"LAST_ACK",
}

// GlobalConnObjects 定义全局链接对象
var GlobalConnObjects []netv3.ConnectionStat

// 定义scoket套接字类型常量
var socketType = map[uint32]string{
	1: "tcp",
	2: "udp",
	3: "icmp",
	4: "ipv6",
}

// 定义协议内容切片常量
var netTypes = []string{
	"all",
	"inet",
	"inet4",
	"inet6",
	"tcp",
	"udp",
	"unix",
}

// ExecuteStatsNet 定义端口扫描执行器
func ExecuteStatsNet() *cobra.Command {
	// statCpu 定义端口扫描子命令
	var statNet = &cobra.Command{
		Use:   "net",
		Short: "stats net information",
		Long:  "stats net information",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	// 注册网卡信息采集子命令
	statNet.AddCommand(executeNetInterfaceGroup())
	statNet.AddCommand(executeConnectGroup())
	return statNet
}

// 初始化链接对象
func init() {
	// 获取网络连接信息
	conns, err := netv3.Connections("all")
	if err != nil {
		logger.Error("Error:", err)
		return
	}
	GlobalConnObjects = conns
}
