FROM alpine:latest
MAINTAINER Docker-Test "Toney"
RUN apk update \
    && apk upgrade \                        # 更新源和应用
    && apk --no-cache add tzdata \          # 只要go里面用到time的包，那就必装
    && apk --no-cache add openssl \         # 如果用到rsa加密解密，也需要装
    && apk --no-cache add ca-certificates \ # ca证书，这个也是必装的
    && rm -rf /var/cache/apk/ \             # 清除安装包
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \  
    && echo "Asia/Shanghai" > /etc/timezone # 这个是设置默认时区

COPY $GOPATH/src/github.com/toney-li/docker-test/main /go/demo/
WORKDIR /go/demo

EXPOSE 6064
#最终运行docker的命令
ENTRYPOINT  ["./main"]

# FROM golang:1.13.1 AS builder
# MAINTAINER Docker-Test "Toney"

# #设置工作目录
# WORKDIR $GOPATH/src/github.com/toney-li/docker-test
# #将服务器的go工程代码加入到docker容器中
# ADD . $GOPATH/src/github.com/toney-li/docker-test
# #go构建可执行文件
# RUN go build .
# #暴露端口
# EXPOSE 6064
# #最终运行docker的命令
# ENTRYPOINT  ["./main"]