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
	"flag"
	"fmt"
	"net"
	"time"
)

func checkUDPPort(target string, port int) {
	conn, err := net.DialTimeout("udp", fmt.Sprintf("%s:%d", target, port), time.Second*2)
	if err != nil {
		fmt.Printf("UDP Port %d on %s is closed\n", port, target)
		return
	}
	defer conn.Close()

	fmt.Printf("UDP Port %d on %s is open\n", port, target)
}

func checkTCPPort(target string, port int) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", target, port), time.Second*2)
	if err != nil {
		fmt.Printf("TCP Port %d on %s is closed\n", port, target)
		return
	}
	defer conn.Close()

	fmt.Printf("TCP Port %d on %s is open\n", port, target)
}

func main() {
	var (
		// UDP 端口列表，可以根据需要添加更多端口
		udpPortsList int
		// TCP 端口列表，可以根据需要添加更多端口
		tcpPortsList int
	)
	target := flag.String("target", "example.com", "Target host to check")
	flag.IntVar(&tcpPortsList, "tcp", 22, "Different TCP ports are separated by commas.")
	flag.IntVar(&udpPortsList, "udp", 53, "Different UDP ports are separated by commas.")
	flag.Parse()

	for _, port := range udpPort {
		go checkUDPPort(*target, port)
	}

	for _, port := range tcpPort {
		go checkTCPPort(*target, port)
	}

	// 防止程序立即退出，等待一段时间让 Goroutines 执行完毕
	time.Sleep(time.Second * 5)
}
