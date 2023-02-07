package ffmpeg

import (
	"bytes"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"image/jpeg"
	"os"
)

func GetVideoCover(videoPath string, frameNum int) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		klog.Fatalf("生成视频封面失败, err: %s\n", err)
		return nil, err
	}
	img, err := imaging.Decode(buf)
	if err != nil {
		klog.Fatalf("图片数据解码失败, err: %s\n", err)
		return nil, err
	}

	buffer := new(bytes.Buffer)
	jpeg.Encode(buffer, img, nil)
	return buffer.Bytes(), nil
}
