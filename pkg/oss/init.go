package oss

import (
	"ByteTech-7815/douyin-zhgg/pkg/consts"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"strings"
)

var Bucket *oss.Bucket
var Path string

func InitOss() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	Path = strings.Split(dir, "/cmd")[0]
	//fmt.Println(Path)
	endpoint := consts.OssEndPoint
	accessKeyId := consts.OssAccessKeyId
	accessKeySecret := consts.OssAccessKeySecret
	bucketName := consts.OssBucketName
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		panic(err)
	}
	Bucket, err = client.Bucket(bucketName)
	if err != nil {
		panic(err)
	}
	fmt.Println(Bucket)
}
