package video_file

import (
	"context"
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"go_douyin/global/variable"
	"io"
	"mime/multipart"
	"os"
	"os/exec"
)

// 保存文件到服务器（基础版）
func SaveFile(file multipart.File, header *multipart.FileHeader) bool {
	filename := header.Filename
	fmt.Println(header.Filename)
	out, err := os.Create(variable.BasePath + "/config/" + filename)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// 截取视频的第一帧
func ExtractFirstFrame(videoPath, outputPath string) error {
	// 注意这里如果换到ubuntu环境可能得改一下
	cmd := exec.Command(variable.BasePath+"/lib/"+
		"ffmpeg", "-i", videoPath, "-vframes", "1", outputPath)
	// Execute command
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	return nil
}

// 上传文件到七牛云
func UploadFileToQiNiu(file multipart.File) (int, string) {
	// 获取文件大小
	fileSize, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		// handle error
	}
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		// handle error
	}
	// 构建鉴权对象（这里需要对应改成自己的）
	var AccessKey = variable.Config.GetString("CDN.AccessKey")
	var SerectKey = variable.Config.GetString("CDN.SercetKey")
	var Bucket = variable.Config.GetString("CDN.Bucket")
	var ImgUrl = variable.Config.GetString("CDN.QiniuServer")
	putPlicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SerectKey)
	upToken := putPlicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err = formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		code := 500
		return code, err.Error()
	}
	url := ImgUrl + "/" + ret.Key
	return 200, url
}
