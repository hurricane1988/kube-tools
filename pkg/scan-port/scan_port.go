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

package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wonderivan/logger"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	Execute()
}

// 定义全局变量
var rootCmd = &cobra.Command{
	Use:   "scan-port",
	Short: "scan-port is a tool to scan ports",
	Long:  "scan-port is a tool to scan ports",
	Run:   runnerScan,
}

// Execute 定义cobra执行器
func Execute() {
	rootCmd.Flags().StringP("host", "H", "localhost", "host to scan")
	rootCmd.Flags().StringP("protocol", "T", "tcp", "protocol to scan")
	rootCmd.Flags().StringP("ports", "P", "22", "ports to scan")
	if err := rootCmd.Execute(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}

// runner 扫描端口
func runnerScan(cmd *cobra.Command, args []string) {
	host, _ := cmd.Flags().GetString("host")
	protocol, _ := cmd.Flags().GetString("protocol")
	portsStr, _ := cmd.Flags().GetString("ports")
	ports := parsePorts(portsStr)
	logger.Info("Scanning...")
	logger.Info("Host:", host)
	logger.Info("Protocol:", protocol)
	logger.Info("Ports:", ports)
	scanner(host, protocol, ports)
	logger.Info("Scanning finished")
}

// parsePorts 解析端口
func parsePorts(portsStr string) []int {
	portsList := strings.Split(portsStr, ",")
	var ports []int
	for _, portStr := range portsList {
		port, err := strconv.Atoi(portStr)
		if err != nil {
			logger.Error(err)
			os.Exit(1)
		}
		ports = append(ports, port)
	}
	return ports
}

// 已并发方式扫描端口
func scanner(host string, protocol string, ports []int) {
	var wg sync.WaitGroup
	for _, port := range ports {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			scanPort(host, protocol, port)
		}(port)
	}
	wg.Wait()
}

func scanPort(host string, protocol string, port int) {
	conn, err := net.DialTimeout(protocol, fmt.Sprintf("%s:%d", host, port), 2*time.Second)
	if err != nil {
		logger.Info("Port %d is close", port)
	} else {
		logger.Info("Port %d is open", port)
		conn.Close()
	}
}
