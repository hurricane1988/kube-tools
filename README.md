# kube-tools - Container Stress Test and Scanning
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
- -H, --host string       host to scan (default "localhost")
- -P, --ports string      ports to scan (default "22")
- -T, --protocol string   protocol to scan (default "tcp")
### 2. simulate memory usage.
This tool is used to simulate memory usage to test the container's memory resource utilization.
#### Usage
```bash
kube-tools memory -P 8888 -S 10240
```
- -P, --port string   Port to listen on (default "8080")
- -S, --size int      Memory size in MB
### 3. metrics
#### Usage
```shell
kube-tools metric -P 8080
```
- -P, --port string   ports to scan (default "8080")
### 4. traceroute
#### Usage
```shell
kube-tools traceroute --host localhost
```
- -H, --host string   host to scan (default "localhost")
### 5. cpu simulate cpu load
#### Usage
```shell
kube-tools cpu --cores 5 --minutes 1
```
- -C, --cores int     the number of cores to use. (default 1)
- -M, --minutes int   the minute time to run the simulation. (default 1)
### 6. generate jwt token
#### Usage
```shell
kube-tools token --issuer kubernetes --username osadmin --password password --port 8888 --expire 5200
```
```json
{
    "access_token": {
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiJwYXNzd29yZCIsImV4cCI6MTY5MTg1ODczMywiaXNzIjoiaXNzdWVyIn0.KBEOkIAYuuGiLoJyqrZnkyRs9OdvOL8qrtnu_eWWR2k",
        "expire_time": 1691858733
    },
    "refresh_token": {
        "expire_time": 1694443533,
        "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjk0NDQzNTMzLCJpc3MiOiJpc3N1ZXIifQ.8KItuXDDnSef5FBiUe1pBeSOJoC2S0Iv88a51Yiy5TM"
    }
}
```
- -E, --expire duration   The token expire time,  default 7200s. (default 7.2µs)
- -I, --issuer string     The issuer of the token. (default "issuer")
- -W, --password string   The password of the token. (default "password")
- -P, --port string       The http web service listening port. (default "8080")
- -U, --username string   The username of the token. (default "admin")
## Contribution
If you find any issues or have any improvement suggestions, feel free to raise an issue or submit a pull request. We welcome your contributions!

## License
This project is licensed under the Apache License. See the LICENSE file for details.
