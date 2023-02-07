[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/gomods/athens)
[![Kitex](https://img.shields.io/badge/Kitex-v0.4.4-green)](https://github.com/cloudwego/kitex)
[![Gorm](https://img.shields.io/badge/Gorm-v1.24.2-blue)](https://gorm.io/)
[![MySQL](https://img.shields.io/badge/MySQL-v8.0.31-red)](https://gorm.io/)
[![Hertz](https://img.shields.io/badge/Hertz-v0.5.2-purple)](https://github.com/cloudwego/hertz)

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
│   ├── public
│   ├── relation
│   └── user
├── docker-compose.yml
├── go.mod
├── idl
│   ├── comment.thrift
│   ├── favorite.thrift
│   ├── feed.thrift
│   ├── public.thrift
│   ├── relation.thrift
│   └── user.thrift
├── kitex_gen
│   ├── comment
│   ├── favorite
│   ├── feed
│   ├── public
│   ├── relation
│   └── user
└── pkg
    ├── configs
    ├── consts
    └── errno
```

## 注意事项

- 把数据访问全放在`dal`目录下


- 错误生成在`pkg/errno`下，之后内容多的话可以把`Errcode`和`Err`单独拎出来放在一个文件


- 一些参数条件判断可以写在thrift文件中，在`根目录`中重新生成
```shell
kitex --thrift-plugin validator -module ByteTech-7815/douyin-zhgg idl/xx.thrift
```
这个在`cmd/xx`执行
```shell
kitex --thrift-plugin validator -module ByteTech-7815/douyin-zhgg -service xx -use ByteTech-7815/douyin-zhgg/kitex_gen ../../idl/xx.thrift
```
如果需要更新`api.thrift`

在`cmd/api`下执行
```shell
hz update -mod ByteTech-7815/douyin-zhgg/cmd/api -idl ../../idl/api.thrift
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
- **暂未包括内容：** `friendUser` 之后可以加入

## 开发注意

```shell
git clone git@github.com:XinceChan/douyin-zhgg.git
```

## 项目启动

### 1、更新依赖

```shell
go mod tidy
```

### 2、启动项目所依赖的环境

```shell
docker network create zhgg-dy_zhggdy_net
docker-compose up
```

补充 可以使用以下命令检查网络
```shell
docker network ls
```

### 3、启动服务 暂时还未补充

```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

```shell
cd cmd/api
go run .
```


### 4、API Requests

推荐`postman`

小队邀请： https://app.getpostman.com/join-team?invite_code=50b3f99fbf2f8f2ce2c6ffe942688ff3&target_code=43549d582ad5959e2d665f4dfbbbcb1b

#### Register
```json
{
    "username":"test1",
    "password":"testtest"
}
```
##### response
```javascript
// successful
{
    "status_code": 0,
    "status_msg": "Success",
    "user_id": 6,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QxIiwidXNlcl9pZCI6NiwiZXhwIjoxNjc1NTkyODg2fQ.yuAu2Ov-Eg0OgR-LTXtu8FAD9ybIRQdQ7EDhTT9Z7x8"
}
// failed
{
    "status_code": 10010,
    "status_msg": "User already exists",
    "data": null
}
```

#### Login
```json
{
  "username":"test3",
  "password":"testtest"
}
```

##### response
```javascript
// successfully
{
    "status_code": 0,
    "status_msg": "Success",
    "user_id": 8,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzU2MDcxOTIsImlkIjo4LCJvcmlnX2lhdCI6MTY3NTYwMzU5Mn0.-OnqilfucCUSwoQ-MuYfhesK1BtUzGEhTwhM1EzQagw"
}
// failed
{
    "status_code": 10012,
    "status_msg": "Authorization failed",
    "data": null
}
```

#### User Info
```
127.0.0.1:8080/douyin/user/?user_id=8
```

```javascript
{
    "status_code": 0,
    "status_msg": "Success",
    "user": {
        "id": 8,
        "name": "test3",
        "follow_count": 0,
        "follower_count": 0,
        "is_follow": true
    }
}
```
