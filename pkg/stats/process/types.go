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
)

// GlobalPids 定义全局对象
var (
	GlobalPids []*processv3.Process
)

// 定义process进程结构体
type processObject struct {
	Pid          int32          `json:"pid"`
	Name         string         `json:"name"`
	Status       string         `json:"status"`
	Parent       int32          `json:"parent"`
	ProcessUids  string         `json:"processUids"`
	NumThreads   int32          `json:"numThreads"`
	ProcessGids  string         `json:"processGids"`
	MemInfo      map[string]int `json:"memInfo"`
	ProcessGroup string         `json:"processGroup"`
	CPUPercent   float64        `json:"CPUPercent"`
	ProcessUser  string         `json:"processUser"`
	CreateTime   string         `json:"createTime"`
}

func init() {
	pids, err := processv3.Processes()
	if err != nil {
		panic(err)
	}
	GlobalPids = pids
}
