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

package cpu

import (
	"github.com/spf13/cobra"
	"github.com/wonderivan/logger"
	"runtime"
	"sync"
	"time"
)

// ExecuteCPUSimulateLoad 定义端口扫描执行器
func ExecuteCPUSimulateLoad() *cobra.Command {
	// cpu 定义网络链路跟踪命令
	var cpu = &cobra.Command{
		Use:   "cpu",
		Short: "cpu core load simulation service.",
		Long:  "cpu core load simulation service.",
		Run:   runner,
	}
	// 初始化命令
	cpu.Flags().IntP("cores", "C", 1, "the number of cores to use.")
	cpu.Flags().IntP("minutes", "M", 1, "the minute time to run the simulation.")
	return cpu
}

func runner(cmd *cobra.Command, args []string) {
	cores, _ := cmd.Flags().GetInt("cores")
	minutes, _ := cmd.Flags().GetInt("time")
	// 设置要使用的核数
	runtime.GOMAXPROCS(cores)
	done := make(chan bool)
	// 使用等待组来等待所有Goroutine完成
	var wg sync.WaitGroup
	// 启动Goroutine
	for i := 0; i < cores; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// 执行CPU测试
			cpuSimulateLoad(id, done)
		}(i)
	}
	// 运行指定时间后停止CPU负载模拟
	runTime := time.Duration(minutes) * time.Minute
	logger.Info("Running load simulation for %v...", runTime)
	time.Sleep(runTime)

	// 停止所有Goroutines
	for i := 0; i < cores; i++ {
		<-done
	}
	logger.Info("All load simulation stopped.")
}

// 模拟CPU负载
func cpuSimulateLoad(coreID int, done chan<- bool) {
	logger.Info("Core %d: Starting load simulation", coreID)
	for {
		// 模式CPU负载
		_ = 2 + 2
	}
}
