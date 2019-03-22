#!/bin/bash
#
# go build -o logServer.exe server.go
# 跨平台编译
# Linux/Ubuntu
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o logServer.exe server.go
# Mac
# CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o logServer.exe server.go

# 还可以远程发送到某个远程服务器上进行部署
# scp ./logServer.exe root@www.xxx.com:/home/xxx/tool/logServer/logServer.exe
