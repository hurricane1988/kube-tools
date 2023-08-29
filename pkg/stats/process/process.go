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

// TODO: https://godocs.io/github.com/shirou/gopsutil/v3/process

package process

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	processv3 "github.com/shirou/gopsutil/v3/process"
	"github.com/spf13/cobra"
	"os"
	"sort"
	"strconv"

	"kube-tools/utils/common"
)

// ExecuteStatsProcess 定义端口扫描执行器
func ExecuteStatsProcess() *cobra.Command {
	// statProcess 定义统计进程子命令
	var statProcess = &cobra.Command{
		Use:   "process",
		Short: "stats process information.",
		Long:  "stats process information.",
		Run:   runner,
	}
	statProcess.Flags().StringP("sort", "s", "none", "support value: cpu, mem, status, none")
	statProcess.Flags().StringP("pid", "p", "", "pid of process.")
	statProcess.Flags().StringP("io", "i", "list", "display process io information.support value: list, read, write")
	// 注册子命令
	return statProcess
}

// executeProcessGroup 定义connect命令
func runner(cmd *cobra.Command, args []string) {
	s, _ := cmd.Flags().GetString("sort")
	switch s {
	case "none":
		ListProcess()
	case "cpu":
		SortProcessByCPU()
	case "mem":
		SortProcessByMem()
	case "status":
		SortProcessByStatus()
	default:
		fmt.Println("unsupported sort value: " + s)
	}
}

func pidList() []*processv3.Process {
	return GlobalPids
}

func pidExists(pid int32) (bool, error) {
	return processv3.PidExists(pid)
}

// ListProcess 列出process明细
func ListProcess() {
	// 初始化table对象
	t := table.NewWriter()
	// 设置输出到终端
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"pid", "name", "user", "status", "parent", "numThreads", "memInfo", "createTime", "cpuPercent"})
	t.AppendSeparator()
	for _, pid := range GlobalPids {
		p := pidObject{
			Process: pid,
		}
		t.AppendRows([]table.Row{
			{
				pid.Pid,
				p.Name(),
				p.ProcessUser(),
				p.Status(),
				p.Parent(),
				p.NumThreads(),
				"rss: " + strconv.Itoa(int(common.ByteToMB(p.MemInfo()["rss"]))) +
					"\nvms: " + strconv.Itoa(int(common.ByteToMB(p.MemInfo()["vms"]))) +
					"\nhwm: " + strconv.Itoa(int(common.ByteToMB(p.MemInfo()["hwm"]))) +
					"\ndata: " + strconv.Itoa(int(common.ByteToMB(p.MemInfo()["data"]))) +
					"\nstack: " + strconv.Itoa(int(common.ByteToMB(p.MemInfo()["stack"]))) +
					"\nlocked: " + strconv.Itoa(int(common.ByteToMB(p.MemInfo()["locked"]))) +
					"\nswap: " + strconv.Itoa(int(common.ByteToMB(p.MemInfo()["swap"]))),
				p.ProcessCreateTime(),
				p.CPUPercent(),
			},
		})
		t.AppendSeparator()
	}
	t.Render()
}

// SortProcessByCPU 根据CPU使用率由高到底进行排序
func SortProcessByCPU() {
	// 创建map存储每个进程的CPU使用率
	processMap := make(map[int32]processObject)
	// 遍历所有进程
	for _, pid := range GlobalPids {
		p := pidObject{
			Process: pid,
		}
		// 存储到map中
		processMap[pid.Pid] = processObject{
			Pid:          pid.Pid,
			Name:         p.Name(),
			Status:       p.Status(),
			Parent:       p.Parent(),
			ProcessUids:  p.ProcessUids(),
			NumThreads:   p.NumThreads(),
			CPUPercent:   p.CPUPercent(),
			ProcessGids:  p.ProcessGids(),
			MemInfo:      p.MemInfo(),
			ProcessGroup: p.ProcessGroup(),
			ProcessUser:  p.ProcessUser(),
			CreateTime:   p.ProcessCreateTime(),
		}
	}
	// 创建一个切片对进程的CPU使用率进行排序
	var processList []processObject
	for _, v := range processMap {
		processList = append(processList, v)
	}
	// 按连接数量从高到低对切片进行排序
	sort.Slice(processList, func(i, j int) bool {
		return processList[i].CPUPercent > processList[j].CPUPercent
	})
	// 初始化table对象
	t := table.NewWriter()
	// 设置输出到终端
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"pid", "name", "user", "status", "parent", "numThreads", "memInfo/MB", "createTime", "cpuPercent/%"})
	t.AppendSeparator()
	for _, process := range processList {
		t.AppendRows([]table.Row{
			{
				process.Pid,
				process.Name,
				process.ProcessUser,
				process.Status,
				process.Parent,
				process.NumThreads,
				"rss: " + strconv.Itoa(int(common.ByteToMB(process.MemInfo["rss"]))) +
					"\nvms: " + strconv.Itoa(int(common.ByteToMB(process.MemInfo["vms"]))) +
					"\nhwm: " + strconv.Itoa(int(common.ByteToMB(process.MemInfo["hwm"]))) +
					"\ndata: " + strconv.Itoa(int(common.ByteToMB(process.MemInfo["data"]))) +
					"\nstack: " + strconv.Itoa(int(common.ByteToMB(process.MemInfo["stack"]))) +
					"\nlocked: " + strconv.Itoa(int(common.ByteToMB(process.MemInfo["locked"]))) +
					"\nswap: " + strconv.Itoa(int(common.ByteToMB(process.MemInfo["swap"]))),
				process.CreateTime,
				process.CPUPercent,
			},
		})
		t.AppendSeparator()
	}
	t.SetAutoIndex(true)
	t.Render()
	return
}

// SortProcessByMem 按照内存排序
func SortProcessByMem() {
	// 创建map存储每个进程的CPU使用率
	processMap := make(map[int32]processObject)
	// 遍历所有进程
	for _, pid := range GlobalPids {
		p := pidObject{
			Process: pid,
		}
		// 存储到map中
		processMap[pid.Pid] = processObject{
			Pid:          pid.Pid,
			Name:         p.Name(),
			Status:       p.Status(),
			Parent:       p.Parent(),
			ProcessUids:  p.ProcessUids(),
			NumThreads:   p.NumThreads(),
			CPUPercent:   p.CPUPercent(),
			ProcessGids:  p.ProcessGids(),
			MemInfo:      p.MemInfo(),
			ProcessGroup: p.ProcessGroup(),
			ProcessUser:  p.ProcessUser(),
			CreateTime:   p.ProcessCreateTime(),
		}
	}
	// 创建一个切片对进程的CPU使用率进行排序
	var processList []processObject
	for _, v := range processMap {
		processList = append(processList, v)
	}
	// 按连接数量从高到低对切片进行排序
	sort.Slice(processList, func(i, j int) bool {
		return processList[i].MemInfo["vms"] > processList[j].MemInfo["vms"]
	})
	// 初始化table对象
	t := table.NewWriter()
	// 设置输出到终端
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"pid", "name", "user", "status", "parent", "numThreads", "memInfo/MB", "createTime", "cpuPercent/%"})
	t.AppendSeparator()
	for _, process := range processList {
		t.AppendRows([]table.Row{
			{
				process.Pid,
				process.Name,
				process.ProcessUser,
				process.Status,
				process.Parent,
				process.NumThreads,
				"rss: " + strconv.Itoa(int(common.ByteToMB(process.MemInfo["rss"]))) +
					"\nvms: " + strconv.Itoa(int(common.ByteToMB(process.MemInfo["vms"]))) +
					"\nhwm: " + strconv.Itoa(int(common.ByteToMB(process.MemInfo["hwm"]))) +
					"\ndata: " + strconv.Itoa(int(common.ByteToMB(process.MemInfo["data"]))) +
					"\nstack: " + strconv.Itoa(int(common.ByteToMB(process.MemInfo["stack"]))) +
					"\nlocked: " + strconv.Itoa(int(common.ByteToMB(process.MemInfo["locked"]))) +
					"\nswap: " + strconv.Itoa(int(common.ByteToMB(process.MemInfo["swap"]))),
				process.CreateTime,
				process.CPUPercent,
			},
		})
		t.AppendSeparator()
	}
	t.SetAutoIndex(true)
	t.Render()
	return
}

// SortProcessByStatus 基于进程状态排序
func SortProcessByStatus() {
	// 创建map存储每个进程的CPU使用率
	processMap := make(map[int32]processObject)
	// 遍历所有进程
	for _, pid := range GlobalPids {
		p := pidObject{
			Process: pid,
		}
		// 存储到map中
		processMap[pid.Pid] = processObject{
			Pid:          pid.Pid,
			Name:         p.Name(),
			Status:       p.Status(),
			Parent:       p.Parent(),
			ProcessUids:  p.ProcessUids(),
			NumThreads:   p.NumThreads(),
			CPUPercent:   p.CPUPercent(),
			ProcessGids:  p.ProcessGids(),
			MemInfo:      p.MemInfo(),
			ProcessGroup: p.ProcessGroup(),
			ProcessUser:  p.ProcessUser(),
			CreateTime:   p.ProcessCreateTime(),
		}
	}
	// 创建一个切片对进程的CPU使用率进行排序
	var processList []processObject
	for _, v := range processMap {
		processList = append(processList, v)
	}
	// 按连接数量从高到低对切片进行排序
	sort.Slice(processList, func(i, j int) bool {
		return processList[i].Status == "running"
	})
	// 初始化table对象
	t := table.NewWriter()
	// 设置输出到终端
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"pid", "name", "user", "status", "parent", "numThreads", "memInfo/MB", "createTime", "cpuPercent/%"})
	t.AppendSeparator()
	for _, process := range processList {
		t.AppendRows([]table.Row{
			{
				process.Pid,
				process.Name,
				process.ProcessUser,
				process.Status,
				process.Parent,
				process.NumThreads,
				"rss: " + strconv.Itoa(int(common.ByteToMB(process.MemInfo["rss"]))) +
					"\nvms: " + strconv.Itoa(int(common.ByteToMB(process.MemInfo["vms"]))) +
					"\nhwm: " + strconv.Itoa(int(common.ByteToMB(process.MemInfo["hwm"]))) +
					"\ndata: " + strconv.Itoa(int(common.ByteToMB(process.MemInfo["data"]))) +
					"\nstack: " + strconv.Itoa(int(common.ByteToMB(process.MemInfo["stack"]))) +
					"\nlocked: " + strconv.Itoa(int(common.ByteToMB(process.MemInfo["locked"]))) +
					"\nswap: " + strconv.Itoa(int(common.ByteToMB(process.MemInfo["swap"]))),
				process.CreateTime,
				process.CPUPercent,
			},
		})
		t.AppendSeparator()
	}
	t.SetAutoIndex(true)
	t.Render()
	return
}
