FROM golang:1.13.3-alpine3.10 AS builder
MAINTAINER Docker-Test "Toney"

# FROM golang:1.13.1 AS builder
# MAINTAINER Docker-Test "Toney"

#设置工作目录
WORKDIR $GOPATH/src/github.com/toney-li/docker-test
#将服务器的go工程代码加入到docker容器中
ADD . $GOPATH/src/github.com/toney-li/docker-test
#go构建可执行文件
RUN go build .
#暴露端口
EXPOSE 6064
#最终运行docker的命令
ENTRYPOINT  ["./main"]