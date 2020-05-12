[![Build Status](https://travis-ci.org/xxjwxc/gmsec.svg?branch=master)](https://travis-ci.org/xxjwxc/gmsec)
[![Go Report Card](https://goreportcard.com/badge/github.com/xxjwxc/gmsec)](https://goreportcard.com/report/github.com/xxjwxc/gmsec)
[![codecov](https://codecov.io/gh/xxjwxc/gmsec/branch/master/graph/badge.svg)](https://codecov.io/gh/xxjwxc/gmsec)
[![GoDoc](https://godoc.org/github.com/xxjwxc/gmsec?status.svg)](https://godoc.org/github.com/xxjwxc/gmsec)

# gmsec 基于[ginprc](https://github.com/xxjwxc/ginrpc) 实现微服务架构集

![img](/image/ginrpc.gif)

## doc 

![doc](/image/ginrpc_doc.gif)

## 目标 
- 实现一套快捷且方便的微服务框架
- 适配主流开源库

## 已实现列表
### gorm 工具及自动生成 [gormt](https://github.com/xxjwxc/gormt)
### 接入gin框架 [go-gin](https://github.com/gin-gonic/gin)
### 自动生成api文档[swagger](https://swagger.io/) ,[markdown/mindoc](https://www.iminho.me/)
### 系统服务注册 [go-service](https://github.com/xxjwxc/go-service)
### 接入 [grpc-go](https://github.com/grpc/grpc-go)
### 服务发现 [dns](github.com/micro/mdns)
### 接入 [micro](https://github.com/gmsec/micro) 微服务

```
sudo ./gmsec install  #安装服务
sudo ./gmsec start  #启动服务
sudo ./gmsec stop #停止服务
sudo ./gmsec remove #移除服务注册
```


## 正在做
- etcd 

## 欢迎一起共建共享