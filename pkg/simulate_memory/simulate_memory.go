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

package simulate_memory

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

// Global variable to hold the simulated memory.
var mem []byte

type HealthyBody struct {
	Code    int       `json:"code"`
	Data    time.Time `json:"data,omitempty"`
	Message string    `json:"message"`
}

// HandleMemorySimulation handles the memory simulation request.
func HandleMemorySimulation(w http.ResponseWriter, r *http.Request) {
	// Check if memory is already allocated.
	if mem == nil {
		http.Error(w, "Memory not allocated. Please set the memory size using command line flag.", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Simulated memory allocated successfully: %d MB", len(mem)/(1024*1024))
}

func HandleHealthyCheck(w http.ResponseWriter, r *http.Request) {
	// 设置响应头，让客户端知道返回的是JSON数据
	w.Header().Set("Content-Type", "application/json")
	h := HealthyBody{
		Code:    http.StatusOK,
		Data:    time.Now(),
		Message: "The service is ok!",
	}
	// 将结构体转为json
	jsonData, err := json.MarshalIndent(h, "", " ")
	if err != nil {
		http.Error(w, "Failed to create JSON data", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

// ExecuteSimulateRaw 定义内存模拟子命令
func ExecuteSimulateRaw() *cobra.Command {
	var simulateRaw = &cobra.Command{
		Use:   "memory",
		Short: "simulate memory service.",
		Long:  "simulate memory service.",
		Run:   runner,
	}
	// 初始化命令
	simulateRaw.Flags().IntP("size", "S", 0, "Memory size in MB")
	simulateRaw.Flags().StringP("port", "P", "8080", "Port to listen on")
	return simulateRaw
}

func runner(cmd *cobra.Command, args []string) {
	memSize, _ := cmd.Flags().GetInt("size")
	port, _ := cmd.Flags().GetString("port")

	// Validate the memory size flag.
	if memSize <= 0 {
		fmt.Println("Invalid memory size. Please set a positive memory size using the -size flag.")
		os.Exit(1)
	}

	// Determine the number of available CPU cores.
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	// Allocate the memory slice with the given size.
	mem = make([]byte, memSize*1024*1024) // Convert from MB to bytes.

	// Register the memory simulation handler.
	http.HandleFunc("/raw", HandleMemorySimulation)
	http.HandleFunc("/healthy", HandleHealthyCheck)
	// Start the HTTP server.
	fmt.Printf("Starting memory simulation server on port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %s", err)
		os.Exit(1)
	}
}
