# Dockerfile多阶构建
# https://yeasy.gitbook.io/docker_practice/image/dockerfile/arg
# 拉取基础镜像镜像
FROM golang:1.20 as builder
# 指定该image镜像的维护者
MAINTAINER hurricane1988
# 为构建设置标签
LABEL image.authors="hurricane1988"
# 此处需要使用ADD,因为有目录文件,如果只使用COPY则只会拷贝文件目录中的文件
COPY ./pkg/simulate-memory/simulate_memory.go /build/simulate_memory.go
WORKDIR /build
# 设置镜像构建代理
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GO111MODULE=on
# 设置编译环境
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOARM=6 go build -installsuffix cgo -o simulate-memory
# 运行阶段
FROM alpine:3.18.2
# 设置容器的健康检查
# HEALTHCHECK --interval=5s --timeout=3s --retries=3 CMD curl -fs http://127.0.0.1:8080/ || exit 1
# 安装基础命令工具
# RUN apk add curl busybox-extras net-tools bind-tools ethtool -U tzdata
# 从编译阶段的中拷贝编译结果到当前镜像中
COPY --from=builder /build/simulate-memory /usr/local/bin/

# 设置服务暴露端口
WORKDIR /
EXPOSE 8080
CMD ["sh"]