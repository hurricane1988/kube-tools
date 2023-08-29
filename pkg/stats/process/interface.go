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

package process

import (
	processv3 "github.com/shirou/gopsutil/v3/process"
	"strconv"
	"strings"

	"kube-tools/utils/common"
)

type pidObject struct {
	Process *processv3.Process
}

type PidInterface interface {
	Name() string
	Status() string
	Parent() int32
	ProcessUids() string
	NumThreads() int32
	ProcessGids() string
	MemInfo() string
	ProcessGroup() string
	CPUPercent() string
	ProcessUser() string
	ProcessCreateTime() string
}

func (p pidObject) Name() string {
	name, err := p.Process.Name()
	if err != nil {
		return ""
	}
	return name
}

func (p pidObject) Status() string {
	status, err := p.Process.Status()
	if err != nil {
		return ""
	}
	return strings.Join(status, ",")
}

func (p pidObject) Parent() int32 {
	parent, err := p.Process.Parent()
	if err != nil {
		return 0
	}
	return parent.Pid
}

func (p pidObject) NumThreads() int32 {
	t, err := p.Process.NumThreads()
	if err != nil {
		return 0
	}
	return t
}

func (p pidObject) ProcessUids() string {
	var temp []string
	t, err := p.Process.Uids()
	if err != nil {
		return ""
	}
	for _, item := range t {
		temp = append(temp, strconv.Itoa(int(item)))
	}
	return strings.Join(common.RemoveDuplicateStrings(temp), ",")
}

func (p pidObject) ProcessGids() string {
	var temp []string
	t, err := p.Process.Gids()
	if err != nil {
		return ""
	}
	for _, item := range t {
		temp = append(temp, strconv.Itoa(int(item)))
	}
	return strings.Join(common.RemoveDuplicateStrings(temp), ",")
}

func (p pidObject) ProcessGroup() string {
	var temp []string
	g, err := p.Process.Groups()
	if err != nil {
		return ""
	}
	for _, item := range g {

		temp = append(temp, strconv.Itoa(int(item)))
	}
	return strings.Join(common.RemoveDuplicateStrings(temp), ",")
}

func (p pidObject) MemInfo() map[string]int {
	m, err := p.Process.MemoryInfo()
	if err != nil {
		return nil
	}
	return map[string]int{
		"rss":    int(m.RSS),
		"vms":    int(m.VMS),
		"hwm":    int(m.HWM),
		"data":   int(m.Data),
		"stack":  int(m.Stack),
		"locked": int(m.Locked),
		"swap":   int(m.Swap),
	}
}

func (p pidObject) ProcessCreateTime() string {
	t, err := p.Process.CreateTime()
	if err != nil {
		return ""
	}
	return common.FormattedTime(t)
}

func (p pidObject) CPUPercent() float64 {
	c, err := p.Process.CPUPercent()
	if err != nil {
		return 0
	}
	// 将浮点数格式化为保留两位小数的字符串
	cpuStr := strconv.FormatFloat(c, 'f', 2, 64)
	// 将格式化后的字符串解析回 float64 类型
	cpuFloat64, _ := strconv.ParseFloat(cpuStr, 64)
	return cpuFloat64
}

func (p pidObject) ProcessUser() string {
	u, err := p.Process.Username()
	if err != nil {
		return ""
	}
	return u
}
