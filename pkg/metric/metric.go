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

package metric

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"net/http"
	"time"
)

// 定义全局metric指标变量
var (
	// Counter 是一个特殊的指标，用于跟踪事件计数器，例如请求的数量、错误的数量等
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "kube_tools_processed_ops_total",
		Help: "Total number of processed ops.",
	})
	opsFailed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "kube_tools_failed_ops_total",
		Help: "Total number of failed ops.",
	})
	// Histogram选项。它有多个字段，用于设置Histogram的属性，如桶的数量、桶的边界等。这个结构体通常用于创建指标的Histogram类型
	opsDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "kube_tools_duration_seconds",
		Help: "Duration of ops in seconds.",
	}, []string{"method", "host", "client", "path", "status"})
)

// ExecuteMetric 定义端口扫描执行器
func ExecuteMetric() *cobra.Command {
	// metric 定义端口扫描子命令
	var metric = &cobra.Command{
		Use:   "metric",
		Short: "expose metric information for testing.",
		Long:  "expose metric information for testing.",
		Run:   startServer,
	}
	// 初始化命令
	metric.Flags().StringP("port", "P", "8080", "ports to scan")
	return metric
}

func startServer(cmd *cobra.Command, args []string) {
	port, _ := cmd.Flags().GetString("port")
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", metricHandler)
	http.ListenAndServe(":"+port, nil)
}

func metricHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	// 模拟判断请求是否成功的逻辑，这里我们随机模拟
	isSuccess := (time.Now().UnixNano() % 2) == 0
	if isSuccess {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success"))
		opsProcessed.Inc()
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Fail"))
		opsFailed.Inc()
	}
	duration := time.Since(start).Seconds()
	// 记录HTTP请求耗时的指标
	opsDuration.WithLabelValues(r.Method, r.Host, r.RemoteAddr, r.URL.Path, fmt.Sprintf("%d", http.StatusOK)).Observe(duration)
}
