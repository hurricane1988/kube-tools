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
	"fmt"
	netv3 "github.com/shirou/gopsutil/v3/net"
	"math"
	"strconv"
	"strings"
	"time"
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
	for _, num := range removeDuplicateInts(intSlice) {
		strBuilder.WriteString(strconv.Itoa(int(num)))
	}
	return strBuilder.String()
}

func removeDuplicateInts(intput []int32) []int32 {
	intMap := make(map[int32]bool)
	var result []int32
	for _, num := range intput {
		if !intMap[num] {
			intMap[num] = true
			result = append(result, num)
		}
	}
	return result
}

func RemoveDuplicateStrings(intput []string) []string {
	strMap := make(map[string]bool)
	var result []string
	for _, item := range intput {
		if !strMap[item] {
			strMap[item] = true
			result = append(result, item)
		}
	}
	return result
}

func FormattedTime(int64Timestamp int64) string {
	nanosecondTimestamp := int64Timestamp * int64(time.Millisecond)
	return time.Unix(0, nanosecondTimestamp).Format("2006-01-02 15:04:05")
}

func ByteToMB(bytes int) float64 {
	mb := float64(bytes / 1024 / 1024)
	return math.Round(mb*100) / 100
}

func StringToInt32(str string) int32 {
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}
	return int32(num)
}
