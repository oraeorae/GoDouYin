package video

import (
	"fmt"
	"go_douyin/global/variable"
	"go_douyin/utils/video_file"
	"io"
	"mime/multipart"
	"os"
)

// 保存文件
func SaveFile(file multipart.File, header *multipart.FileHeader) bool {
	filename := header.Filename
	fmt.Println(header.Filename)
	out, err := os.Create(variable.BasePath + "/output/" + filename)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	//截取第一页作为封面
	video_file.ExtractFirstFrame(variable.BasePath+"/output/"+filename, variable.BasePath+"/output/"+"第一帧图片.jpg")
	//上传到七牛云
	fmt.Println(video_file.UploadFileToQiNiu(file))

	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
