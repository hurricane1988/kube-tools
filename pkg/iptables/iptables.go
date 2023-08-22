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

package iptables

import (
	"fmt"
	"github.com/coreos/go-iptables/iptables"
	"github.com/spf13/cobra"
	"github.com/wonderivan/logger"
)

func ExecuteIptables() *cobra.Command {
	// cpu 定义网络链路跟踪命令
	var iptables = &cobra.Command{
		Use:   "iptables",
		Short: "generate jwt token information.",
		Long:  "generate jwt token information.",
		Run:   runner,
	}
	// 初始化命令
	// iptables.Flags().StringP("port", "P", "8080", "The http web service listening port.")
	// iptables.Flags().StringP("username", "U", "admin", "The username of the token.")
	// iptables.Flags().StringP("password", "W", "password", "The password of the token.")
	// iptables.Flags().IntP("expire", "E", 7200, "The token expire time,  default 7200s.")
	// iptables.Flags().StringP("issuer", "I", "issuer", "The issuer of the token.")
	return iptables
}

func runner(cmd *cobra.Command, args []string) {
	ipt, err := iptables.New()
	if err != nil {
		logger.Error("Error creating iptables client:", err)
	}
	rules, err := ipt.List("filter", "INPUT")
	if err != nil {
		logger.Error("Error listing iptables rules:", err)
	}
	logger.Info("iptables rules:", rules)

	for _, rule := range rules {
		fmt.Println(rule)
	}
}
