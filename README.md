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
