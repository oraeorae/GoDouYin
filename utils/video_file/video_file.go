package video_file

import (
	"fmt"
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
