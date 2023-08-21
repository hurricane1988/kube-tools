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

package common

import (
	netv3 "github.com/shirou/gopsutil/v3/net"
	"strconv"
	"strings"
)

// FormatAddr IP地址处理
func FormatAddr(nets []netv3.InterfaceAddr) string {
	var result []string
	switch len(nets) {
	case 0:
		return ""
	default:
		for _, net := range nets {
			result = append(result, net.Addr)
		}
	}
	return strings.Join(result, "\n")
}

func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// Int32ToString 整型数组转字符串
func Int32ToString(intSlice []int32) string {
	var strBuilder strings.Builder
	for _, num := range intSlice {
		strBuilder.WriteString(strconv.Itoa(int(num)))
		strBuilder.WriteString("\n")
	}
	return strBuilder.String()
}
