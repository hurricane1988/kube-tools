# Golang Tools - Container Stress Test and Scanning
[![License](https://img.shields.io/badge/License-Apache-red?logo=apache)](LICENSE)
[![language](https://img.shields.io/badge/Language-go-blue?logo=go)](language)
[![kubernetes](https://img.shields.io/badge/Platform-kubernetes-blue?logo=kubernetes)](kubernetes)
[![runtime](https://img.shields.io/badge/Runtime-docker-blue?logo=docker)](docker)
[![os](https://img.shields.io/badge/OS-Linux-yellow?logo=linux)](os)
[![readme](https://img.shields.io/badge/Markdown-README-orange?logo=markdown)](readme)

----
## Project Overview

This project is a collection of Golang tools for container port scanning, container memory stress testing, and container CPU stress testing. These tools can help you perform performance testing and resource utilization evaluation on containers.

## Functionality

### 1. Container Port Scanner

This tool is used to scan the open ports inside a container to ensure the container's network configuration is correct and the ports are accessible.

#### Usage

```bash
kube-tools scan -H baidu.com -T tcp -P 80,443
```
### 2. simulate memory usage.
This tool is used to simulate memory usage to test the container's memory resource utilization.

#### Usage

```bash
kube-tools memory -P 8888 -S 10240
```

### 3. simulate CPU usage.
```shell
kube-tools memory -P 8888 -S 10240
```

## Contribution
If you find any issues or have any improvement suggestions, feel free to raise an issue or submit a pull request. We welcome your contributions!

## License
This project is licensed under the Apache License. See the LICENSE file for details.
