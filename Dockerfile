# 使用官方的 Golang 镜像作为基础镜像
FROM golang:1.22  AS builder

# 添加镜像维护人信息
LABEL maintainer="张晓@zhangxiao"

# golang 1.16 之后的版本，需要设置以下环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

# 在容器内创建一个目录来存放我们的应用代码
RUN mkdir /opt/golang/golang-docker

# 将工作目录切换到 /app
WORKDIR /opt/golang/golang-docker

# 将当前目录下的所有文件拷贝到 /app 目录
COPY . .

RUN go mod download && go mod verify

# 编译 Go 应用程序
RUN go build -o golang-docker cmd/main.go



# 运行时轻量镜像
FROM alpine:latest

# 声明环境变量
ENV profile =""
ENV appInfo="http://host:port/inner/appInfo"
ENV logPath=""

# 声明容器将监听的端口
EXPOSE 9091
EXPOSE 8081

WORKDIR /opt/golang/golang-docker

COPY --from=builder /opt/golang/golang-docker .

RUN echo $profile
# CMD 指令中使用 Shell 执行命令，以替换环境变量的值
CMD ["sh", "-c", "./golang-docker --env=$profile"]

