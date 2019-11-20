FROM golang:1.13.4-alpine3.10
MAINTAINER Docker-Test "Toney"
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update \
    && apk upgrade \
    && apk --no-cache add tzdata \
    && apk --no-cache add openssl \
    && apk --no-cache add ca-certificates \
    && rm -rf /var/cache/apk/ \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone
RUN go env -w GO111MODULE="on"
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOPRIVATE=*.gitlab.com,*.gitee.com,*.github.com
RUN go env
WORKDIR /go/src/app
COPY . .
RUN go build -o go-docker .
EXPOSE 8080
EXPOSE 8081
ENTRYPOINT  ["./start"]