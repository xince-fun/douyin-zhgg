# douyin-zhgg

参考课上所讲的 `easy-note` 做的结构。

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
- **暂未包括内容：** friendUser 之后可以加入

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
