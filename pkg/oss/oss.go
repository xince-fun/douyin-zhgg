package oss

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudwego/kitex/pkg/klog"
	"os"
)

// PublishVideoToPublic 保存视频到本地
func PublishVideoToPublic(video []byte, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		klog.Errorf("create %v fail, %v", filePath, err.Error())
		return err
	}
	defer file.Close()
	_, err = file.Write(video)
	if err != nil {
		klog.Errorf("write file fail, %v", err.Error())
		return err
	}
	return nil
}

// PublishVideoToOss 将视频上传到Oss
func PublishVideoToOss(objectKey string, filePath string) error {
	err := Bucket.UploadFile(objectKey, filePath, 1024*1024, oss.Routines(3))
	if err != nil {
		klog.Errorf("public %v to Oss fail, %v ", filePath, err.Error())
		return err
	}
	return nil
}

func PublishCoverToOss(objectKey string, cover *bytes.Reader) error {
	err := Bucket.PutObject(objectKey, cover)
	if err != nil {
		klog.Errorf("public %v to Oss fail, %v ", objectKey, err.Error())
		return err
	}
	return nil
}

// QueryOssVideoURL 从oss上获取播放地址
func QueryOssVideoURL(objectKey string) (string, error) {
	signedURL, err := Bucket.SignURL(objectKey, oss.HTTPPut, 60)
	if err != nil {
		klog.Errorf("Query %v Video URL fail, %v", objectKey, err.Error())
		return "", err
	}
	return signedURL, nil
}

// QueryOssCoverURL 从oss上获取视频封面地址
func QueryOssCoverURL(objectKey string) (string, error) {
	signedURL, err := Bucket.SignURL(objectKey, oss.HTTPPut, 60)
	if err != nil {
		klog.Errorf("Query %v Cover URL fail, %v", objectKey, err.Error())
		return "", err
	}
	return signedURL, nil
}
