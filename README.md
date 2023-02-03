[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/gomods/athens)
[![Kitex](https://img.shields.io/badge/Kitex-v0.4.4-green)](https://github.com/cloudwego/kitex)
[![Gorm](https://img.shields.io/badge/Gorm-v1.24.2-blue)](https://gorm.io/)
[![MySQL](https://img.shields.io/badge/MySQL-v8.0.31-red)](https://gorm.io/)


# douyin-zhgg

参考课上所讲的 `easy-note` 做的结构。

## 项目文档

https://tlkl01qmcs.feishu.cn/docx/Ap3EdPNM9ofV0axcMyccuNpLnLd

## 项目结构图

等待补充....

## 目前文件结构树

```shell
├── LICENSE
├── README.md
├── cmd
│   ├── api
│   ├── comment
│   ├── favorite
│   ├── feed
│   ├── publish
│   ├── relation
│   └── user
├── docker-compose.yml
├── go.mod
├── idl
│   ├── comment.thrift
│   ├── favorite.thrift
│   ├── feed.thrift
│   ├── publish.thrift
│   ├── relation.thrift
│   └── user.thrift
├── kitex_gen
│   ├── comment
│   ├── favorite
│   ├── feed
│   ├── publish
│   ├── relation
│   └── user
└── pkg
    ├── configs
    ├── consts
    └── errno
```

## 包括内容

把数据访问全放在`dal`目录下

对内容有疑问可以直接看idl文件，里面写的比较清晰了。

### base方向

- user 用户注册、登录、信息等
- feed流
- publish 视频上传

### 互动接口

- favorite 内容
- comment 内容

### 社交接口

- follow 内容
- follower 内容
- **暂未包括内容：** `friendUser` 之后可以加入

## 项目启动

### 1、更新依赖

```shell
go mod tidy
```

### 2、启动项目所依赖的环境

```shell
docker network create zhgg-dy_zhhhdy_net
docker-compose up
```

### 3、启动服务 暂时还未补充
